package utilities_test

import (
	"testing"

	"github.com/Gartenschlaeger/strcli/utilities"
	"github.com/stretchr/testify/assert"
)

type parseIntTestCase struct {
	s        string
	d        int
	expected int
}

func TestParseInt(t *testing.T) {
	testCases := []parseIntTestCase{
		{"0", 0, 0},
		{"1", 0, 1},
		{"_", 0, 0},
	}

	for i, tc := range testCases {
		t.Run(utilities.PadInt(i, 2), func(t *testing.T) {
			r := utilities.ParseInt(tc.s, tc.d)

			assert.Equal(t, tc.expected, r)
		})
	}
}
