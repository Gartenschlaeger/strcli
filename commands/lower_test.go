package commands_test

import (
	"testing"

	"github.com/Gartenschlaeger/strcli/commands"
	"github.com/stretchr/testify/assert"
)

type lowerTestCase struct {
	testName       string
	input          string
	expectedResult string
}

func TestLowerCommand(t *testing.T) {
	testCases := []lowerTestCase{
		{"mixed to lower case", "This is a test", "this is a test"},
		{"upper to lower case", "JOHN", "john"},
		{"lower to lower case", "test", "test"},
	}

	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			ctx := commands.NewContext(tc.input)

			commands.LowerCommandHandler(ctx)

			assert.EqualValues(t, tc.expectedResult, ctx.Result)
		})
	}
}
