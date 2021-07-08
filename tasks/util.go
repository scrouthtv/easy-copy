package tasks

import (
	"path/filepath"
	"strings"
)

// removeFirst removes the first folder of a path.
func removeFirst(path string) string {
	path = filepath.Clean(path)
	idx := strings.IndexRune(path, filepath.Separator)

	if idx == -1 {
		return path
	}

	return path[idx+1:]
}
