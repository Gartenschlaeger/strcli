package commands

import (
	"regexp"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

type RegexCommandOptions struct {
	Pattern string
	Replace string
	Group   int
}

func RegexCommandHandler(ctx *CommandContext, opt *RegexCommandOptions) error {
	re, err := regexp.Compile(opt.Pattern)
	if err != nil {
		return err
	}

	if opt.Replace != "" {
		// use replace pattern
		r := re.ReplaceAllString(ctx.Input, opt.Replace)

		ctx.Result = r
	} else {
		// use search pattern
		r := re.FindStringSubmatch(ctx.Input)

		l := len(r)
		if l > opt.Group {
			ctx.Result = r[opt.Group]
		} else {
			// if the given group index is out or range, return an empty result
			ctx.Result = ""
		}
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
			flags.IntVarP(&opt.Group, "group", "g", 0, "The group to be returned")
			flags.StringVarP(&opt.Replace, "replace", "r", "", "Replace pattern")

			cmd.MarkFlagRequired("pattern")
		},
	}

	return cmd
}
