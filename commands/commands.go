package commands

import (
	"fmt"
	"os"
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
	name             string
	description      string
	example          string
	hasSelectionFlag bool
	handler          func(cmd *cobra.Command, args []string) error
	setupFlags       func(flags *pflag.FlagSet)
}

func NewContext(input string) *CommandContext {
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

func ValidateSelection(input string, selection string) (index int, length int) {
	if selection == "" {
		return 0, len(input)
	}

	i, l := parseSelection(selection, input)
	i, l = utilities.ClampStringPartion(input, i, l)

	return i, l
}

func ProcessResult(ctx *CommandContext, handlerCallback func(input string) string) {
	if ctx.Selection == "" {
		ctx.Result = handlerCallback(ctx.Input)
	}

	startIndex, endIndex := ValidateSelection(ctx.Input, ctx.Selection)

	ss := ctx.Input[startIndex:endIndex]
	sr := handlerCallback(ss)

	ctx.Result = ctx.Input[:startIndex] + sr + ctx.Input[endIndex:]
}

func SetupCommand(ctx *CommandContext, cmd *CommandConfiguration) *cobra.Command {
	c := &cobra.Command{
		Use:     cmd.name,
		Short:   cmd.description,
		Example: cmd.example,
		RunE:    cmd.handler,
	}

	flags := c.Flags()
	flags.SetInterspersed(false)

	if cmd.setupFlags != nil {
		cmd.setupFlags(flags)
	}

	if cmd.hasSelectionFlag {
		flags.StringVarP(&ctx.Selection, "selection", "s", "", "Character range selection")
	}

	return c
}

func Execute() {
	input := utilities.GetStandardInputString()

	ctx := NewContext(input)

	rootCmd := &cobra.Command{
		Use:   "str",
		Short: "Performs general string operations",
		PersistentPostRun: func(cmd *cobra.Command, args []string) {
			fmt.Print(ctx.Result)
		},
	}

	rootCmd.Version = "1.6.0"

	rootCmd.AddCommand(
		SetupCommand(ctx, NewFieldCommand(ctx)),
		SetupCommand(ctx, NewReplaceCommand(ctx)),
		SetupCommand(ctx, NewSubCommand(ctx)),
		SetupCommand(ctx, NewRemoveCommand(ctx)),
		SetupCommand(ctx, NewLowerCommand(ctx)),
		SetupCommand(ctx, NewUpperCommand(ctx)),
		SetupCommand(ctx, NewTrimCommand(ctx)),
		SetupCommand(ctx, NewMd5Command(ctx)),
		SetupCommand(ctx, NewSha1Command(ctx)),
		SetupCommand(ctx, NewShuffleCommand(ctx)),
		SetupCommand(ctx, NewReverseCommand(ctx)),
	)

	err := rootCmd.Execute()
	if err != nil {
		fmt.Fprint(os.Stderr, err.Error())
		os.Exit(1)
	}
}
