package kuberbac

import (
	rbacv1 "k8s.io/api/rbac/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"os"
)

func (k *KubeRBAC) CreateClusterRoleBinding() error {
	clusterRoleBinding := createClusterRoleBinding(k.name, k.namespace)
	_, err := k.client.RbacV1().ClusterRoleBindings().Create(ctx, clusterRoleBinding, metav1.CreateOptions{})
	if err != nil {
		return err
	}

	return k.printer.PrintObj(clusterRoleBinding, os.Stdout)
}

func (k *KubeRBAC) DeleteClusterRoleBinding() error {
	return nil
}

func createClusterRoleBinding(name, namespace string) *rbacv1.ClusterRoleBinding {
	clusterRoleBinding := &rbacv1.ClusterRoleBinding{
		TypeMeta: metav1.TypeMeta{
			APIVersion: metav1.SchemeGroupVersion.String(),
			Kind:       ClusterRoleBindingKind,
		},
		ObjectMeta: metav1.ObjectMeta{
			Name: name,
		},
		Subjects: []rbacv1.Subject{
			{
				Kind:      ServiceAccountKind,
				Name:      name,
				Namespace: namespace,
			},
		},
		RoleRef: rbacv1.RoleRef{
			APIGroup: "rbac.authorization.k8s.io",
			Kind:     ClusterRoleKind,
			Name:     name,
		},
	}

	return clusterRoleBinding
}
