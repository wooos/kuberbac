package kuberbac

import (
	"kuberbac/pkg/kuberbac/printer"
	"kuberbac/pkg/kuberbac/util"

	rbacv1 "k8s.io/api/rbac/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (k *KubeRBAC) CreateClusterRole() error {
	clusterRole := createClusterRole(k.name, k.admin)
	_, err := k.client.RbacV1().ClusterRoles().Create(ctx, clusterRole, metav1.CreateOptions{})
	if err != nil {
		return util.CheckError(err)
	}

	return printer.PrintObj(clusterRole, "created")
}

func (k *KubeRBAC) DeleteClusterRole() error {
	return nil
}

func createClusterRole(name string, admin bool) *rbacv1.ClusterRole {
	rules := []rbacv1.PolicyRule{{APIGroups: []string{"*"}, Resources: []string{"*"}, Verbs: []string{"get", "list"}}}
	if admin {
		rules = []rbacv1.PolicyRule{{APIGroups: []string{"*"}, Resources: []string{"*"}, Verbs: []string{"*"}}}
	}

	clusterRole := &rbacv1.ClusterRole{
		TypeMeta: metav1.TypeMeta{
			APIVersion: rbacv1.SchemeGroupVersion.String(),
			Kind:       "ClusterRole",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name: name,
		},
		Rules: rules,
	}

	return clusterRole
}
