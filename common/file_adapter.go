package common

import (
	"os"
)

var FileAdapter IFileAdapter

func init() {
	FileAdapter = SysFiles{}
}

type IFileAdapter interface {
	Open(string) (File, error)
	Create(string) (File, error)
	Stat(string) (os.FileInfo, error)
	Lstat(string) (os.FileInfo, error)
	Rename(string, string) error
	RemoveAll(string) error
}

type SysFiles struct{}

func (s SysFiles) Open(name string) (File, error) {
	return os.Open(name)
}

func (s SysFiles) Create(path string) (File, error) {
	return os.Create(path)
}

func (s SysFiles) Stat(name string) (os.FileInfo, error) {
	return os.Stat(name)
}

func (s SysFiles) Lstat(name string) (os.FileInfo, error) {
	return os.Lstat(name)
}

func (s SysFiles) Rename(oldpath, newpath string) error {
	return os.Rename(oldpath, newpath)
}

func (s SysFiles) RemoveAll(path string) error {
	return os.RemoveAll(path)
}
