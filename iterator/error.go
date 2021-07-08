package iterator

import "io/fs"

type ErrBadType struct {
	Path string
	Info fs.FileInfo
}

func (err *ErrBadType) Error() string {
	return "ignoring " + err.Path + " (" + err.Info.Mode().String() + ")"
}
