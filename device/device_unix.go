// +build linux freebsd openbsd netbsd dragonfly darwin

package device

import (
	"fmt"
	"os"

	"golang.org/x/sys/unix"
)

type unixDevice struct {
	id    uint64
	afile string /* any file on this device */
}

// GetDevice finds the device that hosts the file at the specified path.
// If an error occurs, it is pushed to the modules' error stack.
func GetDevice(path string) Device {
	var stat unix.Stat_t

	err := unix.Stat(path, &stat)
	if err != nil {
		pushError(err)
		return nil
	}

	return &unixDevice{stat.Dev, path}
}

func (d *unixDevice) Usage() SpaceUsage {
	var stat unix.Statfs_t

	err := unix.Statfs(d.afile, &stat)
	if err != nil {
		pushError(err)
		return UnknownUsage
	}

	var free uint64
	if isElevated() {
		free = stat.Bfree * uint64(stat.Bsize)
	} else {
		free = stat.Bavail * uint64(stat.Bsize)
	}

	return SpaceUsage{
		Total: stat.Blocks * uint64(stat.Bsize),
		Free:  free,
	}
}

func isElevated() bool {
	return os.Geteuid() == 0
}

func (d *unixDevice) Name() string {
	return fmt.Sprintf("%d:%d", unix.Major(d.id), unix.Minor(d.id))
}

func (d *unixDevice) Equal(other Device) bool {
	o, ok := other.(*unixDevice)
	if !ok {
		return false
	}

	return d.id == o.id
}
