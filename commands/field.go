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

func DefineFieldCommand(rootCmd *cobra.Command) {

	options := fieldOptions{}

	command := &cobra.Command{
		Use:     "field",
		Example: "field --index 1 --separator ,",
		Short:   "Returns the field at the given index",
		Run: func(cmd *cobra.Command, args []string) {

			input := utilities.GetStandardInputString()

			var fields []string
			if options.ignoreEmptyFields {
				fields = strings.FieldsFunc(input, func(r rune) bool {
					return strings.ContainsRune(options.fieldSeparator, r)
				})
			} else {
				fields = strings.Split(input, options.fieldSeparator)
			}

			fieldsCount := len(fields)
			fieldIndex := utilities.ClampI(options.fieldIndex, 0, fieldsCount-1)

			r := fields[fieldIndex]

			fmt.Print(r)

		},
	}

	command.Flags().IntVarP(&options.fieldIndex, "index", "i", 0, "Zero based field index")
	command.Flags().StringVarP(&options.fieldSeparator, "separator", "s", " ", "Field separator")
	command.Flags().BoolVar(&options.ignoreEmptyFields, "ignore-empty", false, "Ignores empty fields")

	rootCmd.AddCommand(command)

}
