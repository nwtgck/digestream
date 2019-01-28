package cmd

import (
  "fmt"
  "github.com/mattn/go-tty"
  "hash"
  "io"
  "os"
)

func pipeToStdoutAndCalcHash(h hash.Hash) error {
  // Create a tee from stdin to stdout
  tee := io.TeeReader(os.Stdin, os.Stdout)
  // Calculate hash value
  _, err := io.Copy(h, tee)
  if err != nil {
    return err
  }

  // Open tty
  tty, err := tty.Open()
  defer tty.Close()

  // Output of hash
  var hOut io.Writer
  // If not error
  if err == nil {
    // Set hash output as tty
    hOut = tty.Output()
  } else {
    // Set hash output as stderr
    hOut = os.Stderr
  }

  // Print hash value
  fmt.Fprintf(hOut, "%x\n", h.Sum(nil))

  return nil
}
