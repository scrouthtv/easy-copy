package fs

import (
	"reflect"
	"testing"
)

func TestTree(t *testing.T) {
	fs := NewFs()
	foo := NewFolder("foo")
	bar := NewFolder("bar")

	fs.Root.AddFolder(foo)
	fs.Root.AddFolder(bar)

	baz := NewFolder("baz")
	bar.AddFolder(baz)

	a, b, c, d, e := NewFile("a"), NewFile("b"), NewFile("c"), NewFile("d"), NewFile("e")
	foo.AddFile(a)
	bar.AddFile(b)
	baz.AddFile(c)
	baz.AddFile(d)
	fs.Root.AddFile(e)

	empty := NewFolder("empty")
	fs.Root.AddFolder(empty)

	is := fs.Tree()

	should := []string{
		"├──foo/",
		"│  └──a",
		"├──bar/",
		"│  ├──baz/",
		"│  │  ├──c",
		"│  │  └──d",
		"│  └──b",
		"├──empty/",
		"└──e",
	}

	if !reflect.DeepEqual(is, should) {
		t.Error("different tree:")

		if len(is) == len(should) {
			for i, v := range is {
				t.Log(v, "\t\t\t", should[i])
			}
		} else if len(is) < len(should) {
			for i, v := range is {
				t.Log(v, "\t\t\t", should[i])
			}
			for i := len(is); i < len(should); i++ {
				t.Log(should[i])
			}
		} else {
			for i, v := range should {
				t.Log(is[i], "\t\t\t", v)
			}
			for i := len(should); i < len(is); i++ {
				t.Log(is[i])
			}
		}
	}
}
