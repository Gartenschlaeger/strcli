package commands

import (
	"github.com/Gartenschlaeger/strcli/utilities"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

type RemoveCommandOptions struct {
	Index  int
	Length int
}

func RemoveCommandHandler(ctx *CommandContext, opt *RemoveCommandOptions) {
	startIndex, endIndex := utilities.ClampStringPartion(ctx.Input, opt.Index, opt.Length)

	ctx.Result = ctx.Input[0:startIndex] + ctx.Input[endIndex:]
}

func NewRemoveCommand(ctx *CommandContext) *CommandConfiguration {
	opt := RemoveCommandOptions{}

	cmd := &CommandConfiguration{
		name:        "remove",
		description: "Removes a partition",
		handler: func(cmd *cobra.Command, args []string) error {
			RemoveCommandHandler(ctx, &opt)

			return nil
		},
		setup: func(cmd *cobra.Command, flags *pflag.FlagSet) {
			flags.IntVarP(&opt.Index, "index", "i", 0, "Zero based index of the first character")
			flags.IntVarP(&opt.Length, "length", "l", 1, "Number of characters")
		},
	}

	return cmd
}
