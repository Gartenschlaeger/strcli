package commands_test

import (
	"testing"

	"github.com/Gartenschlaeger/strcli/commands"
	"github.com/stretchr/testify/assert"
)

type sha1TestCase struct {
	testName       string
	input          string
	expectedResult string
}

func TestSha1Command(t *testing.T) {
	testCases := []sha1TestCase{
		{"calculates sha1 hash", "1234", "7110eda4d09e062aa5e4a390b0a572ac0d2c0220"},
	}

	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			ctx := commands.NewContext(tc.input)

			commands.ShaCommandHandler(ctx)

			assert.EqualValues(t, tc.expectedResult, ctx.Result)
		})
	}
}
