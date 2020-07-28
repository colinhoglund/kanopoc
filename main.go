package main

import (
	"kanopoc/controller/helm"
	"kanopoc/module/traefik"
)

func main() {
	c := helm.New()
	t1 := traefik.New("t1", "chart data")
	t2 := traefik.New("t2", "chart dataz")

	c.Dump()
	c.Apply(t1)
	c.Dump()
	c.Apply(t2)
	c.Dump()
	c.Delete(t1)
	c.Dump()
}
