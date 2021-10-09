package fs_test

import (
	"easy-copy/fs"
	"reflect"
	"testing"
)

func TestReaddir(t *testing.T) {
	test := fs.CreateFS([]string{
		"foo/",
		"foo/a",
		"foo/b",
		"foo/sub/x",
	})

	foo, err := test.Open("/foo")
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	names, err := foo.Readdirnames(0)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	should := []string{"sub", "a", "b"}

	if !reflect.DeepEqual(should, names) {
		t.Error("readdirnames mismatch")
	}
}
