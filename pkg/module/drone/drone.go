package drone

type Config struct {
	name, data string
}

func New(name, data string) *Config {
	return &Config{name, data}
}

func (t *Config) ReleaseName() string {
	return t.name
}

func (t *Config) Chart() string {
	return t.data
}
