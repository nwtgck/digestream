package cmd

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"github.com/spf13/cobra"
	"os"
)

func init() {
	cobra.OnInitialize()
	RootCmd.AddCommand(md5Cmd)
	RootCmd.AddCommand(sha1Cmd)
	RootCmd.AddCommand(sha256Cmd)
	RootCmd.AddCommand(sha512Cmd)
}

var RootCmd = &cobra.Command{
	Use:   os.Args[0],
	Short: "digestream",
	Long:  "Dige(st + St)ream",
	RunE: func(cmd *cobra.Command, args []string) error {
		return pipeToStdoutAndCalcHash(sha256.New())
	},
}

var md5Cmd = &cobra.Command{
	Use:   "md5",
	Short: "MD5",
	RunE: func(cmd *cobra.Command, args []string) error {
		h := md5.New()
		return pipeToStdoutAndCalcHash(h)
	},
}

var sha1Cmd = &cobra.Command{
	Use:   "sha1",
	Short: "SHA-1",
	RunE: func(cmd *cobra.Command, args []string) error {
		h := sha1.New()
		return pipeToStdoutAndCalcHash(h)
	},
}

var sha256Cmd = &cobra.Command{
	Use:   "sha256",
	Short: "SHA-256",
	RunE: func(cmd *cobra.Command, args []string) error {
		h := sha256.New()
		return pipeToStdoutAndCalcHash(h)
	},
}

var sha512Cmd = &cobra.Command{
	Use:   "sha512",
	Short: "SHA-512",
	RunE: func(cmd *cobra.Command, args []string) error {
		h := sha512.New()
		return pipeToStdoutAndCalcHash(h)
	},
}
