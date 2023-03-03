package commands_test

import (
	"testing"

	"github.com/Gartenschlaeger/strcli/commands"
	"github.com/stretchr/testify/assert"
)

type replaceTestCase struct {
	testName       string
	input          string
	oldValue       string
	newValue       string
	ignoreCase     bool
	replaceAll     bool
	expectedResult string
}

func defineTestCases() *[]replaceTestCase {
	testCases := []replaceTestCase{
		{
			testName:       "replace t with _",
			input:          "test 123",
			oldValue:       "t",
			newValue:       "_",
			ignoreCase:     false,
			replaceAll:     false,
			expectedResult: "_est 123",
		},
		{
			testName:       "replace all t with _",
			input:          "test 123",
			oldValue:       "t",
			newValue:       "_",
			ignoreCase:     false,
			replaceAll:     true,
			expectedResult: "_es_ 123",
		},
		{
			testName:       "replace and ignore casing",
			input:          "TEST 123",
			oldValue:       "t",
			newValue:       "_",
			ignoreCase:     true,
			replaceAll:     false,
			expectedResult: "_EST 123",
		},
		{
			testName:       "replace all and ignore casing",
			input:          "Test 123",
			oldValue:       "t",
			newValue:       "_",
			ignoreCase:     true,
			replaceAll:     true,
			expectedResult: "_es_ 123",
		},
		{
			testName:       "replace all words",
			input:          "Max,Mustermann,Musterweg 123,Musterhausen",
			oldValue:       "Muster",
			newValue:       "Test",
			ignoreCase:     false,
			replaceAll:     true,
			expectedResult: "Max,Testmann,Testweg 123,Testhausen",
		},
	}

	return &testCases
}

func TestReplaceCommand(t *testing.T) {
	testCases := *defineTestCases()

	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			ctx := commands.NewContext(tc.input)

			opt := commands.ReplaceCommandOptions{
				OldValue:   tc.oldValue,
				NewValue:   tc.newValue,
				ReplaceAll: tc.replaceAll,
				IgnoreCase: tc.ignoreCase,
			}

			commands.ReplaceCommandHandler(ctx, &opt)

			assert.EqualValues(t, tc.expectedResult, ctx.Result)
		})
	}
}
