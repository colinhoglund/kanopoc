package traefik

import "kanopoc/pkg/builder/data"

type Config struct {
	name string
	data *data.Object
}

func New() *Config {
	c := &Config{name: "traefik"}
	c.data = data.New().WithName(c.name)
	return c
}

func (t *Config) ReleaseName() string {
	return t.name
}

func (t *Config) Chart() string {
	return t.data.String()
}
