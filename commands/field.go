package commands

import (
	"github.com/Gartenschlaeger/strcli/utilities"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

type FieldCommandOptions struct {
	Index       int
	Separator   string
	IgnoreEmpty bool
}

func FieldCommandHandler(ctx *CommandContext, opt *FieldCommandOptions) error {
	fields := utilities.SplitString(ctx.Input, opt.Separator, opt.IgnoreEmpty)
	fieldsCount := len(fields)

	if fieldsCount > 0 {
		fieldIndex := opt.Index

		if opt.Index < 0 {
			fieldIndex = fieldsCount + opt.Index
		}

		fieldIndex = utilities.ClampI(fieldIndex, 0, fieldsCount-1)

		ctx.Result = fields[fieldIndex]
	}

	return nil
}

func NewFieldCommand(ctx *CommandContext) *CommandConfiguration {
	opt := FieldCommandOptions{}

	cmd := &CommandConfiguration{
		Name:        "field",
		Description: "Returns the field at the specified position",
		Handler: func(cmd *cobra.Command, args []string) error {
			return FieldCommandHandler(ctx, &opt)
		},
		Setup: func(cmd *cobra.Command, flags *pflag.FlagSet) {
			flags.IntVarP(&opt.Index, "index", "i", 0, "Zero based field index")
			flags.StringVarP(&opt.Separator, "separator", "s", " ", "Field separator")
			flags.BoolVar(&opt.IgnoreEmpty, "ignore-empty", false, "Ignores empty fields")
		},
	}

	return cmd
}
