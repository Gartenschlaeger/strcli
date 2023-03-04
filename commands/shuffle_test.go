package commands_test

import (
	"strings"
	"testing"

	"github.com/Gartenschlaeger/strcli/commands"
	"github.com/stretchr/testify/assert"
)

func TestShuffleCommand(t *testing.T) {
	t.Run("should not fail with empty input", func(t *testing.T) {
		input := ""

		ctx := commands.NewCommandContext(input)
		commands.ShuffleCommandHandler(ctx)

		assert.Equal(t, "", "")
	})

	t.Run("should have the same length", func(t *testing.T) {
		input := "1234567890"

		ctx := commands.NewCommandContext(input)
		commands.ShuffleCommandHandler(ctx)

		assert.Equal(t, len(input), len(ctx.Result))
	})

	t.Run("should have the same characters", func(t *testing.T) {
		input := "1234567890"

		ctx := commands.NewCommandContext(input)
		commands.ShuffleCommandHandler(ctx)

		for _, r := range input {
			assert.True(t, strings.ContainsRune(ctx.Result, r))
		}
	})
}
