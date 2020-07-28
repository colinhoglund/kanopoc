package main

import (
	"kanopoc/pkg/controller/modules"
)

func main() {
	c := modules.New()
	c.Dump()
	c.Apply()
	c.Dump()
	c.Apply()
	c.Dump()
}
