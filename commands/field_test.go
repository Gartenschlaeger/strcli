package commands_test

import (
	"strconv"
	"testing"

	"github.com/Gartenschlaeger/strcli/commands"
	"github.com/stretchr/testify/assert"
)

type fieldTestCase struct {
	input       string
	index       int
	separator   string
	ignoreEmpty bool
	expected    string
}

func TestFieldCommand(t *testing.T) {
	testCases := []fieldTestCase{
		{"Max Mustermann", 0, " ", false, "Max"},
		{"Max Mustermann", 1, " ", false, "Mustermann"},
		{"Max Mustermann", -1, " ", false, "Mustermann"},
		{"Max Mustermann", -5, " ", false, "Max"},
		{"Max Mustermann", 10, " ", false, "Mustermann"},
	}

	for i, tc := range testCases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			ctx := commands.NewContext(tc.input)

			opt := commands.FieldCommandOptions{
				Index:       tc.index,
				Separator:   tc.separator,
				IgnoreEmpty: tc.ignoreEmpty,
			}

			commands.FieldCommandHandler(ctx, &opt)

			assert.EqualValues(t, tc.expected, ctx.Result)
		})
	}
}
