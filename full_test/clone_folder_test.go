package full_test

import (
	"easy-copy/fs"
	"testing"
)

func TestCloneFolder(t *testing.T) {
	is := fs.CreateFS([]string{
		"foo/a : qwertz",
		"foo/b",
		"foo/c : thisisafile",
		"foo/sub/empty/",
		"foo/sub/x : x",
	})
	should := fs.CreateFS([]string{
		"foo/a : qwertz",
		"foo/b",
		"foo/c : thisisafile",
		"foo/sub/empty/",
		"foo/sub/x : x",
		"bar/a : qwertz",
		"bar/b",
		"bar/c : thisisafile",
		"bar/sub/empty/",
		"bar/sub/x : x",
	})

	line := "cp foo bar"

	test := NewTest(line)
	test.is = is
	test.should = should

	test.Run(t)

	t.Log("after copying:")
	for _, v := range is.Tree() {
		t.Log(v)
	}

	ok, badpath := is.Equal(should)
	if !ok {
		t.Error("File System differs:", badpath)
	}
}
