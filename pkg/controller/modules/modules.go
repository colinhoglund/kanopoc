package modules

import (
	"kanopoc/pkg/config"
	"kanopoc/pkg/module/drone"
	"kanopoc/pkg/module/traefik"
	"kanopoc/pkg/provider/helm"
)

type Controller struct {
	client *helm.Client
	config *config.Config
}

func New(h *helm.Client, c *config.Config) *Controller {
	return &Controller{client: h, config: c}
}

func (c *Controller) Dump() {
	c.client.Dump()
}

func (c *Controller) Apply() error {
	t := traefik.New()
	if err := c.config.Get("traefik").Scan(t); err != nil {
		return err
	}

	d := drone.New()
	if err := c.config.Get("drone").Scan(d); err != nil {
		return err
	}

	modules := []Releaser{
		t,
		d,
	}

	for _, mod := range modules {
		c.client.Apply(mod.ReleaseName(), mod.Chart())
	}

	return nil
}
