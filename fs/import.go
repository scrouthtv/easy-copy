package fs

import "path/filepath"

// CreateFS creates a file system,
// that contains all folders ("/" ending)
// and files specified in the list.
func CreateFS(list []string) *MockFS {
	fs := NewFs()

	for _, line := range list {
		create(fs, line)
	}

	return fs
}

func create(fs *MockFS, line string) {
	line = filepath.Clean(line)
	paren, rest, _ := fs.Root.resolve(line)
	_ = paren
	_ = rest
}
