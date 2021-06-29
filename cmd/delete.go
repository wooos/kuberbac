package cmd

import (
	"fmt"
	"kuberbac/pkg/kuberbac"
	"os"

	"github.com/spf13/cobra"
	"k8s.io/cli-runtime/pkg/genericclioptions"
)

type DeleteOptions struct {
	Name        string
	Admin       bool
	ConfigFlags *genericclioptions.ConfigFlags
	PrintFlags  *genericclioptions.PrintFlags
}

func NewDeleteCommand(kubeConfigFlags *genericclioptions.ConfigFlags) *cobra.Command {
	opt := DeleteOptions{
		ConfigFlags: kubeConfigFlags,
	}

	cmd := &cobra.Command{
		Use:   "delete <NAME>",
		Short: "Quick delete of kubernetes RBAC",
		Run: func(cmd *cobra.Command, args []string) {
			opt.Validate(cmd, args)
			opt.RunDelete()
		},
	}

	return cmd
}

func (opt *DeleteOptions) Validate(cmd *cobra.Command, args []string) {
	if len(args) != 1 {
		fmt.Fprint(os.Stderr, "Error: NAME is required.\nSee 'kuberbac delete -h' for help and examples\n")
		os.Exit(1)
	}

	opt.Name = args[0]
	return
}

func (opt *DeleteOptions) RunDelete() {
	kubeRBAC, err := kuberbac.NewKubeRBAC(opt.ConfigFlags, opt.PrintFlags, opt.Name, opt.Admin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err)
		os.Exit(2)
	}

	if err := kubeRBAC.Delete(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err)
		os.Exit(2)
	}
}
