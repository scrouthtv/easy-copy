package main

import (
	"easy-copy/device"
	"easy-copy/files"
	"easy-copy/flags"
)

func setOptimalBuffersize() {
	dev, err := device.GetDevice(flags.Current.Target())
	if dev == nil || err != nil {
		return
	}

	files.SetBuffersize(dev.OptimalBuffersize())
}
