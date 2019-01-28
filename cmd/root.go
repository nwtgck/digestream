package cmd

import (
	"crypto/sha256"
	"github.com/spf13/cobra"
	"os"
)

var RootCmd = &cobra.Command{
	Use:   os.Args[0],
	Short: "digestream",
	Long:  "Dige(st + St)ream",
	RunE: func(cmd *cobra.Command, args []string) error {
		return pipeToStdoutAndCalcHash(sha256.New())
	},
}

func init() {
	cobra.OnInitialize()
}
