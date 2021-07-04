package main

import "easy-copy/device"

func setOptimalBuffersize() {
	dev := device.GetDevice(targetBase)
	setBuffersize(dev.OptimalBuffersize())
}
