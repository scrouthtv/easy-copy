package iterator

import (
	"easy-copy/tasks"
	"testing"
)

func TestBasicIterate(t *testing.T) {
	tasks.SetBase("/mnt")
	err := AddAll("/tmp/asdf")
	if err != nil {
		t.Error(err)
	}

	tasks.PrintTasks()
}
