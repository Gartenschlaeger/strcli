package commands

import (
	"strings"

	"github.com/spf13/cobra"
)

func UpperCommandHandler(ctx *CommandContext) {
	ProcessResult(ctx, func(input string) string {
		return strings.ToUpper(input)
	})
}

func NewUpperCommand(ctx *CommandContext) *CommandConfiguration {
	cmd := &CommandConfiguration{
		name:             "upper",
		description:      "Converts characters to upper case",
		hasSelectionFlag: true,
		handler: func(cmd *cobra.Command, args []string) {
			UpperCommandHandler(ctx)
		},
	}

	return cmd
}
