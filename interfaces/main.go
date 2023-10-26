// Assignment 2
package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
  if (len(os.Args) < 2) {
    os.Exit(1)
  }

  filepath := strings.Trim(os.Args[1], " ")
  fmt.Println("Reading contents of file ", filepath)

  file, err := os.Open(filepath)
  if err != nil {
    log.Fatal("Error reading file: ", err)
  }

  io.Copy(os.Stdout, file)
  file.Close()
}