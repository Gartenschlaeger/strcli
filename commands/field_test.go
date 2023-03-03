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
		{"index 0 returns first field", "Max Mustermann", 0, " ", false, "Max"},
		{"index 1 returns second field", "Max Mustermann", 1, " ", false, "Mustermann"},
		{"index -1 returns last field", "Max Mustermann", -1, " ", false, "Mustermann"},
		{"index -5 returns first field", "Max Mustermann", -5, " ", false, "Max"},
		{"index overflow returns last field", "Max Mustermann", 10, " ", false, "Mustermann"},
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
