package commands

import (
	"strings"

	"github.com/spf13/cobra"
)

func LowerCommandHandler(ctx *CommandContext) error {
	return ProcessResult(ctx, func(input string) (string, error) {
		return strings.ToLower(input), nil
	})
}

func NewLowerCommand(ctx *CommandContext) *CommandConfiguration {
	cmd := &CommandConfiguration{
		name:             "lower",
		description:      "Converts all characters to lower case",
		hasSelectionFlag: true,
		handler: func(cmd *cobra.Command, args []string) error {
			return LowerCommandHandler(ctx)
		},
	}

	return cmd
}
