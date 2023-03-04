package commands

import (
	"github.com/Gartenschlaeger/strcli/utilities"
	"github.com/spf13/cobra"
)

func ShaCommandHandler(ctx *CommandContext) error {
	ctx.Result = utilities.GetSha1Hash(ctx.Input)

	return nil
}

func NewSha1Command(ctx *CommandContext) *CommandConfiguration {
	cmd := &CommandConfiguration{
		name:        "sha1",
		description: "Calculates a SHA1 hash",
		handler: func(cmd *cobra.Command, args []string) error {
			return ShaCommandHandler(ctx)
		},
	}

	return cmd
}
