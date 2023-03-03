package commands_test

import (
	"strconv"
	"testing"

	"github.com/Gartenschlaeger/strcli/commands"
	"github.com/stretchr/testify/assert"
)

type upperTestCase struct {
	input    string
	expected string
}

func TestUpperCommand(t *testing.T) {
	testCases := []upperTestCase{
		{"This is a test", "THIS IS A TEST"},
		{"JOHN", "JOHN"},
		{"test", "TEST"},
	}

	for i, tc := range testCases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			ctx := commands.NewContext(tc.input)

			commands.UpperCommandHandler(ctx)

			assert.EqualValues(t, tc.expected, ctx.Result)
		})
	}
}
