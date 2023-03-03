package commands_test

import (
	"strconv"
	"testing"

	"github.com/Gartenschlaeger/strcli/commands"
	"github.com/stretchr/testify/assert"
)

type removeTestCase struct {
	input    string
	index    int
	length   int
	expected string
}

func TestRemoveCommand(t *testing.T) {
	testCases := []removeTestCase{
		{"Hello World", 0, 5, " World"},
		{"Hello World", 5, 6, "Hello"},
		{"Hello World", 5, 99, "Hello"},
		{"1234567890", -3, 3, "1234567"},
		{"1234567890", 0, 99, ""},
	}

	for i, tc := range testCases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			ctx := commands.NewCommandContext(tc.input)

			opt := commands.RemoveCommandOptions{
				Index:  tc.index,
				Length: tc.length,
			}

			commands.RemoveCommandHandler(ctx, &opt)

			assert.Equal(t, tc.expected, ctx.Result)
		})
	}
}
