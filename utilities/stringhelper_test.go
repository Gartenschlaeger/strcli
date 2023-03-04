package utilities_test

import (
	"testing"

	"github.com/Gartenschlaeger/strcli/utilities"
	"github.com/stretchr/testify/assert"
)

type stringShiftTestCase struct {
	input       string
	amount      int
	placeholder rune
	repeat      bool
	expected    string
}

func TestStringShift(t *testing.T) {
	testCases := []stringShiftTestCase{
		{
			input:       "",
			amount:      1,
			repeat:      true,
			placeholder: ' ',
			expected:    "",
		},
		{
			input:       "",
			amount:      1,
			repeat:      false,
			placeholder: ' ',
			expected:    " ",
		},
		{
			input:       "1234",
			amount:      1,
			repeat:      true,
			placeholder: ' ',
			expected:    "4123",
		},
		{
			input:       "1234",
			amount:      3,
			repeat:      true,
			placeholder: ' ',
			expected:    "2341",
		},
		{
			input:       "1234",
			amount:      6,
			repeat:      true,
			placeholder: ' ',
			expected:    "3412",
		},
		{
			input:       "1234",
			amount:      -1,
			repeat:      true,
			placeholder: ' ',
			expected:    "2341",
		},
		{
			input:       "1234",
			amount:      -3,
			repeat:      true,
			placeholder: ' ',
			expected:    "4123",
		},
		{
			input:       "1234",
			amount:      1,
			repeat:      false,
			placeholder: '_',
			expected:    "_123",
		},
		{
			input:       "1234",
			amount:      3,
			repeat:      false,
			placeholder: '_',
			expected:    "___1",
		},
		{
			input:       "1234",
			amount:      -1,
			repeat:      false,
			placeholder: '_',
			expected:    "234_",
		},
		{
			input:       "1234",
			amount:      -3,
			repeat:      false,
			placeholder: '_',
			expected:    "4___",
		},
	}

	for i, tc := range testCases {
		t.Run(utilities.PadInt(i, 2), func(t *testing.T) {
			r := utilities.ShiftString(tc.input, tc.amount, tc.placeholder, tc.repeat)

			assert.Equal(t, tc.expected, r)
		})
	}
}
