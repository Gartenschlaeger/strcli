package commands_test

import (
	"testing"

	"github.com/Gartenschlaeger/strcli/commands"
	"github.com/stretchr/testify/assert"
)

type upperTestCase struct {
	testName       string
	input          string
	expectedResult string
}

func TestUpperCommand(t *testing.T) {
	testCases := []upperTestCase{
		{"mixed to upper case", "This is a test", "THIS IS A TEST"},
		{"upper to upper case", "JOHN", "JOHN"},
		{"lower to upper case", "test", "TEST"},
	}

	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			ctx := commands.NewContext(tc.input)

			commands.UpperCommandHandler(ctx)

			assert.EqualValues(t, tc.expectedResult, ctx.Result)
		})
	}
}
