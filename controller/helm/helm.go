package helm

import (
	helmprovider "kanopoc/provider/helm"

	"kanopoc/module/drone"
	"kanopoc/module/traefik"
)

type Releaser interface {
	ReleaseName() string
	Chart() string
}

type Controller struct {
	client *helmprovider.Client
}

func New() *Controller {
	c := helmprovider.New()
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
