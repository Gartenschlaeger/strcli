package commands

import (
	"crypto/md5"
	"encoding/hex"

	"github.com/spf13/cobra"
)

func getMD5Hash(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}

func Md5CommandHandler(ctx *CommandContext) {
	ctx.Result = getMD5Hash(ctx.Input)
}

func NewMd5Command(ctx *CommandContext) *CommandConfiguration {
	cmd := &CommandConfiguration{
		name:        "md5",
		description: "Calculates a MD5 hash",
		handler: func(cmd *cobra.Command, args []string) error {
			Md5CommandHandler(ctx)

			return nil
		},
	}

	return cmd
}
