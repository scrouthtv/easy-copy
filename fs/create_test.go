package fs

import (
	"reflect"
	"testing"
)

func TestCreate(t *testing.T) {
	create := []string{
		"foo/bar/a",
		"foo/bar/c",
		"foo/q",
		"foo/q/qwertz",
		"foo/",
		"a",
		"baz/",
		"",
	}

	fs := CreateFS(create)

	var _ Integration = fs

	is := fs.Tree()

	for _, v := range is {
		t.Log(v)
	}

	should := []string{
		"├──foo/",
		"│  ├──bar/",
		"│  │  ├──a",
		"│  │  └──c",
		"│  └──q",
		"├──baz/",
		"└──a",
	}

	if !reflect.DeepEqual(is, should) {
		t.Errorf("\n\t%v\n\t%v", is, should)
	}
}
