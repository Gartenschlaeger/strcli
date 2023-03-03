package commands

import (
	"math/rand"
	"time"

	"github.com/spf13/cobra"
)

func shuffleString(s string) string {
	rand.Seed(time.Now().UnixNano())

	runes := []rune(s)
	for i := len(runes) - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		runes[i], runes[j] = runes[j], runes[i]
	}

	return string(runes)
}

func ShuffleCommandHandler(ctx *CommandContext) {
	ctx.Result = shuffleString(ctx.Input)
}

func NewShuffleCommand(ctx *CommandContext) *CommandConfiguration {
	cmd := &CommandConfiguration{
		name:        "shuffle",
		description: "Shuffles the individual characters randomly",
		handler: func(cmd *cobra.Command, args []string) error {
			ShuffleCommandHandler(ctx)

			return nil
		},
	}

	return cmd
}
