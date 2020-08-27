package config

import (
	"io/ioutil"
	"path/filepath"
	"reflect"
	"testing"

	"github.com/colinhoglund/kanopoc/pkg/config/hierarchy"
	"github.com/ghodss/yaml"
)

// create a list of Config objects using different constructors
func testConfigs(t *testing.T) []*Config {
	var configs []*Config

	configFile := "testdata/config.yaml"
	dataDir := "data"
	loadOrder := []string{
		"base.yaml",
		"test/deep/deep.yaml",
		"test/test.yaml",
	}

	configs = append(configs, newTestHierarchyConfig(t, configFile, dataDir, loadOrder...))
	configs = append(configs, newTestObjectConfig(t, configFile, dataDir, loadOrder...))

	return configs
}

func newTestHierarchyConfig(t *testing.T, configFile, dataDir string, loadOrder ...string) *Config {
	conf, err := NewFromHierarchy(hierarchy.New(configFile, dataDir, loadOrder...))
	if err != nil {
		t.Fatal("error creating config:", err)
	}

	return conf
}

// unmarshal files and create new Config from objects
func newTestObjectConfig(t *testing.T, configFile, dataDir string, files ...string) *Config {
	var objects []interface{}

	for _, f := range files {
		var fMap map[string]interface{}

		fBytes, err := ioutil.ReadFile(filepath.Join(filepath.Dir(configFile), dataDir, f))
		if err != nil {
			t.Fatal("error reading file:", err)
		}

		if err := yaml.Unmarshal(fBytes, &fMap); err != nil {
			t.Fatal("error unmarshaling:", err)
		}

		objects = append(objects, fMap)
	}

	c, err := NewFromObjects(objects...)
	if err != nil {
		t.Fatal("error loading objects:", err)
	}

	return c
}

func TestConfig(t *testing.T) {
	for _, c := range testConfigs(t) {
		tests := []struct {
			got  string
			want string
		}{
			{got: c.Get("base").String(""), want: "base"},
			{got: c.Get("test").String(""), want: "test"},
			{got: c.Get("dontinterpolate").String(""), want: "example.com/${notavar}"},
			{got: c.Get("deep").String(""), want: "deep"},
			{got: c.Get("override").String(""), want: "test"},
		}

		for _, test := range tests {
			if test.want != test.got {
				t.Errorf("want %s, got %s", test.want, test.got)
			}
		}
	}
}

func TestSliceMerged(t *testing.T) {
	for _, c := range testConfigs(t) {
		want := []interface{}{"base", "test"}

		got, err := c.SliceMerged("list")
		if err != nil {
			t.Error("error merging lists:", err)
		}

		if !reflect.DeepEqual(want, got) {
			t.Errorf("want %#v, got %#v", want, got)
		}
	}
}
