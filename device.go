package main

import "easy-copy/device"

func setOptimalBuffersize() {
	dev := device.GetDevice(targetBase)
	size := dev.OptimalBuffersize()
	setBuffersize(size)
	verbSetBuffersize(size)
}
