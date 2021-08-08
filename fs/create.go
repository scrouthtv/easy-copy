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
		create(fs, line)
	}

	return fs
}

func create(fs *MockFS, line string) {
	if line == "" {
		return
	}

	lastfolder := line[len(line)-1] == filepath.Separator
	pp, rest, _ := fs.Root.resolve(filepath.Clean(line))
	if rest == "" {
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
