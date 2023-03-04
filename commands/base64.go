package commands

import (
	"github.com/Gartenschlaeger/strcli/utilities"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

type Base64CommandOptions struct {
	Mode string
}

func Base64CommandHandler(ctx *CommandContext, opt *Base64CommandOptions) error {
	if utilities.StringEquals(opt.Mode, "enc", true) {
		ctx.Result = utilities.Base64Encode(ctx.Input)
	} else {
		r, err := utilities.Base64Decode(ctx.Input)
		if err != nil {
			return err
		}

		ctx.Result = r
	}

	return nil
}

func NewBase64Command(ctx *CommandContext) *CommandConfiguration {
	opt := &Base64CommandOptions{}

	cmd := &CommandConfiguration{
		name:        "base64",
		description: "Encodes or decodes Base64",
		handler: func(cmd *cobra.Command, args []string) error {
			return Base64CommandHandler(ctx, opt)
		},
		setup: func(cmd *cobra.Command, flags *pflag.FlagSet) {
			flags.StringVarP(&opt.Mode, "mode", "m", "enc", "Specifies the mode (enc | dec)")
		},
	}

	return cmd
}
