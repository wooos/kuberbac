package main

import "kuberbac/cmd"

func main() {
	command := cmd.NewCommand()
	_ = command.Execute()
}
