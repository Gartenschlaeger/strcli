package commands

import (
	"strings"

	"github.com/spf13/cobra"
)

type UpperCommandOptions struct {
	Select string
}

func UpperCommandHandler(ctx *CommandContext, opt *UpperCommandOptions) {
	ManipulateSelection(ctx, opt.Select, func(input string) string {
		return strings.ToUpper(input)
	})
}

func NewUpperCommand(ctx *CommandContext) *cobra.Command {
	opt := UpperCommandOptions{}

	cmd := &cobra.Command{
		Use:   "upper",
		Short: "Converts all characters to upper case",
		Run: func(cmd *cobra.Command, args []string) {
			UpperCommandHandler(ctx, &opt)
		},
	}

	flags := cmd.Flags()
	flags.SetInterspersed(false)

	flags.StringVarP(&opt.Select, "select", "s", "", "Selection of characters")

	return cmd
}
