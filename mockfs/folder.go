package mockfs

type MockFolder struct {
	subfolders []*MockFolder
	files      []*MockFile
}

func (f *MockFolder) Chdir() error {
	return &ErrOperationNotSupported{Op: "chdir"}
}

func (f *MockFolder) Chmod(mode os.FileMode) error {
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
	return 0, &ErrNotAFile(f.name)
}

func (f *MockFolder) ReadAt(b []byte, off int64) (int, error) {
	return 0, &ErrNotAFile(f.name)
}

func (f *MockFile) ReadDir(count int) ([]os.FileInfo, error) {
	// TODO 
}

func (f *MockFile) ReadFrom(r io.Reader) (int64, error) {
	return 0, &ErrNotAFile(f.name)
}

func (f *MockFile) Readdir(count int) ([]os.FileInfo, error) {
	// TODO
}

func (f *MockFile) Readdirames(count int) ([]os.FileInfo, error) {
	// TODO
}

func (f *MockFile) Seek(offset int64, whence int) (int64, error) {
	return 0, &ErrNotAFile{f.name}
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

func (f *MockFile) Stat() (os.FileInfo, error) {
	return f, nil
}

func (f *MockFile) Sync() error {
	return nil
}

func (f *MockFile) SyscallConn() (interface{}, error) {
	return nil, nil
}

func (f *MockFile) Truncate(size int64) error {
	return &ErrNotAFile{f.name}
}

func (f *MockFile) Write(b []byte) (int, error) {
	return 0, &ErrNotAFile{f.name}
}

func (f *MockFile) WriteAt(b []byte, off int64) (int, error) {
	return 0, &ErrNotAFile{f.name}
}

func (f *MockFile) WriteString(s string) (int, error) {
	return 0, &ErrNotAFile{f.name}
}

// implementation of fs.FileInfo

func (f *MockFile) Size() int64 {
	// TODO
}

func (f *MockFile) Mode() os.FileMode {
	return 0o755
}

func (f *MockFile) ModTime() time.Time {
	// TODO
}

func (f *MockFile) IsDir() bool {
	return true
}

func (f *MockFile) Sys() interface{} {
	return nil
}
