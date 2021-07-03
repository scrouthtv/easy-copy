// +build linux freebsd openbsd netbsd dragonfly darwin

package device

import "fmt"

import "golang.org/x/sys/unix"

type unixDevice struct {
	id uint64
}

func GetDevice(path string) Device {
	var stat unix.Stat_t

	unix.Stat(path, &stat)

	return &unixDevice{stat.Dev}
}

func (d *unixDevice) Usage() SpaceUsage {
	return SpaceUsage{}
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
