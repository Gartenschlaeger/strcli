package commands_test

import (
	"testing"

	"github.com/Gartenschlaeger/strcli/commands"
	"github.com/Gartenschlaeger/strcli/utilities"
	"github.com/stretchr/testify/assert"
)

type upperTestCase struct {
	input     string
	selection string
	expected  string
}

func TestUpperCommand(t *testing.T) {
	testCases := []upperTestCase{
		{"", "", ""},
		{"This is a test", "", "THIS IS A TEST"},
		{"JOHN", "", "JOHN"},
		{"test", "", "TEST"},
		{"test", "2", "TEst"},
		{"test", "-3", "tEST"},
		{"Hello world!", "6:5", "Hello WORLD!"},
	}

	for i, tc := range testCases {
		t.Run(utilities.PadInt(i, 2), func(t *testing.T) {
			ctx := commands.NewCommandContext(tc.input)
			ctx.Selection = tc.selection

			err := commands.UpperCommandHandler(ctx)

			assert.NoError(t, err)
			assert.EqualValues(t, tc.expected, ctx.Result)
		})
	}
}
