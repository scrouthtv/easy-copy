package main

// These constants are for now also hardcoded in the manpage.
const (
	// EasyCopyName is the name of the program.
	EasyCopyName    string = "EasyCopy"

	// EasyCopyVersion is the version of the program.
	EasyCopyVersion string = "0.4.4"
)

const (
	// folderSize is the size that should be used for progress calculation
	// for a folder.
	// It should resemble the "complexity" of creating a folder compared
	// to writing 1 byte.
	folderSize  int = 4

	// symlinkSize is the size that should be used for progress calculation
	// for a symbolic link.
	// It should resemble the "complexity" of creating a link compared
	// to writing 1 byte.
	symlinkSize int = 16
)
