package tasks

import (
	"easy-copy/flags"
	"easy-copy/flags/stub"
	"testing"
)

func TestDestforCreateFolders(t *testing.T) {
	flags.Current = stub.New()

	Setup("bar", true) // create folders in target "bar/"

	is := destFor(&Path{"", "q.txt"})
	exp := "bar/q.txt"
	if is != exp {
		t.Errorf("1. Expected %s, got %s", exp, is)
	}

	is = destFor(&Path{"", "foo/f.txt"})
	exp = "bar/foo/f.txt"
	if is != exp {
		t.Errorf("1. Expected %s, got %s", exp, is)
	}
}

func TestDestforNoFolders(t *testing.T) {
	flags.Current = stub.New()

	Setup("bar", false) // don't create folders in target "bar/"

	is := destFor(&Path{"", "f.txt"})
	exp := "bar/f.txt"
	if is != exp {
		t.Errorf("1. Expected %s, got %s", exp, is)
	}
}
