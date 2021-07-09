package main

import (
	"easy-copy/device"
	"easy-copy/files"
	"easy-copy/flags"
)

func setOptimalBuffersize() {
	dev := device.GetDevice(flags.Current.Target())
	if dev == nil {
		return
	}

	files.SetBuffersize(dev.OptimalBuffersize())
}

func isSameDevice(pathA string, pathB string) bool {
	devA := device.GetDevice(pathA)
	devB := device.GetDevice(pathB)

	return devA.Equal(devB)
}
