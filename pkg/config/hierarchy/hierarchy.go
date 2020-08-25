package hierarchy

import (
	"bytes"
	"io/ioutil"
	"os"
	"path/filepath"
	"text/template"

	"github.com/ghodss/yaml"
)

// Hierarchy stores information about the configuration hierarchy
type Hierarchy struct {
	configFile    string
	DataDirectory string   `json:"dataDirectory"`
	LoadOrder     []string `json:"loadOrder"`
}

// New returns a Hierarchy object
func New(file, dataDir string, loadOrder ...string) *Hierarchy {
	return &Hierarchy{
		configFile:    file,
		DataDirectory: dataDir,
		LoadOrder:     loadOrder,
	}
}

// NewFromTemplate creates a new hierarchy object from a passed in template file and values
func NewFromTemplate(file string, values interface{}) (*Hierarchy, error) {
	h := &Hierarchy{configFile: file}

	// parse and buffer template
	t, err := template.ParseFiles(h.configFile)
	if err != nil {
		return nil, err
	}

	var buf bytes.Buffer
	if err := t.Execute(&buf, values); err != nil {
		return nil, err
	}

	// unmarshal buffer into Hierarchy object
	if err := yaml.Unmarshal(buf.Bytes(), h); err != nil {
		return nil, err
	}

	return h, nil
}

// Files returns a list of full path string for existing configuration files
func (h *Hierarchy) Files() ([]string, error) {
	var files []string

	for _, f := range h.LoadOrder {
		fullpath := filepath.Join(filepath.Dir(h.configFile), h.DataDirectory, f)
		fs, err := buildLoadOrderRecursively(fullpath)
		if err != nil {
			return nil, err
		}
		files = append(files, fs...)
	}

	return files, nil
}

func buildLoadOrderRecursively(path string) ([]string, error) {
	var files []string

	switch fInfo, err := os.Stat(path); {
	case os.IsNotExist(err):
		// skip non-existent config files
		break
	case err != nil:
		return nil, err
	case fInfo.Mode().IsRegular():
		files = append(files, path)
	case fInfo.Mode().IsDir():
		dirFiles, err := ioutil.ReadDir(path)
		if err != nil {
			return nil, err
		}

		for _, file := range dirFiles {
			fs, err := buildLoadOrderRecursively(filepath.Join(path, file.Name()))
			if err != nil {
				return nil, err
			}

			files = append(files, fs...)
		}
	}

	return files, nil
}
