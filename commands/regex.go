package commands

import (
	"regexp"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

type RegexCommandOptions struct {
	Pattern string
	Group   int
}

func RegexCommandHandler(ctx *CommandContext, opt *RegexCommandOptions) error {
	re, err := regexp.Compile(opt.Pattern)
	if err != nil {
		return err
	}

	r := re.FindStringSubmatch(ctx.Input)

	l := len(r)
	if l > opt.Group {
		ctx.Result = r[opt.Group]
	} else {
		ctx.Result = ""
	}

	return nil
}

func NewRegexCommand(ctx *CommandContext) *CommandConfiguration {
	opt := &RegexCommandOptions{}

	cmd := &CommandConfiguration{
		Name:        "regex",
		Description: "Performs a search using a regular expression",
		Handler: func(cmd *cobra.Command, args []string) error {
			return RegexCommandHandler(ctx, opt)
		},
		Setup: func(cmd *cobra.Command, flags *pflag.FlagSet) {
			flags.StringVarP(&opt.Pattern, "pattern", "p", "", "Regular expression pattern")
			flags.IntVarP(&opt.Group, "group", "g", 0, "The group to be returns (if pattern contains groups)")

			cmd.MarkFlagRequired("pattern")
		},
	}

	return cmd
}
