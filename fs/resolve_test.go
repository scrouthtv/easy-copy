package fs

import (
	"path/filepath"
	"testing"
)

func TestPartResolve(t *testing.T) {
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

	tree := fs.Tree()
	for _, l := range tree {
		t.Log(l)
	}

	_, part, _ := fs.Root.resolve(filepath.Clean("bar/baz/c"))
	if part != "" {
		t.Error("Expected part to be empty, is", part)
	}

	_, part, _ = fs.Root.resolve(filepath.Clean("foo/bar/a"))
	if part != "bar/a" {
		t.Error("Expected part to be 'a', is", part)
	}

	_, part, _ = fs.Root.resolve(filepath.Clean("bar/quz/f"))
	if part != "quz/f" {
		t.Error("Expected part to be 'quz/f', is", part)
	}

	_, part, _ = fs.Root.resolve(filepath.Clean("g"))
	if part != "g" {
		t.Error("Expected part to be 'g', is", part)
	}
}

func TestResolve(t *testing.T) {
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

	tree := fs.Tree()
	for _, l := range tree {
		t.Log(l)
	}

	f, err := fs.Resolve("/bar/baz/c")
	if err != nil {
		t.Error("/bar/baz/c", err)
	}
	if f != c {
		t.Error("/bar/baz/c Got wrong file")
	}

	f, err = fs.Resolve("/foo/a")
	if err != nil {
		t.Error("/foo/a", err)
	}
	if f != a {
		t.Error("/foo/a Got wrong file")
	}

	openf, err := fs.Open("/foo/a")
	if err != nil {
		t.Error("/foo/a", err)
	}
	if openf != a {
		t.Error("/foo/a Got wrong file")
	}
	statf, err := fs.Stat("/foo/a")
	if err != nil {
		t.Error("/foo/a", err)
	}
	if statf != a {
		t.Error("/foo/a Got wrong file")
	}
	lstatf, err := fs.Lstat("/foo/a")
	if err != nil {
		t.Error("/foo/a", err)
	}
	if lstatf != a {
		t.Error("/foo/a Got wrong file")
	}

	f, err = fs.Resolve("/bar/")
	if err != nil {
		t.Error("/bar/", err)
	}
	if f != bar {
		t.Error("/bar/ Got wrong file")
	}
}
