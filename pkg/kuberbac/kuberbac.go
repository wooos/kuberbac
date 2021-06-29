package kuberbac

import (
	"context"

	"k8s.io/cli-runtime/pkg/genericclioptions"
	"k8s.io/cli-runtime/pkg/printers"
	"k8s.io/client-go/kubernetes"
)

type KubeRBAC struct {
	client           kubernetes.Interface
	name             string
	namespace        string
	enforceNamespace bool
	admin            bool
	printer          printers.ResourcePrinter
}

var (
	ctx = context.Background()
)

func NewKubeRBAC(kubeConfigFlags *genericclioptions.ConfigFlags, printFlags *genericclioptions.PrintFlags, name string, admin bool) (*KubeRBAC, error) {
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

	printer, err := printFlags.ToPrinter()
	if err != nil {
		return nil, err
	}

	kubeRBAC := &KubeRBAC{
		client:           client,
		name:             name,
		namespace:        namespace,
		enforceNamespace: enforceNamespace,
		admin:            admin,
		printer:          printer,
	}

	return kubeRBAC, nil
}

func (k *KubeRBAC) Create() error {
	if err := k.CreateServiceAccount(); err != nil {
		return err
	}

	if err := k.CreateRole(); err != nil {
		return err
	}

	if err := k.CreateRoleBinding(); err != nil {
		return err
	}

	return nil
}

func (k *KubeRBAC) Delete() error {
	if err := k.DeleteRoleBinding(); err != nil {
		return err
	}

	if err := k.DeleteRole(); err != nil {
		return err
	}

	if err := k.DeleteServiceAccount(); err != nil {
		return err
	}

	return nil
}
