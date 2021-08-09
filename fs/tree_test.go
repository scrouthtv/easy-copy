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
	a.contents = []byte("this is a")
	b.contents = []byte("i am b")
	c.contents = []byte("guess i'm c")

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
		"│  └──a : this is a",
		"├──bar/",
		"│  ├──baz/",
		"│  │  ├──c : guess i'm c",
		"│  │  └──d",
		"│  └──b : i am b",
		"├──empty/",
		"└──e",
	}

	if !reflect.DeepEqual(is, should) {
		t.Error("different tree:")

		switch {
		case len(is) == len(should):
			for i, v := range is {
				t.Log(v, "\t\t\t", should[i])
			}
		case len(is) < len(should):
			for i, v := range is {
				t.Log(v, "\t\t\t", should[i])
			}

			for i := len(is); i < len(should); i++ {
				t.Log(should[i])
			}
		default:
			for i, v := range should {
				t.Log(is[i], "\t\t\t", v)
			}

			for i := len(should); i < len(is); i++ {
				t.Log(is[i])
			}
		}
	}
}
