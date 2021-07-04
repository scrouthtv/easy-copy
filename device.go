package main

import "easy-copy/device"

func setOptimalBuffersize() {
	dev := device.GetDevice(targetBase)
	setBuffersize(dev.OptimalBuffersize())
}

func isSameDevice(pathA string, pathB string) bool {
	devA := device.GetDevice(pathA)
	devB := device.GetDevice(pathB)
	return devA.Equal(devB)
}
