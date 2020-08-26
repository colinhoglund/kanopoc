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

	moduleMap := map[string]Releaser{
		"traefik": traefik.New(),
		"drone":   drone.New(),
	}

	// configure modules
	for path, mod := range moduleMap {
		if err := c.config.Get(path).Scan(mod); err != nil {
			return err
		}
	}

	// apply modules
	for _, mod := range moduleMap {
		c.client.Apply(mod.ReleaseName(), mod.Chart())
	}

	return nil
}
