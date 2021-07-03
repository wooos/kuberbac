package main

import "os"

func main() {
	command := newRootCmd()
	if err := command.Execute(); err != nil {
		os.Exit(1)
	}
}
