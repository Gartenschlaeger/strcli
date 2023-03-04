package commands

import (
	"github.com/Gartenschlaeger/strcli/utilities"
	"github.com/spf13/cobra"
)

func ShuffleCommandHandler(ctx *CommandContext) {
	ctx.Result = utilities.ShuffleString(ctx.Input)
}

func NewShuffleCommand(ctx *CommandContext) *CommandConfiguration {
	cmd := &CommandConfiguration{
		name:        "shuffle",
		description: "Shuffles the individual characters randomly",
		handler: func(cmd *cobra.Command, args []string) error {
			ShuffleCommandHandler(ctx)

			return nil
		},
	}

	return cmd
}
