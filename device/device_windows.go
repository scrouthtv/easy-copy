package device

import (
	"path/filepath"
	"strings"
	"unsafe"

	"golang.org/x/sys/windows"
)

type windowsDevice struct {
	volSerial uint32
	path      string
}

func GetDevice(path string) (Device, error) {
	path = filepath.Clean(path)

	//serial

	return &windowsDevice{path: path}, nil
}

func (d *windowsDevice) Equal(other Device) bool {
	o, ok := other.(*windowsDevice)
	if !ok {
		return false
	}

	return d.Name() == o.Name()
}

// Name returns the drive letter.
func (d *windowsDevice) Name() string {
	idx := strings.IndexRune(d.path, ':')
	if idx == -1 {
		return d.path
	}

	return d.path[:idx+1]
}

func (d *windowsDevice) OptimalBuffersize() int {
	// https://stackoverflow.com/a/15540773/7242251
	return 64 * 1024
}

func (d *windowsDevice) Usage() (*SpaceUsage, error) {
	var tot, free, discard uint64
	err := windows.GetDiskFreeSpaceEx((*uint16)(unsafe.Pointer(&d.path)), &free, &tot, &discard)
	if err != nil {
		return nil, err
	}

	return &SpaceUsage{Free: free, Total: tot}, nil
}
