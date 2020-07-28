package main

import (
	"kanopoc/controller/helm"
)

func main() {
	c := helm.New()
	c.Dump()
	c.Apply()
	c.Dump()
	c.Apply()
	c.Dump()
}
