package main

import (
	"easy-copy/color"
	"easy-copy/flags"
	"easy-copy/flags/impl"
	"easy-copy/iterator"
	"easy-copy/progress"
	"easy-copy/tasks"
	"easy-copy/ui"
	"easy-copy/ui/handler"
	"time"
)

var nodelete []string

func main() {
	go handler.Handle()

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

	go progress.WatchSpeed()
	go progress.Watchdog()

	tasks.CopyLoop()

	//printSummary()

	time.Sleep(1 * time.Minute)
}
