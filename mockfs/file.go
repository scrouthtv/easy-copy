package mockfs

import (
	"io"
	"io/fs"
	"syscall"
	"time"
)

type MockFile struct {
	name     string
	contents []byte
	position int
	modTime  time.Time
}

func (f *MockFile) Chdir() error {
	return &ErrOperationNotSupported{Op: "chdir"}
}

func (f *MockFile) Chmod(mode fs.FileMode) error {
	return nil
}

func (f *MockFile) Chown(uid, gid int) error {
	return nil
}

func (f *MockFile) Close() error {
	return nil
}

func (f *MockFile) Fd() uintptr {
	return 7 // FIXME: this is a placeholder
}

func (f *MockFile) Name() string {
	return f.name
}

func (f *MockFile) Read(b []byte) (int, error) {
	if f.position >= len(f.contents) {
		return 0, io.EOF
	}

	n := len(b)
	if f.position+n > len(f.contents) {
		n = len(f.contents) - f.position
	}

	copy(b, f.contents[f.position:f.position+n])
	f.position += n
	return n, nil
}

func (f *MockFile) ReadAt(b []byte, off int64) (int, error) {
	if off >= int64(len(f.contents)) {
		return 0, io.EOF
	}

	n := len(b)
	if off+int64(n) > int64(len(f.contents)) {
		n = len(f.contents) - int(off)
	}

	copy(b, f.contents[off:int(off)+n])
	return n, nil
}

func (f *MockFile) ReadDir(count int) ([]fs.DirEntry, error) {
	return nil, &ErrNotADirectory{f.name}
}

func (f *MockFile) ReadFrom(r io.Reader) (int64, error) {
	return 0, nil // FIXME what does this do?
}

func (f *MockFile) Readdir(count int) ([]fs.FileInfo, error) {
	return nil, &ErrNotADirectory{f.name}
}

func (f *MockFile) Readdirnames(count int) ([]string, error) {
	return nil, &ErrNotADirectory{f.name}
}

func (f *MockFile) Seek(offset int64, whence int) (int64, error) {
	switch whence {
	case io.SeekStart:
		f.position = int(offset)
	case io.SeekCurrent:
		f.position += int(offset)
	case io.SeekEnd:
		f.position = len(f.contents) + int(offset)
	default:
		return 0, &ErrOperationNotSupported{Op: "seek, whence: " + string(whence)}
	}

	return int64(f.position), nil
}

func (f *MockFile) SetDeadline(t time.Time) error {
	return nil
}

func (f *MockFile) SetReadDeadline(t time.Time) error {
	return nil
}

func (f *MockFile) SetWriteDeadline(t time.Time) error {
	return nil
}

func (f *MockFile) Stat() (fs.FileInfo, error) {
	return f, nil
}

func (f *MockFile) Sync() error {
	return nil
}

func (f *MockFile) SyscallConn() (syscall.RawConn, error) {
	return nil, nil
}

func (f *MockFile) Truncate(size int64) error {
	if size < 0 {
		return &ErrOperationNotSupported{Op: "truncate with negative size"}
	}

	f.contents = f.contents[:size]
	return nil
}

func (f *MockFile) Write(b []byte) (int, error) {
	if f.position >= len(f.contents) {
		f.contents = append(f.contents, b...)
	} else {
		f.contents = append(f.contents[:f.position], b...) // TODO should this overwrite f.contents[f.position:]?
		f.position += len(b)
	}

	return len(b), nil
}

func (f *MockFile) WriteAt(b []byte, off int64) (int, error) {
	if off >= int64(len(f.contents)) {
		f.contents = append(f.contents, b...)
	} else {
		f.contents = append(f.contents[:off], b...)
	}

	return len(b), nil
}

func (f *MockFile) WriteString(s string) (int, error) {
	return f.Write([]byte(s))
}

// implementation of fs.FileInfo

func (f *MockFile) Size() int64 {
	return int64(len(f.contents))
}

func (f *MockFile) Mode() fs.FileMode {
	return 0o644
}

func (f *MockFile) ModTime() time.Time {
	return f.modTime
}

func (f *MockFile) IsDir() bool {
	return false
}

func (f *MockFile) Sys() interface{} {
	return nil
}

// implementation of fs.DirEntry
func (f *MockFile) Info() (fs.FileInfo, error) {
	return f, nil
}

func (f *MockFile) Type() fs.FileMode {
	return 0
}
