package main

import (
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
  fmt.Fprintf(os.Stderr, "%x\n", h.Sum(nil))
}
