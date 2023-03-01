package commands

import (
	"fmt"
	"strings"

	"github.com/Gartenschlaeger/strcli/utilities"
	"github.com/spf13/cobra"
)

type fieldOptions struct {
	fieldIndex     int
	fieldSeparator string
}

func DefineFieldCommand(rootCmd *cobra.Command) {

	options := fieldOptions{}

	command := &cobra.Command{
		Use:   "field",
		Short: "Returns the field at the given index",
		Run: func(cmd *cobra.Command, args []string) {
			input := utilities.GetStandardInputString()

			fields := strings.Split(input, options.fieldSeparator)

			if len(fields) > options.fieldIndex {
				r := fields[options.fieldIndex]
				fmt.Print(r)
			} else {
				cmd.PrintErr("invalid field index")
			}
		},
	}

	command.Flags().IntVarP(&options.fieldIndex, "index", "i", 1, "Zero based field index")
	command.Flags().StringVarP(&options.fieldSeparator, "separator", "s", " ", "Field separator")

	rootCmd.AddCommand(command)

}
