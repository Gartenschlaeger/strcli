package commands_test

import (
	"testing"

	"github.com/Gartenschlaeger/strcli/commands"
	"github.com/Gartenschlaeger/strcli/utilities"
	"github.com/stretchr/testify/assert"
)

type hashTestCase struct {
	input           string
	mode            string
	expected_result string
	expected_error  bool
}

func TestHashCommand(t *testing.T) {
	testCases := []hashTestCase{
		{"", "", "", true},
		{"1234", "INVALID_MODE", "", true},
		{"1234", "MD5", "81dc9bdb52d04dc20036dbd8313ed055", false},
		{"1234", "SHA1", "7110eda4d09e062aa5e4a390b0a572ac0d2c0220", false},
		{"1234", "SHA256", "03ac674216f3e15c761ee1a5e255f067953623c8b388b4459e13f978d7c846f4", false},
		{"1234", "SHA512", "d404559f602eab6fd602ac7680dacbfaadd13630335e951f097af3900e9de176b6db28512f2e000b9d04fba5133e8b1c6e8df59db3a8ab9d60be4b97cc9e81db", false},
	}

	for i, tc := range testCases {
		t.Run(utilities.PadInt(i, 2), func(t *testing.T) {
			ctx := commands.NewCommandContext(tc.input)

			opt := commands.HashCommandOptions{
				Mode: tc.mode,
			}

			err := commands.HashCommandHandler(ctx, &opt)

			if tc.expected_error {
				assert.Error(t, err)
				assert.EqualValues(t, tc.expected_result, ctx.Result)
			} else {
				assert.NoError(t, err)
				assert.EqualValues(t, tc.expected_result, ctx.Result)
			}
		})
	}
}
