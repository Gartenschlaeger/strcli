package commands_test

import (
	"strconv"
	"testing"

	"github.com/Gartenschlaeger/strcli/commands"
	"github.com/stretchr/testify/assert"
)

type upperTestCase struct {
	input    string
	rangeVal string
	expected string
}

func TestUpperCommand(t *testing.T) {
	testCases := []upperTestCase{
		{"This is a test", "", "THIS IS A TEST"},
		{"JOHN", "", "JOHN"},
		{"test", "", "TEST"},
		{"test", "2", "TEst"},
		{"test", "-3", "tEST"},
		{"Hello world!", "6:5", "Hello WORLD!"},
	}

	for i, tc := range testCases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			ctx := commands.NewContext(tc.input)

			opt := commands.UpperCommandOptions{
				Select: tc.rangeVal,
			}

			commands.UpperCommandHandler(ctx, &opt)

			assert.EqualValues(t, tc.expected, ctx.Result)
		})
	}
}
