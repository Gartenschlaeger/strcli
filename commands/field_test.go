package commands_test

import (
	"testing"

	"github.com/Gartenschlaeger/strcli/commands"
	"github.com/stretchr/testify/assert"
)

type fieldTestCase struct {
	testName       string
	input          string
	index          int
	separator      string
	ignoreEmpty    bool
	expectedResult string
}

func TestFieldCommand(t *testing.T) {
	testCases := []fieldTestCase{
		{"index 0 returns first field", "Max Mustermann", 0, " ", false, "Max"},
		{"index 1 returns second field", "Max Mustermann", 1, " ", false, "Mustermann"},
		{"index -1 returns last field", "Max Mustermann", -1, " ", false, "Mustermann"},
		{"index -5 returns first field", "Max Mustermann", -5, " ", false, "Max"},
		{"index overflow returns last field", "Max Mustermann", 10, " ", false, "Mustermann"},
	}

	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			ctx := commands.NewContext(tc.input)

			opt := commands.FieldCommandOptions{
				Index:       tc.index,
				Separator:   tc.separator,
				IgnoreEmpty: tc.ignoreEmpty,
			}

			commands.FieldCommandHandler(ctx, &opt)

			assert.EqualValues(t, tc.expectedResult, ctx.Result)
		})
	}
}
