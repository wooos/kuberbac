package cmd

import (
	"fmt"
	"kuberbac/pkg/kuberbac"

	"github.com/spf13/cobra"
	"k8s.io/cli-runtime/pkg/genericclioptions"
)

type DeleteOptions struct {
	Name        string
	Admin       bool
	Global      bool
	ConfigFlags *genericclioptions.ConfigFlags
}

func NewDeleteCommand(kubeConfigFlags *genericclioptions.ConfigFlags) *cobra.Command {
	opt := DeleteOptions{
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

func (opt *DeleteOptions) Validate(cmd *cobra.Command, args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("NAME is required.\nSee 'kuberbac delete -h' for help and examples")
	}

	if len(args) > 1 {
		return fmt.Errorf("exactly one NAME is required, got %d\nSee 'kuberbac create -h' for help and examples", len(args))
	}

	opt.Name = args[0]
	return nil
}

func (opt *DeleteOptions) RunDelete() error {
	kubeRBAC, err := kuberbac.NewKubeRBAC(opt.ConfigFlags, opt.Name, opt.Admin)
	if err != nil {
		return err
	}

	if err := kubeRBAC.Delete(opt.Global); err != nil {
		return err
	}

	return nil
}
