package main

import (
	"easy-copy/color"
	"easy-copy/flags"
	"easy-copy/flags/impl"
	"easy-copy/iterator"
	"easy-copy/progress"
	"easy-copy/ui"
	"time"
)

var iteratorDone, done bool = false, false

var sources []string
var nodelete []string

var mode int = -1

const (
	// ModeCopy indicates that the files should only be copied.
	ModeCopy = iota

	// ModeMove indicates that the files should be moved.
	ModeMove

	// ModeRemove indicates that the specified files should be deleted.
	ModeRemove
)

// Maybe these are too small:
// uint64 goes up to 18446744073709551615
// or 2097152 TB

func main() {
	flags.Current = impl.New()

	color.Init(color.AutoColors())

	flags.Current.LoadConfig()

	flags.Current.ParseLine()

	if flags.Current.Verbosity() >= flags.VerbInfo {
		ui.PrintVersion()
		flags.VerbFlags()
	}

	//go setOptimalBuffersize()

	progress.Start = time.Now()

	go iterator.Iterate()
	//go speedLoop()
	//go watchdog()

	//copyLoop()

	//printSummary()

	time.Sleep(1 * time.Minute)
}
