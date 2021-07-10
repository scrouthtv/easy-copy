package main

import (
	"easy-copy/flags"
	"easy-copy/progress"
	"easy-copy/ui"
	"fmt"
	"math"
	"strconv"
	"time"
)

const (
	barWidth    = 50 // in px
	maxWidth    = 80
	redrawSpeed = 100 // in ms
)

var lines int = 0

func drawLoop() {
	for !progress.CopyDone {
		for i := 0; i < lines; i++ {
			fmt.Print("\033[1A\033[2K")
		}

		lines = 0

		if flags.Current.Verbosity() > flags.VerbQuiet {
			printBar()

			unit := ui.SizeAutoUnit(float64(progress.FullSize))

			fmt.Print(ui.FormatSize(float64(progress.DoneSize), unit))
			fmt.Print(" / ")
			fmt.Print(ui.FormatSize(float64(progress.FullSize), unit))
			fmt.Println()

			lines++

			printOperation()
			fmt.Println()

			lines++

			fmt.Println()
			lines++
		}

		//printConflict()

		time.Sleep(redrawSpeed * time.Millisecond)
	}
}

func printBar() {
	var barFilled int

	if progress.FullSize == 0 {
		// unneeded as this is only called after the iterator is done
		barFilled = barWidth / 2
	} else {
		barFilled = int(math.Round(float64(barWidth) * float64(progress.DoneSize) / float64(progress.FullSize)))
	}

	fmt.Print("  [")

	for i := 0; i < barFilled-1; i++ {
		fmt.Print("=")
	}

	if barFilled == barWidth {
		fmt.Print("=")
	} else {
		fmt.Print(">")
	}

	for i := barFilled; i < barWidth; i++ {
		fmt.Print(" ")
	}

	fmt.Print("] ")
}

func printOperation() {
	fmt.Print("   ")

	switch flags.Current.Mode() {
	case 1:
		fmt.Print("Copying " + ui.ShrinkPath(progress.CurrentFile, maxWidth/2))
	case 2:
		fmt.Print("Linking " + ui.ShrinkPath(progress.CurrentFile, maxWidth/2))
	case 3:
		fmt.Print("Creating " + ui.ShrinkPath(progress.CurrentFile, maxWidth/2))
	}

	fmt.Print(" @ ")
	fmt.Print(ui.FormatSize(float64(progress.SizePerSecond),
		ui.SizeAutoUnit(float64(progress.SizePerSecond))))
	fmt.Print("/s")

	// remaining time:
	secondsLeft := float32(progress.FullSize-progress.DoneSize) / progress.SizePerSecond

	fmt.Print(", ")
	fmt.Print(ui.FormatSeconds(float64(secondsLeft)))
	fmt.Print(" remaining")
}

/*
func printConflict() {
	if len(piledConflicts) > 0 {
		filesLock.RLock()
		conflictID := piledConflicts[0]
		conflict := fileOrder[conflictID]
		cTarget := filepath.Join(targets[conflict],
			filepath.Base(conflict))
		filesLock.RUnlock()

		fmt.Print(color.FGColors.Yellow, color.Text.Bold)
		fmt.Print(conflict)
		fmt.Print(color.Text.Reset, color.FGColors.Magenta)
		fmt.Print(" already exists in ")
		fmt.Print(color.FGColors.Yellow, color.Text.Bold)
		fmt.Print(filepath.Dir(cTarget))
		fmt.Println(color.Text.Reset + color.FGColors.Magenta)
		lines++

		fmt.Println("[S]kip | Skip [A]ll | [O]verwrite | O[v]erwrite All")
		lines++

		fmt.Print("[I]nfo |      [R]ename target     | [Q]uit")
		fmt.Println(color.Text.Reset)
		lines++

		in := input.GetChoice("saovirq")

		switch in {
		case 's':
			filesLock.Lock()
			piledConflicts = piledConflicts[1:]
			filesLock.Unlock()
			skipFile(cTarget)
		case 'o':
			filesLock.Lock()
			pendingConflicts = append(pendingConflicts, conflictID)
			piledConflicts = piledConflicts[1:]
			filesLock.Unlock()
		case 'a':
			onExistingFile = 0 // skip all

			filesLock.Lock()
			piledConflicts = nil
			filesLock.Unlock()
		case 'v':
			onExistingFile = 1 // overwrite all

			filesLock.Lock()
			pendingConflicts = append(pendingConflicts, piledConflicts...)
			piledConflicts = nil
			filesLock.Unlock()
		case 'i':
			panic("not supported")
		case 'r':
			panic("not supported")
		case 'q':
			os.Exit(0)
		}
	}
}

/**
 * Add the file size to done_size and 1 to done_amount.
*/ /*
func skipFile(path string) {
	stat, err := os.Lstat(path)

	if err != nil {
		doneSize += 0
	} else if stat.Mode().IsRegular() {
		doneSize += uint64(stat.Size())
	} else if stat.Mode()&os.ModeSymlink != 0 {
		doneSize += uint64(symlinkSize)
	}

	doneAmount++
}*/

func printSummary() {
	if flags.Current.Verbosity() <= flags.VerbCrit {
		return
	}

	elapsed := time.Since(progress.Start)

	for i := 0; i < lines; i++ {
		fmt.Print("\033[1A\033[2K")
	}

	fmt.Print("  [")

	for i := 0; i < barWidth; i++ {
		fmt.Print("=")
	}

	fmt.Println("]")
	fmt.Print("   ")

	switch flags.Current.Mode() {
	case flags.ModeCopy:
		fmt.Print("Copied ")
	case flags.ModeMove:
		fmt.Print("Moved ")
	case flags.ModeRemove:
		fmt.Print("Deleted ")
	}

	if progress.FullAmount == 1 {
		fmt.Print("1 file in ")
	} else {
		fmt.Print(strconv.FormatUint(progress.FullAmount, 9))
		fmt.Print(" files in ")
	}

	fullSpeed := float64(progress.FullSize) / elapsed.Seconds()

	fmt.Print(ui.FormatSeconds(elapsed.Seconds()))
	fmt.Print(" (")
	fmt.Print(ui.FormatSize(fullSpeed, ui.SizeAutoUnit(fullSpeed)))
	fmt.Print("/s).")
	fmt.Println()
}
