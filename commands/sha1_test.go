package commands_test

import (
	"strconv"
	"testing"

	"github.com/Gartenschlaeger/strcli/commands"
	"github.com/stretchr/testify/assert"
)

type sha1TestCase struct {
	input    string
	expected string
}

func TestSha1Command(t *testing.T) {
	testCases := []sha1TestCase{
		{"", "da39a3ee5e6b4b0d3255bfef95601890afd80709"},
		{"1234", "7110eda4d09e062aa5e4a390b0a572ac0d2c0220"},
	}

	for i, tc := range testCases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			ctx := commands.NewCommandContext(tc.input)

			err := commands.ShaCommandHandler(ctx)

			assert.NoError(t, err)
			assert.EqualValues(t, tc.expected, ctx.Result)
		})
	}
}
