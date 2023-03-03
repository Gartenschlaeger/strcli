package commands_test

import (
	"testing"

	"github.com/Gartenschlaeger/strcli/commands"
	"github.com/stretchr/testify/assert"
)

type trimTestCase struct {
	testName       string
	input          string
	cutset         string
	expectedResult string
}

func TestTrimCommand(t *testing.T) {
	testCases := []trimTestCase{
		{"trims spefied character", "!!Test!!!", "!", "Test"},
		{"trims spefied character", "  ..   Test .  ", ".", "  ..   Test .  "},
		{"trims spefied character", "..   Test .", ".", "   Test "},
		{"trims spefied characters", "¡¡¡Hello, Gophers!!!", "!¡", "Hello, Gophers"},
		{"trims spefied character", "  \nTest..", ".", "  \nTest"},
	}

	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			ctx := commands.NewContext(tc.input)

			opt := commands.TrimCommandOptions{
				Cutset: tc.cutset,
			}

			commands.TrimCommandHandler(ctx, &opt)

			assert.Equal(t, tc.expectedResult, ctx.Result)
		})
	}
}
