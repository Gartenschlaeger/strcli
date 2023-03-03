package commands

import (
	"github.com/spf13/cobra"
)

func ReverseCommandHandler(ctx *CommandContext) {
	runes := []rune(ctx.Input)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	ctx.Result = string(runes)
}

func NewReverseCommand(ctx *CommandContext) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "reverse",
		Short: "Reverses the order of all characters",
		Run: func(cmd *cobra.Command, args []string) {
			ReverseCommandHandler(ctx)
		},
	}

	flags := cmd.Flags()
	flags.SetInterspersed(false)

	return cmd
}
