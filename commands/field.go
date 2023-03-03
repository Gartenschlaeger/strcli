package commands

import (
	"strings"

	"github.com/Gartenschlaeger/strcli/utilities"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

type FieldCommandOptions struct {
	Index       int
	Separator   string
	IgnoreEmpty bool
}

func splitString(input string, separator string, removeEmptyFields bool) []string {
	if removeEmptyFields {
		return strings.FieldsFunc(input, func(r rune) bool {
			return strings.ContainsRune(separator, r)
		})
	} else {
		return strings.Split(input, separator)
	}
}

func FieldCommandHandler(ctx *CommandContext, opt *FieldCommandOptions) {
	fields := splitString(ctx.Input, opt.Separator, opt.IgnoreEmpty)
	fieldsCount := len(fields)

	fieldIndex := opt.Index

	if opt.Index < 0 {
		fieldIndex = fieldsCount + opt.Index
	}

	fieldIndex = utilities.ClampI(fieldIndex, 0, fieldsCount-1)

	ctx.Result = fields[fieldIndex]
}

func NewFieldCommand(ctx *CommandContext) *CommandConfiguration {
	opt := FieldCommandOptions{}

	cmd := &CommandConfiguration{
		name:        "field",
		description: "Returns the field at the specified position",
		handler: func(cmd *cobra.Command, args []string) error {
			FieldCommandHandler(ctx, &opt)

			return nil
		},
		setup: func(cmd *cobra.Command, flags *pflag.FlagSet) {
			flags.IntVarP(&opt.Index, "index", "i", 0, "Zero based field index")
			flags.StringVarP(&opt.Separator, "separator", "s", " ", "Field separator")
			flags.BoolVar(&opt.IgnoreEmpty, "ignore-empty", false, "Ignores empty fields")
		},
	}

	return cmd
}
