package commands

import (
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

type TrimCommandOptions struct {
	Cutset string
}

func TrimCommandHandler(ctx *CommandContext, opt *TrimCommandOptions) error {
	ctx.Result = strings.Trim(ctx.Input, opt.Cutset)

	return nil
}

func NewTrimCommand(ctx *CommandContext) *CommandConfiguration {
	opt := TrimCommandOptions{}

	cmd := &CommandConfiguration{
		Name:        "trim",
		Description: "Removes all leading and trailing characters from a set of specified characters",
		Handler: func(cmd *cobra.Command, args []string) error {
			return TrimCommandHandler(ctx, &opt)
		},
		Setup: func(cmd *cobra.Command, flags *pflag.FlagSet) {
			flags.StringVarP(&opt.Cutset, "cutset", "c", "\n\r\t ", "Set of characters to be removed")
		},
	}

	return cmd
}
