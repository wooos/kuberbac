package cmd

import (
	"fmt"
	"kuberbac/pkg/kuberbac"

	"github.com/spf13/cobra"
	"k8s.io/cli-runtime/pkg/genericclioptions"
)

type CreateOptions struct {
	Name        string
	Admin       bool
	Global      bool
	Silence     bool
	ConfigFlags *genericclioptions.ConfigFlags
}

func NewCreateCommand(kubeConfigFlags *genericclioptions.ConfigFlags) *cobra.Command {
	opt := CreateOptions{
		ConfigFlags: kubeConfigFlags,
	}

	cmd := &cobra.Command{
		Use:   "create NAME",
		Short: "Quick create of kubernetes RBAC",
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := opt.Validate(cmd, args); err != nil {
				return err
			}

			if err := opt.RunCreate(); err != nil {
				return err
			}

			return nil
		},
	}

	flags := cmd.Flags()
	flags.BoolVar(&opt.Admin, "admin", false, "If true, create admin permission, else create readonly permission")
	flags.BoolVar(&opt.Global, "global", false, "If true, create ClusterRole, ClusterRoleBinding, else create Role, RoleBinding")
	flags.BoolVarP(&opt.Silence, "silence", "s", false, "If true, only print kubeconfig")
	return cmd
}

func (opt *CreateOptions) Validate(cmd *cobra.Command, args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("NAME is required.\nSee 'kuberbac create -h' for help and examples")
	}

	if len(args) > 1 {
		return fmt.Errorf("exactly one NAME is required, got %d\nSee 'kuberbac create -h' for help and examples", len(args))
	}

	opt.Name = args[0]
	return nil
}

func (opt *CreateOptions) RunCreate() error {
	kubeRBAC, err := kuberbac.NewKubeRBAC(opt.ConfigFlags, opt.Name, opt.Admin)
	if err != nil {
		return err
	}

	if err := kubeRBAC.Create(opt.Global); err != nil {
		return err
	}

	return nil
}
