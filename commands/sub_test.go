package commands_test

import (
	"testing"

	"github.com/Gartenschlaeger/strcli/commands"
	"github.com/stretchr/testify/assert"
)

type subTestCase struct {
	testName       string
	input          string
	index          int
	length         int
	expectedResult string
}

func Test(t *testing.T) {
	testCases := []subTestCase{
		{"index 0 length 4", "Some text", 0, 4, "Some"},
		{"index 5 length 5", "Some text", 5, 5, "text"},
		{"index -4 length 4", "Some text", -4, 4, "text"},
		{"index -4 length 3", "Some text", -4, 3, "tex"},
	}

	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			ctx := commands.NewContext(tc.input)

			opt := commands.SubCommandOptions{
				Index:  tc.index,
				Length: tc.length,
			}

			commands.SubCommandHandler(ctx, &opt)

			assert.Equal(t, tc.expectedResult, ctx.Result)
		})
	}
}
