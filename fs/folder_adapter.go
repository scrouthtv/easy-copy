package fs

import (
	"io"
	"io/fs"
	"syscall"
	"time"
)

// implementation of File.

func (f *MockFolder) Chdir() error {
	return &ErrOperationNotSupported{Op: "chdir"}
}

func (f *MockFolder) Chmod(mode fs.FileMode) error {
	return nil
}

func (f *MockFolder) Chown(uid, gid int) error {
	return nil
}

func (f *MockFolder) Close() error {
	return nil
}

func (f *MockFolder) Fd() uintptr {
	return 7 // FIXME: this is a placeholder
}

func (f *MockFolder) Name() string {
	return f.name
}

func (f *MockFolder) Read(b []byte) (int, error) {
	return 0, &ErrNotAFile{Path: f.name}
}

func (f *MockFolder) ReadAt(b []byte, off int64) (int, error) {
	return 0, &ErrNotAFile{Path: f.name}
}

func (f *MockFolder) ReadDir(count int) ([]fs.DirEntry, error) {
	entries := []fs.DirEntry{}

	if count <= 0 {
		next := f.next()

		for next != nil {
			entries = append(entries, next)
			next = f.next()
		}
	} else {
		for i := 0; i < count; i++ {
			next := f.next()
			entries = append(entries, next)
		}
	}

	return entries, nil
}

func (f *MockFolder) ReadFrom(r io.Reader) (int64, error) {
	return 0, &ErrNotAFile{Path: f.name}
}

func (f *MockFolder) Readdir(count int) ([]fs.FileInfo, error) {
	entries := []fs.FileInfo{}

	if count <= 0 {
		next := f.next()

		for next != nil {
			entries = append(entries, next)
			next = f.next()
		}
	} else {
		for i := 0; i < count; i++ {
			next := f.next()
			entries = append(entries, next)
		}
	}

	return entries, nil
}

func (f *MockFolder) Readdirnames(count int) ([]string, error) {
	entries := []string{}

	if count <= 0 {
		next := f.next()

		for next != nil {
			entries = append(entries, next.Name())
			next = f.next()
		}
	} else {
		for i := 0; i < count; i++ {
			next := f.next()
			entries = append(entries, next.Name())
		}
	}

	return entries, nil
}

func (f *MockFolder) Seek(offset int64, whence int) (int64, error) {
	return 0, &ErrNotAFile{Path: f.name}
}

func (f *MockFolder) SetDeadline(t time.Time) error {
	return nil
}

func (f *MockFolder) SetReadDeadline(t time.Time) error {
	return nil
}

func (f *MockFolder) SetWriteDeadline(t time.Time) error {
	return nil
}

func (f *MockFolder) Stat() (fs.FileInfo, error) {
	return f, nil
}

func (f *MockFolder) Sync() error {
	return nil
}

func (f *MockFolder) SyscallConn() (syscall.RawConn, error) {
	return nil, nil
}

func (f *MockFolder) Truncate(size int64) error {
	return &ErrNotAFile{Path: f.name}
}

func (f *MockFolder) Write(b []byte) (int, error) {
	return 0, &ErrNotAFile{Path: f.name}
}

func (f *MockFolder) WriteAt(b []byte, off int64) (int, error) {
	return 0, &ErrNotAFile{Path: f.name}
}

func (f *MockFolder) WriteString(s string) (int, error) {
	return 0, &ErrNotAFile{Path: f.name}
}

// implementation of fs.FileInfo.

func (f *MockFolder) Size() int64 {
	var size int64 = 0

	f.walk(func(f *MockFile) {
		size += f.Size()
	})

	return size
}

func (f *MockFolder) Mode() fs.FileMode {
	return 0o755
}

// ModTime returns the newest Mod Time of all files in this folder and subfolders.
func (f *MockFolder) ModTime() time.Time {
	var modTime time.Time

	f.walk(func(f *MockFile) {
		if f.ModTime().After(modTime) { // this works bc the null value for time.Time is 1.1.1
			modTime = f.ModTime()
		}
	})

	return modTime
}

func (f *MockFolder) IsDir() bool {
	return true
}

func (f *MockFolder) Sys() interface{} {
	return nil
}

// implementation of fs.DirEntry.

func (f *MockFolder) Info() (fs.FileInfo, error) {
	return f, nil
}

func (f *MockFolder) Type() fs.FileMode {
	return fs.ModeDir
}
