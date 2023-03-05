package main

import (
	"errors"
	"fmt"
	"os"
	"regexp"

	"github.com/Gartenschlaeger/strcli/commands"
	"github.com/Gartenschlaeger/strcli/utilities"
	"github.com/spf13/cobra"
)

func setupCommand(ctx *commands.CommandContext, cfg *commands.CommandConfiguration) *cobra.Command {
	cmd := &cobra.Command{
		Use:     cfg.Name,
		Short:   cfg.Description,
		Example: cfg.Example,
		RunE:    cfg.Handler,
	}

	flags := cmd.Flags()
	flags.SetInterspersed(false)

	if cfg.Setup != nil {
		cfg.Setup(cmd, flags)
	}

	if cfg.HasSelectionFlag {
		flags.StringVarP(&ctx.Selection, "selection", "s", "", "Character range selection")
	}

	return cmd
}

func main() {
	input := utilities.GetStandardInputString()

	ctx := commands.NewCommandContext(input)

	rootCmd := &cobra.Command{
		Use:           "str",
		Short:         "Performs general string operations",
		SilenceErrors: false,
		SilenceUsage:  false,
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			if ctx.Selection != "" {
				re, _ := regexp.Compile(`^-?\d+(?::\d+)?$`)
				if !re.MatchString(ctx.Selection) {
					return errors.New("invalid value for --selection")
				}
			}

			return nil
		},
	}

	rootCmd.Version = "1.10.0"

	rootCmd.AddCommand(
		setupCommand(ctx, commands.NewFieldCommand(ctx)),
		setupCommand(ctx, commands.NewReplaceCommand(ctx)),
		setupCommand(ctx, commands.NewSubCommand(ctx)),
		setupCommand(ctx, commands.NewRemoveCommand(ctx)),
		setupCommand(ctx, commands.NewLowerCommand(ctx)),
		setupCommand(ctx, commands.NewUpperCommand(ctx)),
		setupCommand(ctx, commands.NewTrimCommand(ctx)),
		setupCommand(ctx, commands.NewHashCommand(ctx)),
		setupCommand(ctx, commands.NewShuffleCommand(ctx)),
		setupCommand(ctx, commands.NewReverseCommand(ctx)),
		setupCommand(ctx, commands.NewBase64Command(ctx)),
		setupCommand(ctx, commands.NewShiftCommand(ctx)),
		setupCommand(ctx, commands.NewRegexCommand(ctx)),
	)

	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}

	fmt.Print(ctx.Result)
}
