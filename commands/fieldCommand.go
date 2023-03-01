package commands

import (
	"os"
	"strings"

	"github.com/Gartenschlaeger/strcli/utilities"
	"github.com/spf13/cobra"
)

func DefineFieldCommand(rootCmd *cobra.Command) {
	var fieldIndex int
	var fieldSeparator string

	fieldCommand := &cobra.Command{
		Use:   "field",
		Short: "Returns the field at the given index",
		Run: func(cmd *cobra.Command, args []string) {
			input := utilities.GetStandardInputString()

			fields := strings.Split(input, fieldSeparator)

			if len(fields) > fieldIndex {
				cmd.Print(fields[fieldIndex])
			} else {
				cmd.PrintErr("Invalid field index")
				os.Exit(1)
			}
		},
	}

	fieldCommand.Flags().IntVarP(&fieldIndex, "index", "i", 1, "Zero based field index")
	fieldCommand.Flags().StringVarP(&fieldSeparator, "separator", "s", " ", "Field separator")

	rootCmd.AddCommand(fieldCommand)
}
