package commands_test

import (
	"strconv"
	"testing"

	"github.com/Gartenschlaeger/strcli/commands"
	"github.com/stretchr/testify/assert"
)

type trimTestCase struct {
	input    string
	cutset   string
	expected string
}

func TestTrimCommand(t *testing.T) {
	testCases := []trimTestCase{
		{"!!Test!!!", "!", "Test"},
		{"  ..   Test .  ", ".", "  ..   Test .  "},
		{"..   Test .", ".", "   Test "},
		{"¡¡¡Hello, Gophers!!!", "!¡", "Hello, Gophers"},
		{"  \nTest..", ".", "  \nTest"},
	}

	for i, tc := range testCases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			ctx := commands.NewCommandContext(tc.input)

			opt := commands.TrimCommandOptions{
				Cutset: tc.cutset,
			}

			commands.TrimCommandHandler(ctx, &opt)

			assert.Equal(t, tc.expected, ctx.Result)
		})
	}
}
