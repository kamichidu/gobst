package main

import (
	"os"

	"github.com/kamichidu/gobst/cli"
)

func init() {
	cli.Version = "v0.0.5"
}

func main() {
	os.Exit(cli.Run(os.Stdin, os.Stdout, os.Stderr, os.Args))
}
