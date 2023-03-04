package commands

import (
	"strings"

	"github.com/spf13/cobra"
)

func UpperCommandHandler(ctx *CommandContext) error {
	return ProcessResult(ctx, func(input string) (string, error) {
		return strings.ToUpper(input), nil
	})
}

func NewUpperCommand(ctx *CommandContext) *CommandConfiguration {
	cmd := &CommandConfiguration{
		name:             "upper",
		description:      "Converts characters to upper case",
		hasSelectionFlag: true,
		handler: func(cmd *cobra.Command, args []string) error {
			return UpperCommandHandler(ctx)
		},
	}

	return cmd
}
