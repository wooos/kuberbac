package util

import (
	"fmt"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
)

func CheckError(e error) error {
	switch {
	case apierrors.IsAlreadyExists(e):
		fmt.Println(e)
		return nil
	default:
		return e
	}
}
