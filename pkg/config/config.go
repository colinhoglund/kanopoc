package config

import (
	"encoding/json"

	simplejson "github.com/bitly/go-simplejson"
	readerjson "github.com/micro/go-micro/v3/config/reader/json"

	"github.com/colinhoglund/kanopoc/pkg/config/hierarchy"
	"github.com/micro/go-micro/v3/config"
	"github.com/micro/go-micro/v3/config/reader"
	"github.com/micro/go-micro/v3/config/source/file"
	"github.com/micro/go-micro/v3/config/source/memory"
)

// Global represents global (shared between modules) configuration for Kanopy clusters
type Global struct {
	Domain string `json:"domain"`
}

// Config implementation for managing hierarchical configuration
type Config struct {
	config.Config
}

// NewFromHierarchy load data sources from a hierarchy object
func NewFromHierarchy(h *hierarchy.Hierarchy) (*Config, error) {
	// Gather Option funcs to pass to Config constructor.
	// This method of loading collects all Sources into Config.Options().Source
	// to allow slice merging methods to iterate over all source files.
	// https://github.com/micro/go-micro/blob/v2.5.0/config/options.go#L19
	var opts []config.Option

	files, err := h.Files()
	if err != nil {
		return nil, err
	}

	for _, f := range files {
		opts = append(opts, config.WithSource(file.NewSource(file.WithPath(f))))
	}

	return newConfig(opts...)
}

// NewFromObjects loads data sources from the provided interface objects
func NewFromObjects(objects ...interface{}) (*Config, error) {
	var opts []config.Option
	for _, o := range objects {
		data, err := json.Marshal(o)
		if err != nil {
			return nil, err
		}

		opts = append(opts, config.WithSource(memory.NewSource(memory.WithJSON(data))))
	}

	return newConfig(opts...)
}

func newConfig(opts ...config.Option) (*Config, error) {
	// disable micro's default env var interpolation
	opts = append(opts, config.WithReader(readerjson.NewReader(reader.WithDisableReplaceEnvVars())))

	// NewConfig doesnt actually handle errors so Init is called separately to avoid panics
	// https://github.com/micro/go-micro/pull/1648#issuecomment-634818321
	c, _ := config.NewConfig(opts...)
	if err := c.Init(opts...); err != nil {
		return nil, err
	}

	return &Config{c}, nil
}

// SliceMerged uses a lookup path to return a list of items from all sources
func (c *Config) SliceMerged(path ...string) ([]interface{}, error) {
	var slice []interface{}

	for _, source := range c.Options().Source {
		tempConfig, err := config.NewConfig(config.WithSource(source))
		if err != nil {
			return nil, err
		}

		// use go-simplejson (used in underlying micro reader interface) to get interface slices
		j, err := simplejson.NewJson(tempConfig.Get(path...).Bytes())
		if err != nil {
			return nil, err
		}

		slice = append(slice, j.MustArray([]interface{}{})...)
	}

	return slice, nil
}

// Global scans the current configuration into a Global
func (c *Config) Global() (Global, error) {
	g := Global{}

	if err := c.Scan(&g); err != nil {
		return Global{}, err
	}

	return g, nil
}
