package commands

import (
	"github.com/Gartenschlaeger/strcli/utilities"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

type SubCommandOptions struct {
	Index  int
	Length int
}

func SubCommandHandler(ctx *CommandContext, opt *SubCommandOptions) error {
	startIndex, endIndex := utilities.ClampStringPartion(ctx.Input, opt.Index, opt.Length)

	ctx.Result = ctx.Input[startIndex:endIndex]

	return nil
}

func NewSubCommand(context *CommandContext) *CommandConfiguration {
	opt := SubCommandOptions{}

	cmd := &CommandConfiguration{
		name:        "sub",
		description: "Returns a partition",
		handler: func(cmd *cobra.Command, args []string) error {
			return SubCommandHandler(context, &opt)
		},
		setup: func(cmd *cobra.Command, flags *pflag.FlagSet) {
			flags.IntVarP(&opt.Index, "index", "i", 0, "Zero based index of the first character")
			flags.IntVarP(&opt.Length, "length", "l", 1, "Number of characters")
		},
	}

	return cmd
}
