package commands_test

import (
	"testing"

	"github.com/Gartenschlaeger/strcli/commands"
	"github.com/Gartenschlaeger/strcli/utilities"
	"github.com/stretchr/testify/assert"
)

type shiftTestCase struct {
	input       string
	amount      int
	repeat      bool
	placeholder string
	expected    string
}

func TestShiftCommand(t *testing.T) {
	// all test cases are already covered by TestStringShift()
	testCases := []shiftTestCase{}

	for i, tc := range testCases {
		t.Run(utilities.PadInt(i, 2), func(t *testing.T) {
			ctx := commands.NewCommandContext(tc.input)

			opt := &commands.ShiftCommandOptions{
				Amount:      tc.amount,
				Repeat:      tc.repeat,
				Placeholder: tc.placeholder,
			}

			err := commands.ShiftCommandHandler(ctx, opt)

			assert.NoError(t, err)
			assert.EqualValues(t, tc.expected, ctx.Result)
		})
	}
}
