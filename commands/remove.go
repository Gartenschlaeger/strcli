package commands

import (
	"github.com/Gartenschlaeger/strcli/utilities"
	"github.com/spf13/cobra"
)

type RemoveCommandOptions struct {
	Index  int
	Length int
}

func RemoveCommandHandler(ctx *CommandContext, opt *RemoveCommandOptions) {
	startIndex, endIndex := utilities.ClampStringPartion(ctx.Input, opt.Index, opt.Length)

	ctx.Result = ctx.Input[0:startIndex] + ctx.Input[endIndex:]
}

func NewRemoveCommand(ctx *CommandContext) *cobra.Command {
	opt := RemoveCommandOptions{}

	cmd := &cobra.Command{
		Use:   "remove",
		Short: "Removes a partition",
		Run: func(cmd *cobra.Command, args []string) {
			RemoveCommandHandler(ctx, &opt)
		},
	}

	flags := cmd.Flags()
	flags.SetInterspersed(false)

	flags.IntVarP(&opt.Index, "index", "i", 0, "Zero based index of the first character")
	flags.IntVarP(&opt.Length, "length", "l", 1, "Number of characters")

	return cmd
}
