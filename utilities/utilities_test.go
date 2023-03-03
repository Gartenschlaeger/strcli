package utilities_test

import (
	"fmt"
	"testing"

	"github.com/Gartenschlaeger/strcli/utilities"
	"github.com/stretchr/testify/assert"
)

type clampITestCase struct {
	value          int
	min            int
	max            int
	expectedResult int
}

func TestClampI(t *testing.T) {
	testCases := []clampITestCase{
		{10, 1, 3, 3},
		{0, 1, 3, 1},
		{2, 1, 3, 2},
	}

	for _, tc := range testCases {
		testName := fmt.Sprintf("%v = %v - %v = %v", tc.value, tc.min, tc.max, tc.expectedResult)

		t.Run(testName, func(t *testing.T) {
			value := utilities.ClampI(tc.value, tc.min, tc.max)

			assert.Equal(t, tc.expectedResult, value)
		})
	}
}
