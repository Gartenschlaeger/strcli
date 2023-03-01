package main

import (
	"fmt"
	"os"

	"github.com/Gartenschlaeger/strcli/commands"
	"github.com/spf13/cobra"
)

func main() {

	rootCmd := &cobra.Command{
		Use:   "str COMMAND [ARG...]",
		Short: "Runs common string operations.",
	}

	commands.DefineFieldCommand(rootCmd)
	commands.DefineSubCommand(rootCmd)

	err := rootCmd.Execute()
	if err != nil {
		fmt.Fprint(os.Stderr, err.Error())
		os.Exit(1)
	}

}
