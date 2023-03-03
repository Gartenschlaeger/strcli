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
	startIndex, endIndex := utilities.ClampStringPartion(ctx.Input, opt.Index, opt.Length)

	ctx.Result = ctx.Input[startIndex:endIndex]
}

func NewSubCommand(context *CommandContext) *cobra.Command {
	opt := SubCommandOptions{}

	cmd := &cobra.Command{
		Use:   "sub",
		Short: "Returns a partition",
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
