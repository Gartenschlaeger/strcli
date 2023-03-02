package main

import (
	"fmt"
	"os"

	"github.com/Gartenschlaeger/strcli/commands"
	"github.com/spf13/cobra"
)

func main() {
	cmd := &cobra.Command{
		Use:   "str",
		Short: "Performs general string operations",
	}

	commands.AddFieldCommand(cmd)
	commands.AddSubCommand(cmd)

	err := cmd.Execute()
	if err != nil {
		fmt.Fprint(os.Stderr, err.Error())
		os.Exit(1)
	}
}
