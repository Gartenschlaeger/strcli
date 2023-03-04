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
	testCases := []shiftTestCase{
		{
			input:       "",
			amount:      1,
			repeat:      true,
			placeholder: " ",
			expected:    "",
		},
		{
			input:       "",
			amount:      1,
			repeat:      false,
			placeholder: " ",
			expected:    " ",
		},
		{
			input:       "1234",
			amount:      1,
			repeat:      true,
			placeholder: " ",
			expected:    "4123",
		},
		{
			input:       "1234",
			amount:      3,
			repeat:      true,
			placeholder: " ",
			expected:    "2341",
		},
		{
			input:       "1234",
			amount:      6,
			repeat:      true,
			placeholder: " ",
			expected:    "3412",
		},
		{
			input:       "1234",
			amount:      -1,
			repeat:      true,
			placeholder: " ",
			expected:    "2341",
		},
		{
			input:       "1234",
			amount:      -3,
			repeat:      true,
			placeholder: " ",
			expected:    "4123",
		},
		{
			input:       "1234",
			amount:      1,
			repeat:      false,
			placeholder: "_",
			expected:    "_123",
		},
		{
			input:       "1234",
			amount:      3,
			repeat:      false,
			placeholder: "_",
			expected:    "___1",
		},
		{
			input:       "1234",
			amount:      -1,
			repeat:      false,
			placeholder: "_",
			expected:    "234_",
		},
		{
			input:       "1234",
			amount:      -3,
			repeat:      false,
			placeholder: "_",
			expected:    "4___",
		},
	}

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
