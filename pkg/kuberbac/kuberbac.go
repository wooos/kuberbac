package kuberbac

import (
	"context"
	"fmt"
	"k8s.io/cli-runtime/pkg/genericclioptions"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/tools/clientcmd/api"
)

type KubeRBAC struct {
	client           kubernetes.Interface
	config           api.Config
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
	config, err := kubeConfigFlags.ToRESTConfig()
	if err != nil {
		return nil, err
	}

	client, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, err
	}

	namespace, enforceNamespace, err := kubeConfigFlags.ToRawKubeConfigLoader().Namespace()
	if err != nil {
		return nil, err
	}

	apiConfig, err := kubeConfigFlags.ToRawKubeConfigLoader().RawConfig()
	if err != nil {
		return nil, err
	}

	kubeRBAC := &KubeRBAC{
		client:           client,
		config:           apiConfig,
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

func (k *KubeRBAC) ShowConfig(global bool) error {
	serviceAccount, err := k.GetServiceAccount()
	if err != nil {
		return err
	}

	secretName := serviceAccount.Secrets[0].Name
	secret, err := k.GetSecret(secretName)
	if err != nil {
		return err
	}


	currentContext := k.config.Contexts[k.config.CurrentContext]
	currentCluster := k.config.Clusters[currentContext.Cluster]

	apiContext := &api.Context{
		Cluster:          "default",
		AuthInfo:         "default",
		Namespace:        "",
	}

	if !global {
		apiContext.Namespace = k.namespace
	}

	apiConfig := api.Config{
		Clusters: map[string]*api.Cluster{
			"default": currentCluster,
		},
		AuthInfos: map[string]*api.AuthInfo{
			"default": {
				Token: string(secret.Data["token"]),
			},
		},
		Contexts: map[string]*api.Context{
			"default": apiContext,
		},
		CurrentContext: "default",
	}

	data, err := clientcmd.Write(apiConfig)
	if err != nil {
		return err
	}

	fmt.Printf("%s", data)
	return nil
}