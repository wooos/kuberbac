package kuberbac

import (
	"os"

	rbacv1 "k8s.io/api/rbac/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (k *KubeRBAC) CreateRoleBinding() error {
	roleBinding := createRoleBinding(k.name, k.namespace)

	roleBinding, err := k.client.RbacV1().RoleBindings(k.namespace).Create(ctx, roleBinding, metav1.CreateOptions{})
	if err != nil {
		return err
	}

	return k.printer.PrintObj(roleBinding, os.Stdout)
}

func (k *KubeRBAC) DeleteRoleBinding() error {
	return k.client.RbacV1().RoleBindings(k.namespace).Delete(ctx, k.name, metav1.DeleteOptions{})
}

func createRoleBinding(name, namespace string) *rbacv1.RoleBinding {
	roleBinding := &rbacv1.RoleBinding{
		ObjectMeta: metav1.ObjectMeta{
			Name:      name,
			Namespace: namespace,
		},
		Subjects: []rbacv1.Subject{
			{
				Kind:      "ServiceAccount",
				Name:      name,
				Namespace: namespace,
			},
		},
		RoleRef: rbacv1.RoleRef{
			APIGroup: "rbac.authorization.k8s.io",
			Kind:     "Role",
			Name:     name,
		},
	}

	return roleBinding
}
