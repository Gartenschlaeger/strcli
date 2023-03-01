package commands

import (
	"fmt"

	"github.com/Gartenschlaeger/strcli/utilities"
	"github.com/spf13/cobra"
)

type subOptions struct {
	startIndex int
	length     int
}

func DefineSubCommand(rootCmd *cobra.Command) {

	options := subOptions{}

	command := &cobra.Command{
		Use:     "sub",
		Example: "sub --start-index 5 --length 3",
		Short:   "Returns a portion of the string",
		Run: func(cmd *cobra.Command, args []string) {

			input := utilities.GetStandardInputString()

			length := len(input)

			startIndex := utilities.ClampI(options.startIndex, 0, length)
			endIndex := utilities.ClampI(options.startIndex+options.length, startIndex, length)

			r := input[startIndex:endIndex]

			fmt.Print(r)

		},
	}

	command.Flags().IntVarP(&options.startIndex, "start-index", "s", 0, "Zero based index of the first character")
	command.Flags().IntVarP(&options.length, "length", "l", 1, "Number of characters")

	rootCmd.AddCommand(command)

}
