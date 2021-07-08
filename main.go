package main

import (
	"easy-copy/color"
	"os"
	"path/filepath"
	"sync"
	"time"
)

var createFoldersInTarget bool

var (
	unsearchedPaths []string
	uPTargets       map[string]string = make(map[string]string)
	targetBase      string
)

var (
	fileOrder []string
	folders   []string
	targets   map[string]string = make(map[string]string)
)

// filesLock is an exclusion lock for the three arrays above.
var filesLock = sync.RWMutex{}

var iteratorDone, done bool = false, false

var sources []string
var nodelete []string

var (
	doneAmount uint64 = 0
	fullAmount uint64 = 0
	doneSize   uint64 = 0
	fullSize   uint64 = 0
)

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

var start time.Time

func main() {
	var err error

	color.Init(color.AutoColors())

	readConfig()

	parseArgs()

	if verbose >= VerbInfo {
		printVersion()
		verbFlags()
	}

	if len(unsearchedPaths) < 2 {
		errEmptySource()
	}

	targetBase, err = filepath.Abs(unsearchedPaths[len(unsearchedPaths)-1])
	if err != nil {
		errResolvingTarget(unsearchedPaths[len(unsearchedPaths)-1], err)
	}

	unsearchedPaths = unsearchedPaths[0 : len(unsearchedPaths)-1]
	sources = unsearchedPaths

	if len(unsearchedPaths) == 1 {
		// if there is only one source, we want to duplicate it:

		uPTargets[unsearchedPaths[0]] = targetBase

		stat, err := os.Stat(targetBase)

		// if the target already exists as a folder, we want to copy into it:
		if err == nil && stat.IsDir() {
			createFoldersInTarget = true
		} else {
			createFoldersInTarget = false
		}

		// if the source is a folder, we have to create the duplicated folder:
		stat, err = os.Stat(unsearchedPaths[0])
		if err != nil {
			errMissingFile(err, unsearchedPaths[0])
		}

		if stat.IsDir() && !dryrun {
			err := os.MkdirAll(targetBase, 0o755)
			if err != nil {
				errCreatingFile(err, targetBase)
			}
		}
	} else {
		// if there is more than one source, we want to copy the files
		// into the target directory:
		stat, err := os.Stat(targetBase)
		if os.IsNotExist(err) && !dryrun {
			err = os.MkdirAll(targetBase, 0o755)
			if err != nil {
				errCreatingFile(err, targetBase)
			}
		} else if err != nil {
			errCreatingFile(err, targetBase)
		} else if !stat.IsDir() {
			errTargetNoDir(targetBase)
		}

		createFoldersInTarget = true
		for _, uP := range unsearchedPaths {
			uPTargets[uP] = targetBase
		}
	}

	if createFoldersInTarget {
		createFolders([]string{targetBase})
	}

	verbSearchStart()

	go setOptimalBuffersize()

	start = time.Now()

	go iteratePaths()
	go speedLoop()
	go watchdog()

	copyLoop()

	printSummary()
}
