package commands_test

import (
	"testing"

	"github.com/Gartenschlaeger/strcli/commands"
	"github.com/Gartenschlaeger/strcli/utilities"
	"github.com/stretchr/testify/assert"
)

type reverseTestCase struct {
	input    string
	expected string
}

func TestReverseCommand(t *testing.T) {
	testCases := []reverseTestCase{
		{"", ""},
		{"123", "321"},
		{"123456789", "987654321"},
	}

	for i, tc := range testCases {
		t.Run(utilities.PadInt(i, 2), func(t *testing.T) {
			ctx := commands.NewCommandContext(tc.input)

			err := commands.ReverseCommandHandler(ctx)

			assert.NoError(t, err)
			assert.Equal(t, tc.expected, ctx.Result)
		})
	}
}
