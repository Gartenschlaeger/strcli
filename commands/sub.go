package commands

import (
	"github.com/Gartenschlaeger/strcli/utilities"
	"github.com/spf13/cobra"
)

type SubCommandOptions struct {
	Index  int
	Length int
}

func SubCommandHandler(ctx *CommandContext, opt *SubCommandOptions) {
	length := len(ctx.Input)

	startIndex := utilities.ClampI(opt.Index, 0, length)
	endIndex := utilities.ClampI(opt.Index+opt.Length, startIndex, length)

	ctx.Result = ctx.Input[startIndex:endIndex]
}

func NewSubCommand(context *CommandContext) *cobra.Command {
	opt := SubCommandOptions{}

	cmd := &cobra.Command{
		Use:   "sub",
		Short: "Returns a partition beginning at index",
		Run: func(cmd *cobra.Command, args []string) {
			SubCommandHandler(context, &opt)
		},
	}

	flags := cmd.Flags()
	flags.SetInterspersed(false)

	flags.IntVarP(&opt.Index, "index", "i", 0, "Zero based index of the first character")
	flags.IntVarP(&opt.Length, "length", "l", 1, "Number of characters")

	return cmd
}
