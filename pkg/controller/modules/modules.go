package modules

import (
	"kanopoc/pkg/module/drone"
	"kanopoc/pkg/module/traefik"
	"kanopoc/pkg/provider/helm"
)

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
		traefik.New(),
		drone.New(),
	}

	for _, mod := range modules {
		c.client.Apply(mod.ReleaseName(), mod.Chart())
	}
}
