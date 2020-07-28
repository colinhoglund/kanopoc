package modules

type Releaser interface {
	ReleaseName() string
	Chart() string
}
