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
	"os"
	"strings"
	"time"
)

func main() {
	go handler.Handle()

	flags.Current = impl.New()

	color.Init(color.AutoColors())

	flags.Current.LoadConfig(os.Args)

	args, err := flags.Sep(strings.Join(os.Args, " "))
	if err != nil {
		ui.Error(err)
	}

	flags.Current.ParseLine(args)

	if flags.Current.Verbosity() >= flags.VerbInfo {
		ui.PrintVersion()
		flags.VerbFlags()
	}

	if flags.Current.Parallel() {
		runParallel()
		return
	}

	// go setOptimalBuffersize()

	progress.Start = time.Now()

	go iterator.Iterate()

	go drawLoop()

	go progress.WatchSpeed()
	go progress.Watchdog()

	tasks.CopyLoop()

	printSummary()
}

func runParallel() {
	iterator.Iterate()
	println("confirm?")
	buf := make([]byte, 8)
	os.Stdin.Read(buf) // wait for confirmation on task list
	tasks.CopyLoop()
	printSummary()
}
