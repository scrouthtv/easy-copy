package progress

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
