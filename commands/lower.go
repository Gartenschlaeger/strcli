package commands

import (
	"strings"

	"github.com/spf13/cobra"
)

func LowerCommandHandler(ctx *CommandContext) {
	ProcessResult(ctx, func(input string) string {
		return strings.ToLower(input)
	})
}

func NewLowerCommand(ctx *CommandContext) *CommandConfiguration {
	cmd := &CommandConfiguration{
		name:             "lower",
		description:      "Converts all characters to lower case",
		hasSelectionFlag: true,
		handler: func(cmd *cobra.Command, args []string) {
			LowerCommandHandler(ctx)
		},
	}

	return cmd
}
