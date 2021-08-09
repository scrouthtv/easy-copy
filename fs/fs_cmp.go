package fs

// Equal compares if fs and other
// contain the same file paths
// and those files have the same contents.
// It does not compare file attributes or mod time.
func (fs *MockFS) Equal(other *MockFS) bool {
	return fs.equal(fs.Root, other.Root)
}

func (fs *MockFS) equal(a, b *MockFolder) bool {
	if len(a.subfolders) != len(b.subfolders) {
		return false
	}
	if len(a.files) != len(b.files) {
		return false
	}

	for _, Asub := range a.subfolders {
		_, Bsub, err := b.getSubfolder(Asub.Name())
		if err != nil {
			return false
		}

		if !fs.equal(Asub, Bsub) {
			return false
		}
	}

	for _, Afile := range a.files {
		_, Bfile, err := b.getFile(Afile.Name())
		if err != nil {
			return false
		}

		if !fs.file_equal(Afile, Bfile) {
			return false
		}
	}

	return true
}

func (fs *MockFS) file_equal(a, b *MockFile) bool {
	return a.Name() == b.Name() &&
		string(a.contents) == string(b.contents)
}
