package tasks

import (
	"path/filepath"
	"testing"
)

func TestRemoveFirst(t *testing.T) {
	f := "mypath/subfolder/folder/f"
	exp := filepath.Clean("subfolder/folder/f")

	if removeFirst(f) != exp {
		t.Errorf("removing first from %s gave %s, expected %s",
			f, removeFirst(f), exp)
	}
}
