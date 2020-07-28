package helm

import "fmt"

type release struct {
	versions int
	chart    string
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
		r.chart = chart
		r.versions++
	}
}

func (c *Client) Delete(name string) {
	delete(c.releases, name)
}

func (c *Client) Dump() {
	fmt.Println("===start===")
	for name, rel := range c.releases {
		fmt.Println(name, rel.versions, rel.chart)
	}
	fmt.Println("===end===")
}
