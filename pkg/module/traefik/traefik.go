package traefik

import "kanopoc/pkg/builder/data"

type Config struct {
	Name string `json:"name"`
}

func New() *Config {
	return &Config{Name: "traefik"}
}

func (c *Config) ReleaseName() string {
	return c.Name
}

func (c *Config) Chart() string {
	return data.New().WithName(c.Name).String()
}
