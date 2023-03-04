package commands

import (
	"strings"

	"github.com/Gartenschlaeger/strcli/utilities"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

type CommandContext struct {
	Input     string
	Result    string
	Selection string
}

type CommandConfiguration struct {
	Name             string
	Description      string
	Example          string
	HasSelectionFlag bool
	Handler          func(cmd *cobra.Command, args []string) error
	Setup            func(cmd *cobra.Command, flags *pflag.FlagSet)
}

type CommandHandler = func(input string) (string, error)

func NewCommandContext(input string) *CommandContext {
	return &CommandContext{
		Input:  input,
		Result: "",
	}
}

func parseSelection(selection string, input string) (index int, length int) {
	p := strings.Split(selection, ":")
	if len(p) == 2 {
		return utilities.ParseInt(p[0], 0), utilities.ParseInt(p[1], 0)
	}

	f := utilities.ParseInt(p[0], 0)

	if f < 0 {
		return len(input) + f, -f
	}

	return 0, f
}

func ProcessResult(ctx *CommandContext, handler CommandHandler) error {
	if ctx.Selection == "" {
		r, err := handler(ctx.Input)
		if err != nil {
			return err
		}

		ctx.Result = r
	} else {
		startIndex, endIndex := parseSelection(ctx.Selection, ctx.Input)
		startIndex, endIndex = utilities.ClampStringPartion(ctx.Input, startIndex, endIndex)

		ss := ctx.Input[startIndex:endIndex]
		sr, err := handler(ss)
		if err != nil {
			return err
		}

		ctx.Result = ctx.Input[:startIndex] + sr + ctx.Input[endIndex:]
	}

	return nil
}
