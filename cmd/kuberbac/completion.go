package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var completionShells = map[string]func(cmd *cobra.Command){
	"bash": runCompletionBash,
	"zsh": runCompletionZsh,
}

func newCompletionCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "completion SHELL",
		Short: "Generate auto completions script for kuberbac for the specified shell (bash or zsh)",
		Run:  RunCompletionCommand,
	}

	return cmd
}

func RunCompletionCommand(cmd *cobra.Command, args []string) {
	if len(args) == 0 {
		fmt.Fprintf(os.Stderr, "Error: Shell not specified.\nSee 'kuberbac completion -h' for help and examples\n")
		os.Exit(1)
	}

	if len(args) > 1 {
		fmt.Fprintf(os.Stderr, "%s", "Too many arguments. Expected only the shell type.")
		os.Exit(1)
	}

	run, found := completionShells[args[0]]
	if !found {
		fmt.Fprintf(os.Stderr, "Unsupported shell type %q.", args[0])
		os.Exit(1)
	}

	run(cmd)
}

func runCompletionBash(cmd *cobra.Command) {
	cmd.GenBashCompletion(os.Stdout)
}

func runCompletionZsh(cmd *cobra.Command) {
	cmd.GenZshCompletion(os.Stderr)
}