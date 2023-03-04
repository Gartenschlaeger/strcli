package commands_test

import (
	"strconv"
	"testing"

	"github.com/Gartenschlaeger/strcli/commands"
	"github.com/stretchr/testify/assert"
)

type subTestCase struct {
	input    string
	index    int
	length   int
	expected string
}

func TestSubCommand(t *testing.T) {
	testCases := []subTestCase{
		{"", 0, 1, ""},
		{"Some text", 0, 4, "Some"},
		{"Some text", 5, 5, "text"},
		{"Some text", -4, 4, "text"},
		{"Some text", -4, 3, "tex"},
	}

	for i, tc := range testCases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			ctx := commands.NewCommandContext(tc.input)

			opt := commands.SubCommandOptions{
				Index:  tc.index,
				Length: tc.length,
			}

			commands.SubCommandHandler(ctx, &opt)

			assert.Equal(t, tc.expected, ctx.Result)
		})
	}
}
