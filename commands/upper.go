package commands

import (
	"strings"

	"github.com/spf13/cobra"
)

func UpperCommandHandler(ctx *CommandContext) {
	ctx.Result = strings.ToUpper(ctx.Input)
}

func NewUpperCommand(ctx *CommandContext) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "upper",
		Short: "Converts all characters to upper case",
		Run: func(cmd *cobra.Command, args []string) {
			UpperCommandHandler(ctx)
		},
	}

	flags := cmd.Flags()
	flags.SetInterspersed(false)

	return cmd
}
