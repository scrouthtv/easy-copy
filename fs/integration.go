package fs

import (
	"io"
	"io/fs"
	"os"
	"syscall"
	"time"
)

type Integration interface {
	Open(string) (File, error)
	Stat(string) (fs.FileInfo, error)
	Lstat(string) (fs.FileInfo, error)
}

type SysInt struct{}

func (s SysInt) Open(name string) (File, error) {
	return os.Open(name)
}

func (s SysInt) Stat(name string) (os.FileInfo, error) {
	return os.Stat(name)
}

func (s SysInt) LStat(name string) (fs.FileInfo, error) {
	return os.Lstat(name)
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
