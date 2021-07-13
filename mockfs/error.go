package mockfs

type ErrOperationNotSupported struct {
	Op string
}

func (e *ErrOperationNotSupported) Error() string {
	return "operation not supported: " + e.Op
}

type ErrNotADirectory struct {
	Path string
}

func (e *ErrNotADirectory) Error() string {
	return "not a directory: " + e.Path
}
