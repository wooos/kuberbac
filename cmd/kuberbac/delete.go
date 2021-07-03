package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"k8s.io/cli-runtime/pkg/genericclioptions"
	"kuberbac/pkg/kuberbac"
)

type deleteOptions struct {
	Name        string
	Admin       bool
	Global      bool
	ConfigFlags *genericclioptions.ConfigFlags
}

func newDeleteCommand(kubeConfigFlags *genericclioptions.ConfigFlags) *cobra.Command {
	opt := deleteOptions{
		ConfigFlags: kubeConfigFlags,
	}

	cmd := &cobra.Command{
		Use:   "delete <NAME>",
		Short: "Quick delete of kubernetes RBAC",
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := opt.Validate(cmd, args); err != nil {
				return err
			}

			if err := opt.RunDelete(); err != nil {
				return err
			}

			return nil
		},
	}

	flags := cmd.Flags()
	flags.BoolVar(&opt.Global, "global", false, "If true, delete ClusterRole, ClusterRoleBinding, else delete Role, RoleBinding")

	return cmd
}

func (opt *deleteOptions) Validate(cmd *cobra.Command, args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("NAME is required.\nSee 'kuberbac delete -h' for help and examples")
	}

	if len(args) > 1 {
		return fmt.Errorf("exactly one NAME is required, got %d\nSee 'kuberbac create -h' for help and examples", len(args))
	}

	opt.Name = args[0]
	return nil
}

func (opt *deleteOptions) RunDelete() error {
	kubeRBAC, err := kuberbac.NewKubeRBAC(opt.ConfigFlags, opt.Name, opt.Admin)
	if err != nil {
		return err
	}

	if err := kubeRBAC.Delete(opt.Global); err != nil {
		return err
	}

	return nil
}
