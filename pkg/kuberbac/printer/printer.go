package printer

import (
	"fmt"
	"k8s.io/apimachinery/pkg/api/meta"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"strings"
)

func PrintObj(obj runtime.Object, operation string) error {
	var name string
	if acc, err := meta.Accessor(obj); err == nil {
		if n := acc.GetName(); len(n) > 0 {
			name = n
		}
	}

	return printObj(name, operation, GetObjectGroupKind(obj))
}

func GetObjectGroupKind(obj runtime.Object) schema.GroupKind {
	return obj.GetObjectKind().GroupVersionKind().GroupKind()
}

func printObj(name string, operation string, groupKind schema.GroupKind) error {
	if len(operation) > 0 {
		operation = " " + operation
	}

	if len(groupKind.Group) == 0 {
		fmt.Printf("%s/%s%s\n", strings.ToLower(groupKind.Kind), name, operation)
		return nil
	}

	fmt.Printf("%s.%s/%s%s\n", strings.ToLower(groupKind.Kind), groupKind.Group, name, operation)
	return nil
}
