package commands_test

import (
	"strconv"
	"testing"

	"github.com/Gartenschlaeger/strcli/commands"
	"github.com/stretchr/testify/assert"
)

type replaceTestCase struct {
	input        string
	oldValue     string
	newValue     string
	ignoreCasing bool
	replaceAll   bool
	expected     string
}

func defineTestCases() *[]replaceTestCase {
	testCases := []replaceTestCase{
		{
			input:        "",
			oldValue:     "x",
			newValue:     "y",
			ignoreCasing: false,
			replaceAll:   false,
			expected:     "",
		},
		{
			input:        "test 123",
			oldValue:     "t",
			newValue:     "_",
			ignoreCasing: false,
			replaceAll:   false,
			expected:     "_est 123",
		},
		{
			input:        "test 123",
			oldValue:     "t",
			newValue:     "_",
			ignoreCasing: false,
			replaceAll:   true,
			expected:     "_es_ 123",
		},
		{
			input:        "TEST 123",
			oldValue:     "t",
			newValue:     "_",
			ignoreCasing: true,
			replaceAll:   false,
			expected:     "_EST 123",
		},
		{
			input:        "Test 123",
			oldValue:     "t",
			newValue:     "_",
			ignoreCasing: true,
			replaceAll:   true,
			expected:     "_es_ 123",
		},
		{
			input:        "Max,Mustermann,Musterweg 123,Musterhausen",
			oldValue:     "Muster",
			newValue:     "Test",
			ignoreCasing: false,
			replaceAll:   true,
			expected:     "Max,Testmann,Testweg 123,Testhausen",
		},
	}

	return &testCases
}

func TestReplaceCommand(t *testing.T) {
	testCases := *defineTestCases()

	for i, tc := range testCases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			ctx := commands.NewCommandContext(tc.input)

			opt := commands.ReplaceCommandOptions{
				OldValue:     tc.oldValue,
				NewValue:     tc.newValue,
				ReplaceAll:   tc.replaceAll,
				IgnoreCasing: tc.ignoreCasing,
			}

			commands.ReplaceCommandHandler(ctx, &opt)

			assert.EqualValues(t, tc.expected, ctx.Result)
		})
	}
}
