package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

func NewCompletionCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "completion",
		Short: "Generate auto completions script for kuberbac for the specified shell (bash)",
		RunE:  runCompletionCommand,
	}

	return cmd
}

func runCompletionCommand(cmd *cobra.Command, args []string) error {
	err := cmd.Root().GenBashCompletion(os.Stdout)
	bashrc := `
if [[ $(type -t compopt) = "builtin"  ]]; then
	complete -o default -F __start_kuberbac kuberbac
else
	complete -o default -o nospace -F __start_kuberbac kuberbac
fi
`
	fmt.Printf("%s", bashrc)
	return err
}
