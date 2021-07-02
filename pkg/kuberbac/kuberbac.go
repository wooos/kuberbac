package kuberbac

import (
	"context"

	"k8s.io/cli-runtime/pkg/genericclioptions"
	"k8s.io/client-go/kubernetes"
)

type KubeRBAC struct {
	client           kubernetes.Interface
	name             string
	namespace        string
	enforceNamespace bool
	admin            bool
}

const (
	ServiceAccountKind     = "ServiceAccount"
	RoleKind               = "Role"
	RoleBindingKind        = "RoleBinding"
	ClusterRoleKind        = "ClusterRole"
	ClusterRoleBindingKind = "ClusterRoleBinding"
)

var (
	ctx = context.TODO()
)

func NewKubeRBAC(kubeConfigFlags *genericclioptions.ConfigFlags, name string, admin bool) (*KubeRBAC, error) {
	cfg, err := kubeConfigFlags.ToRESTConfig()
	if err != nil {
		return nil, err
	}

	client, err := kubernetes.NewForConfig(cfg)
	if err != nil {
		return nil, err
	}

	namespace, enforceNamespace, err := kubeConfigFlags.ToRawKubeConfigLoader().Namespace()
	if err != nil {
		return nil, err
	}

	kubeRBAC := &KubeRBAC{
		client:           client,
		name:             name,
		namespace:        namespace,
		enforceNamespace: enforceNamespace,
		admin:            admin,
	}

	return kubeRBAC, nil
}

func (k *KubeRBAC) Create(global bool) error {
	if err := k.CreateServiceAccount(); err != nil {
		return err
	}

	if global {
		if err := k.CreateClusterRole(); err != nil {
			return err
		}

		if err := k.CreateClusterRoleBinding(); err != nil {
			return err
		}

		return nil
	}

	if err := k.CreateRole(); err != nil {
		return err
	}

	if err := k.CreateRoleBinding(); err != nil {
		return err
	}

	return nil
}

func (k *KubeRBAC) Delete(global bool) error {
	if global {
		if err := k.DeleteClusterRoleBinding(); err != nil {
			return err
		}

		if err := k.DeleteClusterRole(); err != nil {
			return err
		}
	} else {
		if err := k.DeleteRoleBinding(); err != nil {
			return err
		}

		if err := k.DeleteRole(); err != nil {
			return err
		}
	}

	if err := k.DeleteServiceAccount(); err != nil {
		return err
	}

	return nil
}
