package commands

import (
	"fmt"
	"strings"

	"github.com/Gartenschlaeger/strcli/utilities"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

type HashCommandOptions struct {
	Mode string
}

func HashCommandHandler(ctx *CommandContext, opt *HashCommandOptions) error {
	switch strings.ToUpper(opt.Mode) {
	case "MD5":
		ctx.Result = utilities.GetMD5Hash(ctx.Input)
	case "SHA1":
		ctx.Result = utilities.GetSHA1Hash(ctx.Input)
	case "SHA256":
		ctx.Result = utilities.GetSHA256Hash(ctx.Input)
	case "SHA512":
		ctx.Result = utilities.GetSHA512Hash(ctx.Input)

	default:
		return fmt.Errorf("invalid mode : %v", opt.Mode)
	}

	return nil
}

func NewHashCommand(ctx *CommandContext) *CommandConfiguration {
	opt := &HashCommandOptions{}

	cmd := &CommandConfiguration{
		Name:        "hash",
		Description: "Calculates a hash",
		Handler: func(cmd *cobra.Command, args []string) error {
			return HashCommandHandler(ctx, opt)
		},
		Setup: func(cmd *cobra.Command, flags *pflag.FlagSet) {
			flags.StringVarP(&opt.Mode, "mode", "m", "", "Hash mode (MD5, SHA1, SHA256, SHA512)")

			cmd.MarkFlagRequired("mode")
		},
	}

	return cmd
}
