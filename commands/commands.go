package commands

import (
	"fmt"
	"os"

	"github.com/Gartenschlaeger/strcli/utilities"
	"github.com/spf13/cobra"
)

type CommandContext struct {
	Input  string
	Result string
}

func NewContext(input string) *CommandContext {
	return &CommandContext{
		Input:  input,
		Result: "",
	}
}

var rootCmd *cobra.Command

func Execute() {
	input := utilities.GetStandardInputString()

	ctx := NewContext(input)

	rootCmd = &cobra.Command{
		Use:   "str",
		Short: "Performs general string operations",
		PersistentPostRun: func(cmd *cobra.Command, args []string) {
			fmt.Print(ctx.Result)
		},
	}

	rootCmd.Version = "1.3.0"

	rootCmd.AddCommand(
		NewFieldCommand(ctx),
		NewReplaceCommand(ctx),
		NewSubCommand(ctx),
		NewRemoveCommand(ctx),
		NewLowerCommand(ctx),
		NewUpperCommand(ctx),
		NewTrimCommand(ctx),
		NewMd5Command(ctx),
		NewSha1Command(ctx),
	)

	err := rootCmd.Execute()
	if err != nil {
		fmt.Fprint(os.Stderr, err.Error())
		os.Exit(1)
	}
}
