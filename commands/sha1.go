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

func NewSha1Command(ctx *CommandContext) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "sha1",
		Short: "Calculates the SHA1 hash",
		Run: func(cmd *cobra.Command, args []string) {
			ShaCommandHandler(ctx)
		},
	}

	flags := cmd.Flags()
	flags.SetInterspersed(false)

	return cmd
}
