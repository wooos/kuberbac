package kuberbac

import (
	"kuberbac/pkg/kuberbac/printer"
	"kuberbac/pkg/kuberbac/util"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (k *KubeRBAC) CreateServiceAccount() error {
	serviceAccount := createServiceAccount(k.name, k.namespace)

	_, err := k.client.CoreV1().ServiceAccounts(k.namespace).Create(ctx, serviceAccount, metav1.CreateOptions{})
	if err != nil {
		return util.CheckError(err)
	}

	return printer.PrintObj(serviceAccount, "created")
}

func (k *KubeRBAC) DeleteServiceAccount() error {
	return k.client.CoreV1().ServiceAccounts(k.namespace).Delete(ctx, k.name, metav1.DeleteOptions{})
}

func createServiceAccount(name, namespace string) *corev1.ServiceAccount {
	serviceAccount := &corev1.ServiceAccount{
		ObjectMeta: metav1.ObjectMeta{
			Name:      name,
			Namespace: namespace,
		},
		TypeMeta: metav1.TypeMeta{APIVersion: corev1.SchemeGroupVersion.String(), Kind: ServiceAccountKind},
	}

	return serviceAccount
}

func (k *KubeRBAC) GetServiceAccount() (*corev1.ServiceAccount, error) {
	serviceAccount, err := k.client.CoreV1().ServiceAccounts(k.namespace).Get(ctx, k.name, metav1.GetOptions{})
	if err != nil {
		return nil, err
	}

	return serviceAccount, nil
}
