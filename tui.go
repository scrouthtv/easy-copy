package main

import (
	"easy-copy/color"
	"easy-copy/input"
	"fmt"
	"math"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

const (
	barWidth int = 50
	maxWidth int = 80
)

// contains ids to files that should be recopied after the
//  dialog whether to overwrite files has been answered.
// once their respective dialogs have been answered, they are either
//  added to pendingOverwrites or simply removed from piledOverwrites.
var (
	piledConflicts   []int
	pendingConflicts []int
)

// 0 undefined
// 1 Copying
// 2 Linking
// 3 Creating Folder
// 4 Deleting
var (
	currentTaskType int = -1
	currentFile     string
)

var lines int = 0

func drawLoop() {
	go speedLoop()
	for !done {

		var i int
		for i = 0; i < lines; i++ {
			fmt.Print("\033[1A\033[2K")
		}
		lines = 0

		if verbose > VerbQuiet {
			var barFilled int

			if fullSize == 0 {
				// unneeded as this is only called after the iterator is done
				barFilled = barWidth / 2
			} else {
				barFilled = int(math.Round(float64(barWidth) * float64(doneSize) / float64(fullSize)))
			}

			fmt.Print("  [")
			var i int
			for i = 0; i < barFilled-1; i++ {
				fmt.Print("=")
			}
			if barFilled == barWidth {
				fmt.Print("=")
			} else {
				fmt.Print(">")
			}
			for i = barFilled; i < barWidth; i++ {
				fmt.Print(" ")
			}
			fmt.Print("] ")
			unit := sizeAutoUnit(float64(fullSize))
			fmt.Print(formatSize(float64(doneSize), unit))
			fmt.Print(" / ")
			fmt.Print(formatSize(float64(fullSize), unit))
			fmt.Println()
			lines++

			// speed:
			fmt.Print("   ")
			switch currentTaskType {
			case 1:
				fmt.Print("Copying " + shrinkPath(currentFile, maxWidth/2))
			case 2:
				fmt.Print("Linking " + shrinkPath(currentFile, maxWidth/2))
			case 3:
				fmt.Print("Creating " + shrinkPath(currentFile, maxWidth/2))
			}
			fmt.Print(" @ ")
			fmt.Print(formatSize(float64(sizePerSecond),
				sizeAutoUnit(float64(sizePerSecond))))
			fmt.Print("/s")

			// remaining time:
			fmt.Print(", ")
			secondsLeft := float32(fullSize-doneSize) / sizePerSecond
			fmt.Print(formatSeconds(float64(secondsLeft)))
			fmt.Println(" remaining")
			lines++
		}

		if len(piledConflicts) > 0 {
			filesLock.RLock()
			var conflictID int = piledConflicts[0]
			var conflict string = fileOrder[conflictID]
			var cTarget string = filepath.Join(targets[conflict],
				filepath.Base(conflict))
			filesLock.RUnlock()
			fmt.Println()
			lines++
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
			in := input.GetChoice("soavidreq")

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

		time.Sleep(100 * time.Millisecond)
	}
}

/**
 * Add the file size to done_size and 1 to done_amount.
 */
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
}

func printSummary() {
	elapsed := time.Since(start)
	for i := 0; i < lines; i++ {
		fmt.Print("\033[1A\033[2K")
	}

	if verbose > VerbQuiet {
		fmt.Print("  [")

		for i := 0; i < barWidth; i++ {
			fmt.Print("=")
		}

		fmt.Println("]")
		fmt.Print("   ")
		switch mode {
		case ModeCopy:
			fmt.Print("Copied ")
		case ModeMove:
			fmt.Print("Moved ")
		case ModeRemove:
			fmt.Print("Deleted ")
		}
		if fullAmount == 1 {
			fmt.Print("1 file in ")
		} else {
			fmt.Print(strconv.FormatUint(fullAmount, 9))
			fmt.Print(" files in ")
		}
		fmt.Print(formatSeconds(elapsed.Seconds()))
		fmt.Print(" (")
		fullSpeed := float64(fullSize) / float64(elapsed.Seconds())
		fmt.Print(formatSize(float64(fullSpeed),
			sizeAutoUnit(float64(fullSpeed))))
		fmt.Print("/s).")
		fmt.Println()
	}
}
