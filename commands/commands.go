package commands

import (
	"github.com/Gartenschlaeger/strcli/utilities"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

// Represents the execution context for sub commands
type CommandContext struct {
	Input     string
	Result    string
	Selection string
}

// Represents the command configuration used to setup a sub command
type CommandConfiguration struct {
	Name             string
	Description      string
	Example          string
	HasSelectionFlag bool
	Handler          func(cmd *cobra.Command, args []string) error
	Setup            func(cmd *cobra.Command, flags *pflag.FlagSet)
}

// Command handler used by ProcessSelectionResult
type SelectionCommandHandler = func(input string) (string, error)

// Returns a new CommandContext object
func NewCommandContext(input string) *CommandContext {
	return &CommandContext{
		Input:  input,
		Result: "",
	}
}

// Must be used in command handler in case the command configuration HasSelectionFlag is true
func ProcessSelectionResult(ctx *CommandContext, handler SelectionCommandHandler) error {
	if ctx.Selection == "" {
		r, err := handler(ctx.Input)
		if err != nil {
			return err
		}

		ctx.Result = r
	} else {
		startIndex, endIndex := utilities.ParseSelection(ctx.Selection, ctx.Input)
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
