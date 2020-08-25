package hierarchy

import (
	"reflect"
	"testing"
)

func newTestHierarchy(t *testing.T) *Hierarchy {
	file := "../testdata/config.yaml"
	values := struct{ Value string }{Value: "test"}

	h, err := NewFromTemplate(file, values)
	if err != nil {
		t.Fatal("error loading testdata:", err)
	}

	return h
}

func TestFiles(t *testing.T) {
	h := newTestHierarchy(t)

	want := []string{
		"../testdata/data/base.yaml",
		"../testdata/data/test/deep/deep.yaml",
		"../testdata/data/test/test.yaml",
	}
	got, err := h.Files()
	if err != nil {
		t.Errorf("error getting files: %s", err)
	}

	if !reflect.DeepEqual(want, got) {
		t.Errorf("want %#v, got %#v", want, got)
	}
}
