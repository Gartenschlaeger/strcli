package commands

import (
	"github.com/Gartenschlaeger/strcli/utilities"
	"github.com/spf13/cobra"
)

func ShuffleCommandHandler(ctx *CommandContext) error {
	ctx.Result = utilities.ShuffleString(ctx.Input)

	return nil
}

func NewShuffleCommand(ctx *CommandContext) *CommandConfiguration {
	cmd := &CommandConfiguration{
		Name:        "shuffle",
		Description: "Shuffles the individual characters randomly",
		Handler: func(cmd *cobra.Command, args []string) error {
			return ShuffleCommandHandler(ctx)
		},
	}

	return cmd
}
