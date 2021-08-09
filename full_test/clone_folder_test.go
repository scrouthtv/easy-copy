package full_test

import (
	"easy-copy/fs"
	"testing"
)

func TestCloneFolder(t *testing.T) {
	is := fs.CreateFS([]string{
		"/foo/a",
		"/foo/b",
		"/foo/c",
		"/foo/sub/empty/",
		"/foo/sub/x",
	})
	line := "cp foo bar"
	should := fs.CreateFS([]string{
		"/foo/a",
		"/foo/b",
		"/foo/c",
		"/foo/sub/empty/",
		"/foo/sub/x",
		"/bar/a",
		"/bar/b",
		"/bar/c",
		"/bar/sub/empty/",
		"/bar/sub/x",
	})

	test := NewTest(line)
	test.is = is
	test.should = should

	test.Run(t)

	t.Log("after copying:")
	tree := is.Tree()
	for _, v := range tree {
		t.Log(v)
	}
}
