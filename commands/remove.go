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

func RemoveCommandHandler(ctx *CommandContext, opt *RemoveCommandOptions) error {
	startIndex, endIndex := utilities.ClampStringPartion(ctx.Input, opt.Index, opt.Length)

	ctx.Result = ctx.Input[0:startIndex] + ctx.Input[endIndex:]

	return nil
}

func NewRemoveCommand(ctx *CommandContext) *CommandConfiguration {
	opt := RemoveCommandOptions{}

	cmd := &CommandConfiguration{
		Name:        "remove",
		Description: "Removes a partition",
		Handler: func(cmd *cobra.Command, args []string) error {
			return RemoveCommandHandler(ctx, &opt)
		},
		Setup: func(cmd *cobra.Command, flags *pflag.FlagSet) {
			flags.IntVarP(&opt.Index, "index", "i", 0, "Zero based index of the first character")
			flags.IntVarP(&opt.Length, "length", "l", 1, "Number of characters")
		},
	}

	return cmd
}
