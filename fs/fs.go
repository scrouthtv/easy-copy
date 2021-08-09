package fs

import (
	"easy-copy/common"
	"errors"
	"io/fs"
	"path/filepath"
	"strings"
	"time"
)

type MockFS struct {
	Root *MockFolder
}

func (fs *MockFS) Resolve(path string) (MockEntry, error) {
	if path[0] == '/' {
		f, _, err := fs.Root.resolve(filepath.Clean(path[1:]))
		if err != nil {
			errFNF := &ErrFileNotFound{}
			if errors.As(err, &errFNF) {
				errFNF.Path = "/" + errFNF.Path
				return nil, errFNF
			}

			return nil, err
		}

		return f, nil
	} else {
		return nil, &ErrFileNotFound{Path: path}
	}
}

// Rewind rewinds all foldder iterators to the beginning.
// This is required after subdirectories on a folder have been read.
func (fs *MockFS) Rewind() {
	fs.Root.walkF(func(folder *MockFolder) {
		folder.itpos = 0
	})
}

func (fs *MockFS) Open(name string) (common.File, error) {
	return fs.Resolve(name)
}

func (fs *MockFS) Create(path string) (common.File, error) {
	f, rest, err := fs.Root.resolve(filepath.Clean(path)[1:])
	if err == nil {
		return f, nil
	}

	if rest == filepath.Base(path) {
		paren, ok := f.(*MockFolder)
		if !ok {
			return nil, &ErrNotADirectory{Path: path}
		}

		file := NewFile(filepath.Base(path))
		paren.AddFile(file)

		return file, nil
	} else {
		return nil, &ErrFileNotFound{Path: path}
	}
}

func (fs *MockFS) MkdirAll(path string, mode fs.FileMode) (err error) {
	// TODO mode is currently ignored
	path = filepath.Clean(path[1:])
	if path[len(path)-1] != filepath.Separator {
		path += string(filepath.Separator)
	}

	create(fs, path, "")

	return nil
}

func (fs *MockFS) Stat(name string) (fs.FileInfo, error) {
	return fs.Resolve(name)
}

func (fs *MockFS) Lstat(name string) (fs.FileInfo, error) {
	return fs.Resolve(name)
}

func (fs *MockFS) parentFolder(path string) (*MockFolder, string, error) {
	paren := filepath.Clean(path)
	idx := strings.LastIndex(paren, string(filepath.Separator))

	if idx <= 0 {
		return fs.Root, "", nil
	}

	if idx == len(paren)-1 {
		idx = strings.LastIndex(paren[:len(paren)-1], string(filepath.Separator))
	}

	parenEntry, rest, err := fs.Root.resolve(paren[:idx])
	if err != nil {
		return nil, "", err
	}

	parenFolder, ok := parenEntry.(*MockFolder)
	if !ok {
		return nil, "", &ErrFileNotFound{Path: path}
	}

	return parenFolder, rest, nil
}

func (fs *MockFS) Rename(old, new string) error {
	return nil // TODO
}

func (fs *MockFS) RemoveAll(path string) error {
	if path == "" || path == "/" {
		fs.Root = NewFolder("")
		return nil
	}

	parenFolder, rest, err := fs.parentFolder(path)
	if err != nil {
		return err
	}

	return parenFolder.RemoveSub(rest)
}

// Tree returns a string representation of the folder structure.
func (fs *MockFS) Tree() []string {
	return fs.Root.tree(0)
}

func (f *MockFolder) tree(depth int) []string {
	s := make([]string, 0, 2*len(f.subfolders)+len(f.files))

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

		if len(file.contents) > 0 {
			s = append(s, prefix+file.Name()+" : "+string(file.contents))
		} else {
			s = append(s, prefix+file.Name())
		}
	}

	return s
}

// MockEntry groups all information about a file.
type MockEntry interface {
	common.File
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
