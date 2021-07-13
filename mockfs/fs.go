package mockfs

import "io/fs"

type MockFS struct {
	Root *MockFolder
}

// MockEntry groups all information about a file.
type MockEntry interface {
	File
	fs.FileInfo
	fs.DirEntry
}
