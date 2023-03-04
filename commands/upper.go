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
		Name:             "upper",
		Description:      "Converts characters to upper case",
		HasSelectionFlag: true,
		Handler: func(cmd *cobra.Command, args []string) error {
			return UpperCommandHandler(ctx)
		},
	}

	return cmd
}
