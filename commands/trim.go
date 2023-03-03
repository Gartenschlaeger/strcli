package commands

import (
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

type TrimCommandOptions struct {
	Cutset string
}

func TrimCommandHandler(ctx *CommandContext, opt *TrimCommandOptions) {
	ctx.Result = strings.Trim(ctx.Input, opt.Cutset)
}

func NewTrimCommand(ctx *CommandContext) *CommandConfiguration {
	opt := TrimCommandOptions{}

	cmd := &CommandConfiguration{
		name:        "trim",
		description: "Removes all leading and trailing characters from a set of specified characters",
		handler: func(cmd *cobra.Command, args []string) error {
			TrimCommandHandler(ctx, &opt)

			return nil
		},
		setup: func(cmd *cobra.Command, flags *pflag.FlagSet) {
			flags.StringVarP(&opt.Cutset, "cutset", "c", "\n\r\t ", "Set of characters to be removed")
		},
	}

	return cmd
}
