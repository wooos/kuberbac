package kuberbac

import (
	"kuberbac/pkg/kuberbac/printer"
	"kuberbac/pkg/kuberbac/util"

	rbacv1 "k8s.io/api/rbac/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// CreateRole create a Role in kubernetes
func (k *KubeRBAC) CreateRole() error {
	role := createRole(k.name, k.namespace, k.admin)
	_, err := k.client.RbacV1().Roles(k.namespace).Create(ctx, role, metav1.CreateOptions{})
	if err != nil {
		return util.CheckError(err)
	}

	return printer.PrintObj(role, "created")
}

func (k *KubeRBAC) DeleteRole() error {
	return k.client.RbacV1().Roles(k.namespace).Delete(ctx, k.name, metav1.DeleteOptions{})
}

func createRole(name, namespace string, admin bool) *rbacv1.Role {
	rules := []rbacv1.PolicyRule{{APIGroups: []string{"*"}, Resources: []string{"*"}, Verbs: []string{"GET", "LIST"}}}
	if admin {
		rules = []rbacv1.PolicyRule{{APIGroups: []string{"*"}, Resources: []string{"*"}, Verbs: []string{"*"}}}
	}

	role := &rbacv1.Role{
		TypeMeta: metav1.TypeMeta{
			Kind:       RoleKind,
			APIVersion: rbacv1.SchemeGroupVersion.String(),
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      name,
			Namespace: namespace,
		},
		Rules: rules,
	}

	return role
}
