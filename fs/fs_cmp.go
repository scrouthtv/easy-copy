package fs

import "path/filepath"

// Equal compares if fs and other
// contain the same file paths
// and those files have the same contents.
// It does not compare file attributes or mod time.
// If the fs don't match, the first folder that is missing
// a file is returned.
// If files only differ, the first non-matching file path is returned.
func (fs *MockFS) Equal(other *MockFS) (bool, string) {
	return fs.equal(fs.Root, other.Root, "")
}

func (fs *MockFS) equal(a, b *MockFolder, prefix string) (bool, string) {
	prefix += a.Name() + string(filepath.Separator)

	if len(a.subfolders) != len(b.subfolders) {
		return false, prefix
	}

	if len(a.files) != len(b.files) {
		return false, prefix
	}

	for _, Asub := range a.subfolders {
		_, Bsub, err := b.getSubfolder(Asub.Name())
		if err != nil {
			return false, prefix + Asub.Name() + string(filepath.Separator)
		}

		ok, badpath := fs.equal(Asub, Bsub, prefix)
		if !ok {
			return false, badpath
		}
	}

	for _, Afile := range a.files {
		_, Bfile, err := b.getFile(Afile.Name())
		if err != nil {
			return false, prefix + Afile.Name()
		}

		if !fs.fileEqual(Afile, Bfile) {
			return false, prefix + Afile.Name()
		}
	}

	return true, ""
}

func (fs *MockFS) fileEqual(a, b *MockFile) bool {
	return a.Name() == b.Name() &&
		string(a.contents) == string(b.contents)
}
