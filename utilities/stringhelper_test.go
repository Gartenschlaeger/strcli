package utilities_test

import (
	"testing"

	"github.com/Gartenschlaeger/strcli/utilities"
	"github.com/stretchr/testify/assert"
)

type padStringTestCase struct {
	input    string
	prefix   rune
	length   int
	expected string
}

func TestPadString(t *testing.T) {
	testCases := []padStringTestCase{
		{"1", '0', 5, "00001"},
		{"10", '0', 3, "010"},
		{"10", '.', 10, "........10"},
	}

	for i, tc := range testCases {
		t.Run(utilities.PadInt(i, 2), func(t *testing.T) {
			r := utilities.PadString(tc.input, tc.prefix, tc.length)

			assert.EqualValues(t, tc.expected, r)
		})
	}
}

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
