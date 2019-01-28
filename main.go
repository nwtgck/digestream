package main

import (
  "github.com/mattn/go-tty"
  "crypto/sha256"
  "fmt"
  "io"
  "log"
  "os"
)

func main () {
  tee := io.TeeReader(os.Stdin, os.Stdout)
  h := sha256.New()
  _, err := io.Copy(h, tee)
  if err != nil {
    log.Fatal(err)
    fmt.Fprintf(os.Stderr, "Error: %s\n", err.Error())
    os.Exit(1)
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
}
