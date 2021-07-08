package files

import "fmt"

// ErrWritingData is returned by the copy loop
// if not all data could be written,
// but no other is returned.
type ErrWritingData struct {
	read    int
	written int
}

func (e *ErrWritingData) Error() string {
	return fmt.Sprintf("could only write %d b out of %d b", e.written, e.read)
}
