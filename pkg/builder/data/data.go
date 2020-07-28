package data

import "fmt"

type Object struct {
	Name string
}

func New() *Object {
	return &Object{}
}

func (o *Object) WithName(name string) *Object {
	o.Name = name
	return o
}

func (o *Object) String() string {
	return fmt.Sprintf("%s chart", o.Name)
}
