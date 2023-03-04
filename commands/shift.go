package commands

import (
	"github.com/Gartenschlaeger/strcli/utilities"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

type ShiftCommandOptions struct {
	Amount      int
	Repeat      bool
	Placeholder string
}

func ShiftCommandHandler(ctx *CommandContext, opt *ShiftCommandOptions) error {
	p := ' '
	if len(opt.Placeholder) > 0 {
		p = rune(opt.Placeholder[0])
	}

	ctx.Result = utilities.ShiftString(ctx.Input, opt.Amount, p, opt.Repeat)

	return nil
}

func NewShiftCommand(ctx *CommandContext) *CommandConfiguration {
	opt := &ShiftCommandOptions{}

	cmd := &CommandConfiguration{
		Name:        "shift",
		Description: "Shifts all characters",
		Handler: func(cmd *cobra.Command, args []string) error {
			return ShiftCommandHandler(ctx, opt)
		},
		Setup: func(cmd *cobra.Command, flags *pflag.FlagSet) {
			flags.IntVarP(&opt.Amount, "amount", "a", 1, "The shift amount")
			flags.BoolVarP(&opt.Repeat, "repeat", "r", true, "Characters are repeated")
			flags.StringVarP(&opt.Placeholder, "placeholder", "p", " ", "Placeholder character")
		},
	}

	return cmd
}
