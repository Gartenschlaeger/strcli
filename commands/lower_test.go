package commands_test

import (
	"strconv"
	"testing"

	"github.com/Gartenschlaeger/strcli/commands"
	"github.com/stretchr/testify/assert"
)

type lowerTestCase struct {
	input    string
	expected string
}

func TestLowerCommand(t *testing.T) {
	testCases := []lowerTestCase{
		{"This is a test", "this is a test"},
		{"JOHN", "john"},
		{"test", "test"},
	}

	for i, tc := range testCases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			ctx := commands.NewContext(tc.input)

			commands.LowerCommandHandler(ctx)

			assert.EqualValues(t, tc.expected, ctx.Result)
		})
	}
}
