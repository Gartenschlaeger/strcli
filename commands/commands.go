package commands

import (
	"errors"
	"fmt"
	"os"
	"regexp"
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
	setup            func(cmd *cobra.Command, flags *pflag.FlagSet)
}

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

func ProcessResult(ctx *CommandContext, handlerCallback func(input string) string) {
	if ctx.Selection == "" {
		ctx.Result = handlerCallback(ctx.Input)
	} else {
		startIndex, endIndex := parseSelection(ctx.Selection, ctx.Input)
		startIndex, endIndex = utilities.ClampStringPartion(ctx.Input, startIndex, endIndex)

		ss := ctx.Input[startIndex:endIndex]
		sr := handlerCallback(ss)

		ctx.Result = ctx.Input[:startIndex] + sr + ctx.Input[endIndex:]
	}
}

func setupCommand(ctx *CommandContext, cfg *CommandConfiguration) *cobra.Command {
	cmd := &cobra.Command{
		Use:     cfg.name,
		Short:   cfg.description,
		Example: cfg.example,
		RunE:    cfg.handler,
	}

	flags := cmd.Flags()
	flags.SetInterspersed(false)

	if cfg.setup != nil {
		cfg.setup(cmd, flags)
	}

	if cfg.hasSelectionFlag {
		flags.StringVarP(&ctx.Selection, "selection", "s", "", "Character range selection")
	}

	return cmd
}

func Execute() {
	input := utilities.GetStandardInputString()

	ctx := NewCommandContext(input)

	rootCmd := &cobra.Command{
		Use:           "str",
		Short:         "Performs general string operations",
		SilenceErrors: false,
		SilenceUsage:  false,
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			if ctx.Selection != "" {
				re, _ := regexp.Compile(`^-?\d+(?:,\d+)?$`)
				if !re.MatchString(ctx.Selection) {
					return errors.New("invalid value for --selection")
				}
			}

			return nil
		},
	}

	rootCmd.Version = "1.6.0"

	rootCmd.AddCommand(
		setupCommand(ctx, NewFieldCommand(ctx)),
		setupCommand(ctx, NewReplaceCommand(ctx)),
		setupCommand(ctx, NewSubCommand(ctx)),
		setupCommand(ctx, NewRemoveCommand(ctx)),
		setupCommand(ctx, NewLowerCommand(ctx)),
		setupCommand(ctx, NewUpperCommand(ctx)),
		setupCommand(ctx, NewTrimCommand(ctx)),
		setupCommand(ctx, NewMd5Command(ctx)),
		setupCommand(ctx, NewSha1Command(ctx)),
		setupCommand(ctx, NewShuffleCommand(ctx)),
		setupCommand(ctx, NewReverseCommand(ctx)),
	)

	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}

	fmt.Print(ctx.Result)
}
