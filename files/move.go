package files

import (
	"easy-copy/common"
	"easy-copy/device"
)

func Move(source string, dest string) error {
	if isSameDevice(source, dest) {
		err := common.FileAdapter.Rename(source, dest)
		if err == nil { // native move successful
			return nil
		}
	}

	// move "by hand": copy + delete
	s, err := common.FileAdapter.Open(source)
	if err != nil {
		return err
	}

	d, err := common.FileAdapter.Create(dest)
	if err != nil {
		return err
	}

	CopyFile(s, d)
	Syncdel(&[]string{source})

	return nil
}

func isSameDevice(pathA string, pathB string) bool {
	devA := device.GetDevice(pathA)
	devB := device.GetDevice(pathB)

	return devA.Equal(devB)
}
