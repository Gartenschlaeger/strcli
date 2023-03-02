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

func AddSubCommand(rootCmd *cobra.Command) {
	opt := subOptions{}

	cmd := &cobra.Command{
		Use:   "sub",
		Short: "Returns a partition beginning at index",
		Run: func(cmd *cobra.Command, args []string) {
			input := utilities.GetStandardInputString()

			length := len(input)

			startIndex := utilities.ClampI(opt.startIndex, 0, length)
			endIndex := utilities.ClampI(opt.startIndex+opt.length, startIndex, length)

			r := input[startIndex:endIndex]

			fmt.Print(r)
		},
	}

	flags := cmd.Flags()
	flags.SetInterspersed(false)

	flags.IntVarP(&opt.startIndex, "index", "i", 0, "Zero based index of the first character")
	flags.IntVarP(&opt.length, "length", "l", 1, "Number of characters")

	rootCmd.AddCommand(cmd)
}
