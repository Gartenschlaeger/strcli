package commands

import (
	"github.com/Gartenschlaeger/strcli/utilities"
	"github.com/spf13/cobra"
)

func Md5CommandHandler(ctx *CommandContext) error {
	ctx.Result = utilities.GetMd5Hash(ctx.Input)

	return nil
}

func NewMd5Command(ctx *CommandContext) *CommandConfiguration {
	cmd := &CommandConfiguration{
		name:        "md5",
		description: "Calculates a MD5 hash",
		handler: func(cmd *cobra.Command, args []string) error {
			return Md5CommandHandler(ctx)
		},
	}

	return cmd
}
