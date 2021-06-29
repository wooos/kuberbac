package cmd

import (
	"fmt"
	"kuberbac/pkg/kuberbac"
	"os"

	"github.com/spf13/cobra"
	"k8s.io/cli-runtime/pkg/genericclioptions"
)

type CreateOptions struct {
	Name        string
	Admin       bool
	Global bool
	ConfigFlags *genericclioptions.ConfigFlags
	PrintFlags  *genericclioptions.PrintFlags
}

func NewCreateCommand(kubeConfigFlags *genericclioptions.ConfigFlags) *cobra.Command {
	opt := CreateOptions{
		ConfigFlags: kubeConfigFlags,
		PrintFlags:  genericclioptions.NewPrintFlags("created"),
	}

	cmd := &cobra.Command{
		Use:   "create NAME",
		Short: "Quick create of kubernetes RBAC",
		Run: func(cmd *cobra.Command, args []string) {
			opt.Validate(cmd, args)
			opt.RunCreate()
		},
	}

	opt.PrintFlags.AddFlags(cmd)

	flags := cmd.Flags()
	flags.BoolVar(&opt.Admin, "admin", false, "If true, create admin permission, else create readonly permission")
	flags.BoolVar(&opt.Global, "global", false, "If true, create or delete ClusterRole, ClusterRoleBinding, else create or delete Role, RoleBinding")
	return cmd
}

func (opt *CreateOptions) Validate(cmd *cobra.Command, args []string) {
	if len(args) != 1 {
		fmt.Fprint(os.Stderr, "Error: NAME is required.\nSee 'kuberbac create -h' for help and examples\n")
		os.Exit(1)
	}

	opt.Name = args[0]
	return
}

func (opt *CreateOptions) RunCreate() {
	kubeRBAC, err := kuberbac.NewKubeRBAC(opt.ConfigFlags, opt.PrintFlags, opt.Name, opt.Admin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err)
		os.Exit(2)
	}

	if err := kubeRBAC.Create(opt.Global); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err)
		os.Exit(2)
	}
}
