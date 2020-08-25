package helm

import "fmt"

type release struct {
	version int
	chart   string
}

type Client struct {
	releases map[string]*release
}

func New() *Client {
	c := &Client{map[string]*release{}}
	return c
}

func (c *Client) Apply(name, chart string) {
	_, exists := c.releases[name]
	switch {
	case !exists:
		// install
		c.releases[name] = &release{1, chart}
	default:
		// upgrade
		r := c.releases[name]
		if chart != r.chart {
			r.chart = chart
			r.version++
		}
	}
}

func (c *Client) Delete(name string) {
	delete(c.releases, name)
}

func (c *Client) Dump() {
	fmt.Println("=> Dumping Releases")
	for name, rel := range c.releases {
		fmt.Printf("==> name: %s, version: %d, chart: %s\n", name, rel.version, rel.chart)
	}
}
