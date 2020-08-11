package modules

import (
	"kanopoc/pkg/module/drone"
	"kanopoc/pkg/module/traefik"
	"kanopoc/pkg/provider/helm"
)

type Controller struct {
	client *helm.Client
}

func New(h *helm.Client) *Controller {
	return &Controller{h}
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
