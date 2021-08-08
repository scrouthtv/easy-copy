package fs

import (
	"io/fs"
	"path/filepath"
	"strings"
	"time"
)

type MockFS struct {
	Root *MockFolder
}

func (f *MockFS) Resolve(path string) (MockEntry, error) {
	if path[0] == '/' {
		return f.Root.resolve(filepath.Clean(path[1:]))
	} else {
		return nil, &ErrFileNotFound{path}
	}
}

func (f *MockFS) Tree() []string {
	return f.Root.tree(0)
}

func (f *MockFolder) tree(depth int) []string {
	var s []string

	p := strings.Repeat("│  ", depth)
	var prefix string

	for i, sub := range f.subfolders {
		if len(f.files) == 0 && i == len(f.subfolders)-1 {
			prefix = p + "└──"
		} else {
			prefix = p + "├──"
		}

		s = append(s, prefix+sub.Name()+"/")
		s = append(s, sub.tree(depth+1)...)
	}
	for i, file := range f.files {
		if i == len(f.files)-1 {
			prefix = p + "└──"
		} else {
			prefix = p + "├──"
		}

		s = append(s, prefix+file.Name())
	}

	return s
}

// MockEntry groups all information about a file.
type MockEntry interface {
	File
	fs.FileInfo
	fs.DirEntry
}

func NewFs() *MockFS {
	return &MockFS{
		Root: NewFolder(""),
	}
}

func NewFolder(name string) *MockFolder {
	return &MockFolder{
		name:       name,
		subfolders: nil,
		files:      nil,
		itpos:      0,
	}
}

// Creates an empty file that was last modified at the current time.
func NewFile(name string) *MockFile {
	return &MockFile{
		name:     name,
		contents: nil,
		position: 0,
		modTime:  time.Now(),
	}
}
