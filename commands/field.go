package commands

import (
	"fmt"
	"strings"

	"github.com/Gartenschlaeger/strcli/utilities"
	"github.com/spf13/cobra"
)

type fieldOptions struct {
	fieldIndex        int
	fieldSeparator    string
	ignoreEmptyFields bool
}

func AddFieldCommand(rootCmd *cobra.Command) {
	opt := fieldOptions{}

	cmd := &cobra.Command{
		Use:   "field",
		Short: "Returns the field at the specified position",
		Run: func(cmd *cobra.Command, args []string) {
			input := utilities.GetStandardInputString()

			var fields []string
			if opt.ignoreEmptyFields {
				fields = strings.FieldsFunc(input, func(r rune) bool {
					return strings.ContainsRune(opt.fieldSeparator, r)
				})
			} else {
				fields = strings.Split(input, opt.fieldSeparator)
			}

			fieldsCount := len(fields)
			fieldIndex := utilities.ClampI(opt.fieldIndex, 0, fieldsCount-1)

			r := fields[fieldIndex]

			fmt.Print(r)
		},
	}

	flags := cmd.Flags()
	flags.SetInterspersed(false)

	flags.IntVarP(&opt.fieldIndex, "index", "i", 0, "Zero based field index")
	flags.StringVarP(&opt.fieldSeparator, "separator", "s", " ", "Field separator")
	flags.BoolVar(&opt.ignoreEmptyFields, "ignore-empty", false, "Ignores empty fields")

	rootCmd.AddCommand(cmd)
}
