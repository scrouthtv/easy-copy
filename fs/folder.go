package fs

import (
	"path/filepath"
	"strings"
)

type MockFolder struct {
	name       string
	subfolders []*MockFolder
	files      []*MockFile
	itpos      int
}

// next advances returns the current element ( = folder or file)
// and advances the iterator by 1. If no more elements are available,
// nil is returned.
func (f *MockFolder) next() MockEntry {
	if f.itpos >= len(f.subfolders)+len(f.files) {
		return nil
	}

	defer func() { f.itpos++ }()

	if f.itpos < len(f.subfolders) {
		return f.subfolders[f.itpos]
	}

	return f.files[f.itpos-len(f.subfolders)]
}

// walk calls the consumer on every file in this folder and subfolders.
func (f *MockFolder) walk(consumer func(f *MockFile)) {
	for _, sub := range f.subfolders {
		sub.walk(consumer)
	}

	for _, file := range f.files {
		consumer(file)
	}
}

func (f *MockFolder) walkF(consumer func(f *MockFolder)) {
	consumer(f)

	for _, sub := range f.subfolders {
		sub.walkF(consumer)
	}
}

// resolve recursively searches for the given path in this folder
// It returns the last subfolder that is found and optionally
// the subpath that couldn't be found and an error
// explaining why.
//
// The path must be separated with filepath.Separator.
func (f *MockFolder) resolve(path string) (MockEntry, string, error) {
	base := path
	rest := ""
	idx := strings.IndexRune(path, filepath.Separator)

	if idx != -1 {
		base = path[:idx]
		rest = path[idx+1:]
	}

	for _, sub := range f.subfolders {
		if sub.Name() == base {
			if rest == "" {
				return sub, "", nil
			}

			return sub.resolve(rest)
		}
	}

	for _, file := range f.files {
		if file.Name() == base {
			if rest == "" {
				return file, "", nil
			}

			return file, rest, &ErrNotADirectory{Path: path}
		}
	}

	return f, path, &ErrFileNotFound{Path: path}
}

func (f *MockFolder) RemoveSub(name string) error {
	idx, _, _ := f.getSubfolder(name)
	if idx != -1 {
		f.removeSubFolder(idx)
		return nil
	}

	idx, _, _ = f.getFile(name)
	if idx != -1 {
		f.removeFile(idx)
		return nil
	}

	return &ErrFileNotFound{Path: name}
}

func (f *MockFolder) getSubfolder(name string) (int, *MockFolder, error) {
	for i, sub := range f.subfolders {
		if sub.Name() == name {
			return i, sub, nil
		}
	}

	return -1, nil, &ErrFileNotFound{Path: name}
}

func (f *MockFolder) getFile(name string) (int, *MockFile, error) {
	for i, v := range f.files {
		if v.Name() == name {
			return i, v, nil
		}
	}

	return -1, nil, &ErrFileNotFound{Path: name}
}

func (f *MockFolder) removeSubFolder(idx int) {
	if idx == 0 {
		f.subfolders = f.subfolders[1:]
	} else if idx == len(f.subfolders)-1 {
		f.subfolders = f.subfolders[:len(f.subfolders)-1]
	} else {
		f.subfolders = append(f.subfolders[:idx], f.subfolders[idx+1:]...)
	}
}

func (f *MockFolder) removeFile(idx int) {
	if idx == 0 {
		f.files = f.files[1:]
	} else if idx == len(f.files)-1 {
		f.files = f.files[:len(f.files)-1]
	} else {
		f.files = append(f.files[:idx], f.files[idx+1:]...)
	}
}

func (f *MockFolder) AddFolder(folder *MockFolder) {
	f.subfolders = append(f.subfolders, folder)
}

func (f *MockFolder) AddFile(file *MockFile) {
	f.files = append(f.files, file)
}
