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

func FieldCommandHandler(ctx *CommandContext, opt *FieldCommandOptions) {
	var fields []string
	if opt.IgnoreEmpty {
		fields = strings.FieldsFunc(ctx.Input, func(r rune) bool {
			return strings.ContainsRune(opt.Separator, r)
		})
	} else {
		fields = strings.Split(ctx.Input, opt.Separator)
	}

	fieldsCount := len(fields)
	fieldIndex := utilities.ClampI(opt.Index, 0, fieldsCount-1)

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
