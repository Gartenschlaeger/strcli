package commands_test

import (
	"testing"

	"github.com/Gartenschlaeger/strcli/commands"
	"github.com/Gartenschlaeger/strcli/utilities"
	"github.com/stretchr/testify/assert"
)

type regexTestCase struct {
	input    string
	pattern  string
	replace  string
	group    int
	expected string
}

func TestRegexCommand(t *testing.T) {
	testCases := []regexTestCase{
		{"", `\s`, "", 0, ""},
		{"05.10.2023", `(\d{1,2}).(\d{1,2}).(\d{4}|\d{2})`, "", 0, "05.10.2023"},
		{"05.10.2023", `(\d{1,2}).(\d{1,2}).(\d{4}|\d{2})`, "", 1, "05"},
		{"05.10.2023", `(\d{1,2}).(\d{1,2}).(\d{4}|\d{2})`, "", 2, "10"},
		{"05.10.2023", `(\d{1,2}).(\d{1,2}).(\d{4}|\d{2})`, "", 3, "2023"},
		{"05.10.2023", `(\d{1,2}).(\d{1,2}).(\d{4}|\d{2})`, "$3-$2-$1", 0, "2023-10-05"},
	}

	for i, tc := range testCases {
		t.Run(utilities.PadInt(i, 2), func(t *testing.T) {
			ctx := commands.NewCommandContext(tc.input)

			opt := &commands.RegexCommandOptions{
				Pattern: tc.pattern,
				Replace: tc.replace,
				Group:   tc.group,
			}

			err := commands.RegexCommandHandler(ctx, opt)

			assert.NoError(t, err)
			assert.Equal(t, tc.expected, ctx.Result)
		})
	}
}
