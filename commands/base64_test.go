package commands_test

import (
	"strconv"
	"testing"

	"github.com/Gartenschlaeger/strcli/commands"
	"github.com/stretchr/testify/assert"
)

type base64TestCase struct {
	input    string
	mode     string
	expected string
}

func TestBaseCommand(t *testing.T) {
	testCases := []base64TestCase{
		{"", "enc", ""},
		{"", "dec", ""},
		{"1234", "enc", "MTIzNA=="},
		{"MTIzNA==", "dec", "1234"},
	}

	for i, tc := range testCases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			opt := commands.Base64CommandOptions{
				Mode: tc.mode,
			}

			ctx := commands.NewCommandContext(tc.input)

			err := commands.Base64CommandHandler(ctx, &opt)

			assert.NoError(t, err)
			assert.EqualValues(t, tc.expected, ctx.Result)
		})
	}
}
