package mockfs

import (
	"io"
	"io/fs"
	"os"
	"syscall"
	"time"
)

type Opener func(name string) (File, error)

var OsOpener Opener = func(name string) (File, error) {
	return os.Open(name)
}

type File interface {
	Chdir() error
	Chmod(mode os.FileMode) error
	Chown(uid, gid int) error
	Close() error
	Fd() uintptr
	Name() string
	Read(b []byte) (n int, err error)
	ReadAt(b []byte, off int64) (n int, err error)
	ReadDir(n int) (fi []fs.DirEntry, err error)
	ReadFrom(r io.Reader) (n int64, err error)
	Readdir(n int) (fi []fs.FileInfo, err error)
	Readdirnames(n int) (names []string, err error)
	Seek(offset int64, whence int) (int64, error)
	SetDeadline(timeout time.Time) error
	SetReadDeadline(timeout time.Time) error
	SetWriteDeadline(timeout time.Time) error
	Stat() (os.FileInfo, error)
	Sync() error
	SyscallConn() (syscall.RawConn, error)
	Truncate(size int64) error
	Write(b []byte) (n int, err error)
	WriteAt(b []byte, off int64) (n int, err error)
	WriteString(s string) (ret int, err error)
}
