package commands_test

import (
	"testing"

	"github.com/Gartenschlaeger/strcli/commands"
	"github.com/stretchr/testify/assert"
)

type md5TestCase struct {
	testName       string
	input          string
	expectedResult string
}

func TestMd5Command(t *testing.T) {
	testCases := []md5TestCase{
		{"creates md5", "1234", "81dc9bdb52d04dc20036dbd8313ed055"},
	}

	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			ctx := commands.NewContext(tc.input)

			commands.Md5CommandHandler(ctx)

			assert.EqualValues(t, tc.expectedResult, ctx.Result)
		})
	}
}
