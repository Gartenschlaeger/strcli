package commands

import (
	"strings"

	"github.com/Gartenschlaeger/strcli/utilities"
	"github.com/spf13/cobra"
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

func NewFieldCommand(context *CommandContext) *cobra.Command {
	opt := FieldCommandOptions{}

	cmd := &cobra.Command{
		Use:   "field",
		Short: "Returns the field at the specified position",
		Run: func(cmd *cobra.Command, args []string) {
			FieldCommandHandler(context, &opt)
		},
	}

	flags := cmd.Flags()
	flags.SetInterspersed(false)

	flags.IntVarP(&opt.Index, "index", "i", 0, "Zero based field index")
	flags.StringVarP(&opt.Separator, "separator", "s", " ", "Field separator")
	flags.BoolVar(&opt.IgnoreEmpty, "ignore-empty", false, "Ignores empty fields")

	return cmd
}
