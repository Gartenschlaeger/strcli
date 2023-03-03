package commands

import (
	"strings"

	"github.com/spf13/cobra"
)

func LowerCommandHandler(ctx *CommandContext) {
	ctx.Result = strings.ToLower(ctx.Input)
}

func NewLowerCommand(ctx *CommandContext) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "lower",
		Short: "Converts all characters to lower case",
		Run: func(cmd *cobra.Command, args []string) {
			LowerCommandHandler(ctx)
		},
	}

	flags := cmd.Flags()
	flags.SetInterspersed(false)

	return cmd
}
