package commands

import (
	"crypto/sha1"
	"encoding/hex"

	"github.com/spf13/cobra"
)

func getSha1Hash(input string) string {
	h := sha1.New()
	h.Write([]byte(input))
	bs := h.Sum(nil)

	return hex.EncodeToString(bs)
}

func ShaCommandHandler(ctx *CommandContext) {
	ctx.Result = getSha1Hash(ctx.Input)
}

func NewSha1Command(ctx *CommandContext) *CommandConfiguration {
	cmd := &CommandConfiguration{
		name:        "sha1",
		description: "Calculates a SHA1 hash",
		handler: func(cmd *cobra.Command, args []string) error {
			ShaCommandHandler(ctx)

			return nil
		},
	}

	return cmd
}
