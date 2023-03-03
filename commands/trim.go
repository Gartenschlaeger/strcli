package commands

import (
	"strings"

	"github.com/spf13/cobra"
)

type TrimCommandOptions struct {
	Cutset string
}

func TrimCommandHandler(ctx *CommandContext, opt *TrimCommandOptions) {
	ctx.Result = strings.Trim(ctx.Input, opt.Cutset)
}

func NewTrimCommand(ctx *CommandContext) *cobra.Command {
	opt := TrimCommandOptions{}

	cmd := &cobra.Command{
		Use:   "trim",
		Short: "Removes all leading and trailing characters from a set of specified characters",
		Run: func(cmd *cobra.Command, args []string) {
			TrimCommandHandler(ctx, &opt)
		},
	}

	flags := cmd.Flags()
	flags.SetInterspersed(false)

	flags.StringVarP(&opt.Cutset, "cutset", "c", "\n\r\t ", "Set of characters to be removed")

	return cmd
}
