package main

import (
	"fmt"
	"kanopoc/internal/cli"
	"os"
)

func main() {
	if err := cli.NewRootCommand().Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
