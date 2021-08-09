package fs

import (
	"strings"
	"testing"
)

func TestCreateSingleFile(t *testing.T) {
	is := NewFs()

	f, err := is.Create("/test.txt")
	if err != nil {
		t.Error(err)
	}

	_, ok := f.(*MockFile)
	if !ok {
		t.Error("Created test.txt should be a File")
	}

	err = is.MkdirAll("/foo/bar/quz", 0o755)
	if err != nil {
		t.Error(err)
	}

	err = is.MkdirAll("/foo/baz", 0o755)
	if err != nil {
		t.Error(err)
	}

	err = is.MkdirAll("/foo/", 0o755)
	if err != nil {
		t.Error(err)
	}

	should := CreateFS([]string{
		"test.txt",
		"foo/bar/quz/",
		"foo/baz/",
	})

	ok, bad := is.Equal(should)
	if !ok {
		t.Error("FS are not equal:", bad)
	}

	t.Log("should:")
	t.Log(strings.Join(should.Tree(), "\n"))
	t.Log("is:")
	t.Log(strings.Join(is.Tree(), "\n"))
}
