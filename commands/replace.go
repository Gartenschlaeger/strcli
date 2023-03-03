package commands

import (
	"errors"
	"strings"

	"github.com/spf13/cobra"
)

type ReplaceCommandOptions struct {
	OldValue     string
	NewValue     string
	ReplaceAll   bool
	IgnoreCasing bool
}

func replaceStringPortion(input string, index int, oldValue string) string {
	if index+len(oldValue) > len(input) {
		return input
	}

	return input[:index] + oldValue + input[index+len(oldValue):]
}

func ReplaceCommandHandler(ctx *CommandContext, opt *ReplaceCommandOptions) {
	if !opt.IgnoreCasing {
		if opt.ReplaceAll {
			ctx.Result = strings.ReplaceAll(ctx.Input, opt.OldValue, opt.NewValue)
		} else {
			ctx.Result = strings.Replace(ctx.Input, opt.OldValue, opt.NewValue, 1)
		}
	} else {
		inputLower := strings.ToLower(ctx.Input)
		oldValueLower := strings.ToLower(opt.OldValue)

		findIndicees := []int{}
		for i := 0; i < len(inputLower); i++ {
			if strings.HasPrefix(inputLower[i:], oldValueLower) {
				findIndicees = append(findIndicees, i)
				i += len(oldValueLower)

				if !opt.ReplaceAll {
					break
				}
			}
		}

		result := ctx.Input
		for _, v := range findIndicees {
			result = replaceStringPortion(result, v, opt.NewValue)
		}

		ctx.Result = result
	}
}

func NewReplaceCommand(ctx *CommandContext) *cobra.Command {
	opt := ReplaceCommandOptions{}

	cmd := &cobra.Command{
		Use:     "replace",
		Example: "str replace -o \" \" -n \"_\" -a",
		Short:   "Replaces occurrences with a new value",
		RunE: func(cmd *cobra.Command, args []string) error {
			if opt.OldValue == "" {
				return errors.New("flag --old is required")
			}

			ReplaceCommandHandler(ctx, &opt)

			return nil
		},
	}

	flags := cmd.Flags()
	flags.SetInterspersed(false)

	flags.StringVarP(&opt.OldValue, "old", "o", "", "Value to be replaced")
	flags.StringVarP(&opt.NewValue, "new", "n", "", "New value to replace the old value")
	flags.BoolVarP(&opt.ReplaceAll, "replace-all", "a", false, "Replace all occurrences instead of first one only")
	flags.BoolVarP(&opt.IgnoreCasing, "ignore-casing", "i", false, "Ignore casing when comparing old values")

	return cmd
}
