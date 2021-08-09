package common

import (
	"os"
)

var FileAdapter IFileAdapter

func init() {
	FileAdapter = SysInt{}
}

type IFileAdapter interface {
	Open(string) (File, error)
	Stat(string) (os.FileInfo, error)
	Lstat(string) (os.FileInfo, error)
}

type SysInt struct{}

func (s SysInt) Open(name string) (File, error) {
	return os.Open(name)
}

func (s SysInt) Stat(name string) (os.FileInfo, error) {
	return os.Stat(name)
}

func (s SysInt) Lstat(name string) (os.FileInfo, error) {
	return os.Lstat(name)
}
