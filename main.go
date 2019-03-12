package main

import (
	"fmt"
	"github.com/nwtgck/digestream/cmd"
	"os"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err.Error())
		os.Exit(1)
	}
}
