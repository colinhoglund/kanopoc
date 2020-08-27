package drone

import "github.com/colinhoglund/kanopoc/pkg/builder/data"

type Config struct {
	Name string `json:"name"`
}

func New() *Config {
	return &Config{Name: "drone"}
}

func (c *Config) ReleaseName() string {
	return c.Name
}

func (c *Config) Chart() string {
	return data.New().WithName(c.Name).String()
}
