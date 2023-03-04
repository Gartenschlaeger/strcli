package commands_test

import (
	"strconv"
	"testing"

	"github.com/Gartenschlaeger/strcli/commands"
	"github.com/stretchr/testify/assert"
)

type lowerTestCase struct {
	input     string
	selection string
	expected  string
}

func TestLowerCommand(t *testing.T) {
	testCases := []lowerTestCase{
		{"", "", ""},
		{"This is a test", "", "this is a test"},
		{"JOHN", "", "john"},
		{"test", "", "test"},
		{"HELLO WORLD", "5", "hello WORLD"},
	}

	for i, tc := range testCases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			ctx := commands.NewCommandContext(tc.input)
			ctx.Selection = tc.selection

			err := commands.LowerCommandHandler(ctx)

			assert.NoError(t, err)
			assert.EqualValues(t, tc.expected, ctx.Result)
		})
	}
}
