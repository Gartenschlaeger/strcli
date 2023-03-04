package commands_test

import (
	"testing"

	"github.com/Gartenschlaeger/strcli/commands"
	"github.com/Gartenschlaeger/strcli/utilities"
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
		{"", 0, " ", false, ""},
		{"", 0, " ", true, ""},
		{"Max Mustermann", 0, " ", false, "Max"},
		{"Max Mustermann", 1, " ", false, "Mustermann"},
		{"Max Mustermann", -1, " ", false, "Mustermann"},
		{"Max Mustermann", -5, " ", false, "Max"},
		{"Max Mustermann", 10, " ", false, "Mustermann"},
	}

	for i, tc := range testCases {
		t.Run(utilities.PadInt(i, 2), func(t *testing.T) {
			ctx := commands.NewCommandContext(tc.input)

			opt := commands.FieldCommandOptions{
				Index:       tc.index,
				Separator:   tc.separator,
				IgnoreEmpty: tc.ignoreEmpty,
			}

			err := commands.FieldCommandHandler(ctx, &opt)

			assert.NoError(t, err)
			assert.EqualValues(t, tc.expected, ctx.Result)
		})
	}
}
