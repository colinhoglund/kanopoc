package modules

import (
	"kanopoc/module/drone"
	"kanopoc/module/traefik"
	"kanopoc/provider/helm"
)

type Releaser interface {
	ReleaseName() string
	Chart() string
}

type Controller struct {
	client *helm.Client
}

func New() *Controller {
	c := helm.New()
	return &Controller{c}
}

func (c *Controller) Dump() {
	c.client.Dump()
}

func (c *Controller) Apply() {
	modules := []Releaser{
		traefik.New("traefik", "traefik data"),
		drone.New("drone", "drone data"),
	}

	for _, mod := range modules {
		c.client.Apply(mod.ReleaseName(), mod.Chart())
	}
}
