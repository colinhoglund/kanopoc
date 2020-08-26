package modules

import (
	"errors"
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

	moduleMap := map[string]interface{}{
		"traefik": traefik.New(),
		"drone":   drone.New(),
	}

	releasers, err := c.configureReleasers(moduleMap)
	if err != nil {
		return err
	}

	for _, mod := range releasers {
		c.client.Apply(mod.ReleaseName(), mod.Chart())
	}

	return nil
}

func (c *Controller) configureReleasers(moduleMap map[string]interface{}) ([]Releaser, error) {
	var releasers []Releaser

	for path, mod := range moduleMap {
		if err := c.config.Get(path).Scan(mod); err != nil {
			return nil, err
		}

		r, ok := mod.(Releaser)
		if !ok {
			return nil, errors.New("Releaser type assertion error")
		}

		releasers = append(releasers, r)
	}

	return releasers, nil
}
