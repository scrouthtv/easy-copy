package main

import (
	"easy-copy/color"
	"fmt"
	"math"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

const (
	BAR_WIDTH int = 50
	MAX_WIDTH int = 80
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
			var BAR_FILLED int

			if full_size == 0 {
				// unneeded as this is only called after the iterator is done
				BAR_FILLED = BAR_WIDTH / 2
			} else {
				BAR_FILLED = int(math.Round(float64(BAR_WIDTH) * float64(done_size) / float64(full_size)))
			}

			fmt.Print("  [")
			var i int
			for i = 0; i < BAR_FILLED-1; i++ {
				fmt.Print("=")
			}
			if BAR_FILLED == BAR_WIDTH {
				fmt.Print("=")
			} else {
				fmt.Print(">")
			}
			for i = BAR_FILLED; i < BAR_WIDTH; i++ {
				fmt.Print(" ")
			}
			fmt.Print("] ")
			var unit int
			unit = sizeAutoUnit(float64(full_size))
			fmt.Print(formatSize(float64(done_size), unit))
			fmt.Print(" / ")
			fmt.Print(formatSize(float64(full_size), unit))
			fmt.Println()
			lines++

			// speed:
			fmt.Print("   ")
			switch currentTaskType {
			case 1:
				fmt.Print("Copying " + shrinkPath(currentFile, MAX_WIDTH/2))
			case 2:
				fmt.Print("Linking " + shrinkPath(currentFile, MAX_WIDTH/2))
			case 3:
				fmt.Print("Creating " + shrinkPath(currentFile, MAX_WIDTH/2))
			}
			fmt.Print(" @ ")
			fmt.Print(formatSize(float64(sizePerSecond),
				sizeAutoUnit(float64(sizePerSecond))))
			fmt.Print("/s")

			// remaining time:
			fmt.Print(", ")
			var secondsLeft float32 = float32(full_size-done_size) / sizePerSecond
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
			var in rune = getChoice("soavidreq")
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
	var stat os.FileInfo
	stat, _ = os.Lstat(path)

	if stat.Mode().IsRegular() {
		done_size += uint64(stat.Size())
	} else if stat.Mode()&os.ModeSymlink != 0 {
		done_size += uint64(symlinkSize)
	}

	done_amount += 1
}

func printSummary() {
	var elapsed time.Duration
	elapsed = time.Now().Sub(start)
	var i int
	for i = 0; i < lines; i++ {
		fmt.Print("\033[1A\033[2K")
	}

	if verbose > VerbQuiet {
		fmt.Print("  [")
		for i = 0; i < BAR_WIDTH; i++ {
			fmt.Print("=")
		}
		fmt.Println("]")
		fmt.Print("   Copied " + strconv.FormatUint(full_amount, 9))
		fmt.Print(" files in ")
		fmt.Print(formatSeconds(elapsed.Seconds()))
		fmt.Print(" (")
		var full_speed float64
		full_speed = float64(full_size) / float64(elapsed.Seconds())
		fmt.Print(formatSize(float64(full_speed),
			sizeAutoUnit(float64(full_speed))))
		fmt.Print("/s).")
		fmt.Println()
	}
}
