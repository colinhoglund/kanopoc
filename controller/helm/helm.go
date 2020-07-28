package helm

import helmprovider "kanopoc/provider/helm"

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

func (c *Controller) Apply(r Releaser) {
	c.client.Apply(r.ReleaseName(), r.Chart())
}

func (c *Controller) Delete(r Releaser) {
	c.client.Delete(r.ReleaseName())
}
