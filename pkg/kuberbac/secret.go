package kuberbac

import (
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (k *KubeRBAC) GetSecret(name string) (*corev1.Secret, error) {
	secret, err := k.client.CoreV1().Secrets(k.namespace).Get(ctx, name, metav1.GetOptions{})
	if err != nil {
		return nil, err
	}

	return secret, nil
}