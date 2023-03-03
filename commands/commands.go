package commands

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/Gartenschlaeger/strcli/utilities"
	"github.com/spf13/cobra"
)

type CommandContext struct {
	Input  string
	Result string
}

func NewContext(input string) *CommandContext {
	return &CommandContext{
		Input:  input,
		Result: "",
	}
}

func parseInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		i = 0
	}

	return i
}

func parseSelection(selection string, input string) (index int, length int) {
	p := strings.Split(selection, ":")
	if len(p) == 2 {
		return parseInt(p[0]), parseInt(p[1])
	}

	f := parseInt(p[0])

	if f < 0 {
		return len(input) + f, -f
	}

	return 0, f
}

func ValidateSelection(input string, selection string) (index int, length int) {
	if selection == "" {
		return 0, len(input)
	}

	i, l := parseSelection(selection, input)
	i, l = utilities.ClampStringPartion(input, i, l)

	return i, l
}

func ManipulateSelection(ctx *CommandContext, selection string, callback func(input string) string) {
	if selection == "" {
		ctx.Result = callback(ctx.Input)
	}

	startIndex, endIndex := ValidateSelection(ctx.Input, selection)

	ss := ctx.Input[startIndex:endIndex]
	sr := callback(ss)

	ctx.Result = ctx.Input[:startIndex] + sr + ctx.Input[endIndex:]
}

var rootCmd *cobra.Command

func Execute() {
	input := utilities.GetStandardInputString()

	ctx := NewContext(input)

	rootCmd = &cobra.Command{
		Use:   "str",
		Short: "Performs general string operations",
		PersistentPostRun: func(cmd *cobra.Command, args []string) {
			fmt.Print(ctx.Result)
		},
	}

	rootCmd.Version = "1.5.0"

	rootCmd.AddCommand(
		NewFieldCommand(ctx),
		NewReplaceCommand(ctx),
		NewSubCommand(ctx),
		NewRemoveCommand(ctx),
		NewLowerCommand(ctx),
		NewUpperCommand(ctx),
		NewTrimCommand(ctx),
		NewMd5Command(ctx),
		NewSha1Command(ctx),
		NewShuffleCommand(ctx),
		NewReverseCommand(ctx),
	)

	err := rootCmd.Execute()
	if err != nil {
		fmt.Fprint(os.Stderr, err.Error())
		os.Exit(1)
	}
}
