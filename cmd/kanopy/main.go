package main

import (
	"fmt"
	"os"

	"github.com/colinhoglund/kanopoc/internal/cli"
)

func main() {
	if err := cli.NewRootCommand().Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
