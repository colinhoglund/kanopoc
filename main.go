package main

import (
	"kanopoc/controller/modules"
)

func main() {
	c := modules.New()
	c.Dump()
	c.Apply()
	c.Dump()
	c.Apply()
	c.Dump()
}
