package traefik

import (
	"fmt"
	"kanopoc/pkg/builder/data"
	"kanopoc/pkg/config"
)

type Config struct {
	Name   string `json:"name"`
	Domain string `json:"domain"`
}

func New(conf config.Global) *Config {
	return &Config{
		Name:   "traefik",
		Domain: fmt.Sprintf("traefik.%s", conf.Domain),
	}
}

func (c *Config) ReleaseName() string {
	return c.Name
}

func (c *Config) Chart() string {
	return data.New().WithName(c.Name).AddValue(c.Domain).String()
}
