package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"k8s.io/cli-runtime/pkg/genericclioptions"
	"kuberbac/pkg/kuberbac"
)

type createOptions struct {
	name            string
	admin           bool
	global          bool
	silence         bool
	kubeConfigFlags *genericclioptions.ConfigFlags
}

func newCreateCmd(kubeConfigFlags *genericclioptions.ConfigFlags) *cobra.Command {
	opt := createOptions{
		kubeConfigFlags: kubeConfigFlags,
	}

	cmd := &cobra.Command{
		Use:   "create NAME",
		Short: "Quick create of kubernetes RBAC",
		Args:  opt.Validate,
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := opt.RunCreate(); err != nil {
				return err
			}

			return nil
		},
	}

	flags := cmd.Flags()
	flags.BoolVar(&opt.admin, "admin", false, "If true, create admin permission, else create readonly permission")
	flags.BoolVar(&opt.global, "global", false, "If true, create ClusterRole, ClusterRoleBinding, else create Role, RoleBinding")
	flags.BoolVarP(&opt.silence, "silence", "s", false, "If true, only print kubeconfig")
	return cmd
}

func (opt *createOptions) Validate(cmd *cobra.Command, args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("NAME is required.\nSee 'kuberbac create -h' for help and examples")
	}

	if len(args) > 1 {
		return fmt.Errorf("exactly one NAME is required, got %d\nSee 'kuberbac create -h' for help and examples", len(args))
	}

	opt.name = args[0]
	return nil
}

func (opt *createOptions) RunCreate() error {
	kubeRBAC, err := kuberbac.NewKubeRBAC(opt.kubeConfigFlags, opt.name, opt.admin)
	if err != nil {
		return err
	}

	if err = kubeRBAC.Create(opt.global); err != nil {
		return err
	}

	if err = kubeRBAC.ShowConfig(opt.global); err != nil {
		return err
	}

	return nil
}
