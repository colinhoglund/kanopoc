package data

import "fmt"

type Object struct {
	Name   string
	Values []string
}

func New() *Object {
	return &Object{}
}

func (o *Object) WithName(name string) *Object {
	o.Name = name
	return o
}

func (o *Object) AddValue(s string) *Object {
	o.Values = append(o.Values, s)
	return o
}

func (o *Object) String() string {
	return fmt.Sprintf("Name: %s chart, Values: %v", o.Name, o.Values)
}
