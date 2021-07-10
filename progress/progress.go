package progress

const (
	// FolderSize is the size that should be used for progress calculation
	// for a folder.
	// It should resemble the "complexity" of creating a folder compared
	// to writing 1 byte.
	FolderSize int = 4

	// SymlinkSize is the size that should be used for progress calculation
	// for a symbolic link.
	// It should resemble the "complexity" of creating a link compared
	// to writing 1 byte.
	SymlinkSize int = 16
)

var (
	DoneAmount uint64 = 0
	FullAmount uint64 = 0
	DoneSize   uint64 = 0
	FullSize   uint64 = 0
)

// Maybe these are too small:
// uint64 goes up to 18446744073709551615
// or 2097152 TB

var (
	CurrentTask TaskType
	CurrentFile string
)

type TaskType uint8

const (
	TaskNone TaskType = iota
	TaskCopy
	TaskLink
	TaskDelete
	TaskMkdir
)

var (
	IteratorDone = false
	CopyDone     = false
)
