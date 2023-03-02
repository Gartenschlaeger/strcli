package commands_test

import (
	"testing"

	"github.com/Gartenschlaeger/strcli/commands"
	"github.com/stretchr/testify/assert"
)

type fieldTestSet struct {
	testName       string
	input          string
	index          int
	separator      string
	ignoreEmpty    bool
	expectedResult string
}

func TestFieldCommand(t *testing.T) {
	tests := []fieldTestSet{
		{"should return first field for index 0", "Max Mustermann", 0, " ", false, "Max"},
		{"should return second field for index 1", "Max Mustermann", 1, " ", false, "Mustermann"},
		{"should return first field for negative index", "Max Mustermann", -3, " ", false, "Max"},
		{"should return last field for index overflow", "Max Mustermann", 10, " ", false, "Mustermann"},
	}

	for _, ts := range tests {
		t.Run(ts.testName, func(t *testing.T) {
			ctx := commands.NewContext(ts.input)

			opt := commands.FieldCommandOptions{
				Index:       ts.index,
				Separator:   ts.separator,
				IgnoreEmpty: ts.ignoreEmpty,
			}

			commands.FieldCommandHandler(ctx, &opt)

			assert.EqualValues(t, ts.expectedResult, ctx.Result)
		})
	}
}
