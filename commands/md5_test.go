package commands_test

import (
	"strconv"
	"testing"

	"github.com/Gartenschlaeger/strcli/commands"
	"github.com/stretchr/testify/assert"
)

type md5TestCase struct {
	input    string
	expected string
}

func TestMd5Command(t *testing.T) {
	testCases := []md5TestCase{
		{"1234", "81dc9bdb52d04dc20036dbd8313ed055"},
	}

	for i, tc := range testCases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			ctx := commands.NewContext(tc.input)

			commands.Md5CommandHandler(ctx)

			assert.EqualValues(t, tc.expected, ctx.Result)
		})
	}
}
