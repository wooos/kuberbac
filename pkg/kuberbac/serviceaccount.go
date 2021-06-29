package kuberbac

import (
	"os"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (k *KubeRBAC) CreateServiceAccount() error {
	serviceAccount := createServiceAccount(k.name, k.namespace)

	serviceAccount, err := k.client.CoreV1().ServiceAccounts(k.namespace).Create(ctx, serviceAccount, metav1.CreateOptions{})
	if err != nil {
		return err
	}

	return k.printer.PrintObj(serviceAccount, os.Stdout)
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
		TypeMeta: metav1.TypeMeta{APIVersion: corev1.SchemeGroupVersion.String(), Kind: "ServiceAccount"},
	}

	return serviceAccount
}
