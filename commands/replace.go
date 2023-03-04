package commands

import (
	"github.com/Gartenschlaeger/strcli/utilities"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

type ReplaceCommandOptions struct {
	OldValue     string
	NewValue     string
	ReplaceAll   bool
	IgnoreCasing bool
}

func ReplaceCommandHandler(ctx *CommandContext, opt *ReplaceCommandOptions) error {
	ctx.Result = utilities.ReplaceString(ctx.Input, opt.OldValue, opt.NewValue, opt.ReplaceAll, opt.IgnoreCasing)

	return nil
}

func NewReplaceCommand(ctx *CommandContext) *CommandConfiguration {
	opt := ReplaceCommandOptions{}

	cmd := &CommandConfiguration{
		name:        "replace",
		description: "Replaces occurrences with a new value",
		example:     "str replace -o \" \" -n \"_\" -a",
		handler: func(cmd *cobra.Command, args []string) error {
			return ReplaceCommandHandler(ctx, &opt)
		},
		setup: func(cmd *cobra.Command, flags *pflag.FlagSet) {
			flags.StringVarP(&opt.OldValue, "old", "o", "", "Value to be replaced")
			flags.StringVarP(&opt.NewValue, "new", "n", "", "New value to replace the old value")
			flags.BoolVarP(&opt.ReplaceAll, "replace-all", "a", false, "Replace all occurrences instead of first one only")
			flags.BoolVarP(&opt.IgnoreCasing, "ignore-casing", "i", false, "Ignore casing when comparing old values")

			cmd.MarkFlagRequired("old")
		},
	}

	return cmd
}
