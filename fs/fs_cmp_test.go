package fs

import "testing"

func TestEqualFS(t *testing.T) {
	a := CreateFS([]string{
		"foo/empty/",
		"bar/a",
		"bar/b",
		"bar/c",
		"foo/x",
		"test.txt",
	})
	b := CreateFS([]string{ // same as a
		"foo/empty/",
		"bar/a",
		"bar/b",
		"bar/c",
		"foo/x",
		"test.txt",
	})
	c := CreateFS([]string{ // missing /foo/empty/
		"bar/a",
		"bar/b",
		"bar/c",
		"foo/x",
		"test.txt",
	})
	d := CreateFS([]string{ // missing /bar/b
		"foo/empty/",
		"bar/a",
		"bar/c",
		"foo/x",
		"test.txt",
	})
	e := CreateFS([]string{ // missing /test.txt
		"foo/empty/",
		"bar/a",
		"bar/b",
		"bar/c",
		"foo/x",
	})
	f := CreateFS([]string{ // extra /bar/quz/
		"foo/empty/",
		"bar/quz/",
		"bar/a",
		"bar/b",
		"bar/c",
		"foo/x",
		"test.txt",
	})
	g := CreateFS([]string{ // all of the above
		"bar/quz/",
		"bar/a",
		"bar/c",
		"foo/x",
	})

	if !a.Equal(b) {
		t.Error("a != b")
	}

	if a.Equal(c) {
		t.Error("a == c")
	}
	if a.Equal(d) {
		t.Error("a == d")
	}
	if a.Equal(e) {
		t.Error("a == e")
	}
	if a.Equal(f) {
		t.Error("a == f")
	}
	if a.Equal(g) {
		t.Error("a == g")
	}
}
