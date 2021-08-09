package fs

import (
	"path/filepath"
	"strings"
)

// CreateFS creates a file system,
// that contains all folders ("/" ending)
// and files specified in the list.
// Folders that already exist as a file are skipped.
func CreateFS(list []string) *MockFS {
	fs := NewFs()

	for _, line := range list {

		params := strings.Split(line, " : ")
		if len(params) == 1 {
			create(fs, params[0], "")
		} else {
			create(fs, params[0], params[1])
		}
	}

	return fs
}

func create(fs *MockFS, path string, content string) {
	if path == "" {
		return
	}

	lastfolder := path[len(path)-1] == filepath.Separator

	pp, rest, err := fs.Root.resolve(filepath.Clean(path))
	if err == nil && rest == "" {
		return
	}

	paren, ok := pp.(*MockFolder)
	if !ok {
		return
	}

	split := strings.Split(rest, string(filepath.Separator))

	for i, seg := range split {
		if i < len(split)-1 {
			f := NewFolder(seg)
			paren.AddFolder(f)
			paren = f
		} else {
			if lastfolder {
				f := NewFolder(seg)
				paren.AddFolder(f)
			} else {
				f := NewFile(seg)
				paren.AddFile(f)
			}
			return
		}
	}
}
