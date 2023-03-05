package utilities_test

import (
	"testing"

	"github.com/Gartenschlaeger/strcli/utilities"
	"github.com/stretchr/testify/assert"
)

type clampITestCase struct {
	v        int
	min      int
	max      int
	expected int
}

func TestClampI(t *testing.T) {
	testCases := []clampITestCase{
		{-1, 1, 3, 1},
		{0, 1, 3, 1},
		{2, 1, 3, 2},
		{3, 1, 3, 3},
		{4, 1, 3, 3},
	}

	for i, tc := range testCases {
		t.Run(utilities.PadInt(i, 2), func(t *testing.T) {
			r := utilities.ClampI(tc.v, tc.min, tc.max)

			assert.Equal(t, tc.expected, r)
		})
	}
}

type minITestCase struct {
	a        int
	b        int
	expected int
}

func TestMinI(t *testing.T) {
	testCases := []minITestCase{
		{1, 5, 1},
		{8, 3, 3},
		{-1, 1, -1},
	}

	for i, tc := range testCases {
		t.Run(utilities.PadInt(i, 2), func(t *testing.T) {
			r := utilities.MinI(tc.a, tc.b)

			assert.Equal(t, tc.expected, r)
		})
	}
}

type maxITestCase struct {
	a        int
	b        int
	expected int
}

func TestMaxI(t *testing.T) {
	testCases := []maxITestCase{
		{1, 5, 5},
		{8, 3, 8},
		{-1, 1, 1},
	}

	for i, tc := range testCases {
		t.Run(utilities.PadInt(i, 2), func(t *testing.T) {
			r := utilities.MaxI(tc.a, tc.b)

			assert.Equal(t, tc.expected, r)
		})
	}
}

type absITestCase struct {
	v        int
	expected int
}

func TestAbsI(t *testing.T) {
	testCases := []absITestCase{
		{-5, 5},
		{3, 3},
	}

	for i, tc := range testCases {
		t.Run(utilities.PadInt(i, 2), func(t *testing.T) {
			r := utilities.AbsI(tc.v)

			assert.Equal(t, tc.expected, r)
		})
	}
}

type modITestCase struct {
	v        int
	m        int
	expected int
}

func TestModI(t *testing.T) {
	testCases := []modITestCase{
		{10, 3, 1},
		{15, 6, 3},
		{3, 3, 0},
		{-3, 3, 0},
		{3, -3, 0},
		{8, -4, 0},
	}

	for i, tc := range testCases {
		t.Run(utilities.PadInt(i, 2), func(t *testing.T) {
			r := utilities.ModI(tc.v, tc.m)

			assert.Equal(t, tc.expected, r)
		})
	}
}
