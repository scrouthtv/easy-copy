package files

import (
	"easy-copy/common"
	"easy-copy/device"
	"easy-copy/ui"
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
	devA, errA := device.GetDevice(pathA)
	devB, errB := device.GetDevice(pathB)

	if errA != nil {
		ui.Warns <- errA
		return false
	}

	if errB != nil {
		ui.Warns <- errB
		return false
	}

	return devA.Equal(devB)
}
