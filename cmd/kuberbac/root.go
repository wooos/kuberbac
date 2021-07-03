package main

import (
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"k8s.io/cli-runtime/pkg/genericclioptions"
)

func newRootCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:          "kuberbac",
		Short:        "Quick create or delete of kubernetes RBAC",
		SilenceUsage: true,
	}

	flags := cmd.PersistentFlags()

	kubeConfigFlags := genericclioptions.NewConfigFlags(true).WithDeprecatedPasswordFlag()
	addKubeConfigFlags(kubeConfigFlags, flags)

	cmd.AddCommand(
		newCreateCmd(kubeConfigFlags),
		newDeleteCommand(kubeConfigFlags),
		newCompletionCommand(),
	)
	return cmd
}

func addKubeConfigFlags(f *genericclioptions.ConfigFlags, flags *pflag.FlagSet) {
	if f.Namespace != nil {
		flags.StringVarP(f.Namespace, "namespace", "n", *f.Namespace, "If present, the namespace scope for this CLI request")
	}
	if f.KubeConfig != nil {
		flags.StringVar(f.KubeConfig, "kubeconfig", *f.KubeConfig, "Path to the kubeconfig file to use for CLI requests.")
	}
	if f.Context != nil {
		flags.StringVar(f.Context, "kube-context", *f.Context, "Path to the kubeconfig file to use for CLI requests.")
	}
}