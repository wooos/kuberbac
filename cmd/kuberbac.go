package cmd

import (
	"github.com/spf13/cobra"
	"k8s.io/cli-runtime/pkg/genericclioptions"
)

func NewCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "kuberbac",
		Short: "Quick create or delete of kubernetes RBAC",
	}

	flags := cmd.PersistentFlags()

	kubeConfigFlags := genericclioptions.NewConfigFlags(true).WithDeprecatedPasswordFlag()
	kubeConfigFlags.AddFlags(flags)

	cmd.AddCommand(
		NewCreateCommand(kubeConfigFlags),
		NewDeleteCommand(kubeConfigFlags),
		NewCompletionCommand(),
	)
	return cmd
}
