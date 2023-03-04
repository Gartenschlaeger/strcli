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
		Name:        "sha1",
		Description: "Calculates a SHA1 hash",
		Handler: func(cmd *cobra.Command, args []string) error {
			return ShaCommandHandler(ctx)
		},
	}

	return cmd
}
