package full_test

import (
	"easy-copy/fs"
	"testing"
)

func TestDuplicateRenaming(t *testing.T) {
	is := fs.CreateFS([]string{
		"foo/a",
		"bar/",
	})

	should := fs.CreateFS([]string{
		"foo/a",
		"bar/b",
	})

	t.Log("before copying:")
	for _, v := range is.Tree() {
		t.Log(v)
	}

	line := "cp foo/a bar/b --debug"

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
		t.Error("fs differs:", badpath)
	}
}
