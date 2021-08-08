package iterator

import (
	"easy-copy/flags"
	"easy-copy/flags/stub"
	"easy-copy/tasks"
	"testing"
)

func TestIterateSingleFolder(t *testing.T) {
	flags.Current = stub.New()

	tasks.Setup("d:/tmp/target", true)
	err := add(&tasks.Path{Base: "c:/tmp/asdf/foo", Sub: ""})
	if err != nil {
		t.Error(err)
	}

	tasks.PrintTasks()
}

func TestIterateMultiFolders(t *testing.T) {
	flags.Current = stub.New()

	tasks.Setup("d:/tmp/target", true)
	err := add(&tasks.Path{Base: "c:/tmp/asdf", Sub: ""})
	if err != nil {
		t.Error(err)
	}

	tasks.PrintTasks()
}
