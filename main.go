package main

import (
	"easy-copy/color"
	"easy-copy/config"
	"os"
	"path/filepath"
	"strings"
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

// read/write exclusion lock for the three arrays above
var filesLock = sync.RWMutex{}

var iteratorDone, done bool = false, false

var sources []string

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

var doReadConfig bool = true

// Maybe these are too small:
// uint64 goes up to 18446744073709551615
// or 2097152 TB

var start time.Time

func main() {
	var err error

	color.Init(color.AutoColors())
	parseArgs()
	if doReadConfig {
		var kvs []string
		kvs, err = config.Load()
		if err == nil {
			for _, line := range kvs {
				parseOption(line)
			}
		} else {
			warnConfig(err)
		}
	}

	if verbose >= VerbInfo {
		printVersion()
		verbFlags()
	}

	if len(unsearchedPaths) < 3 {
		errEmptySource()
	}

	switch strings.ToLower(unsearchedPaths[0]) {
	case "cp":
		mode = ModeCopy
	case "mv":
		mode = ModeMove
	case "rm":
		mode = ModeRemove
		panic("This mode is not implemented (yet).")
	default:
		errInvalidMode(strings.ToLower(unsearchedPaths[0]), "cp, mv")
	}
	unsearchedPaths = unsearchedPaths[1:]

	targetBase, err = filepath.Abs(unsearchedPaths[len(unsearchedPaths)-1])
	if err != nil {
		errResolvingTarget(unsearchedPaths[len(unsearchedPaths)-1], err)
	}
	unsearchedPaths = unsearchedPaths[0 : len(unsearchedPaths)-1]
	sources = unsearchedPaths
	if len(unsearchedPaths) > 1 {
		stat, err := os.Stat(targetBase)
		if err != nil {
			errMissingFile(err, targetBase)
		}
		if os.IsNotExist(err) {
			errMissingFile(err, targetBase)
		} else if !stat.IsDir() {
			errTargetNoDir(targetBase)
		}
	}
	if len(unsearchedPaths) == 1 {
		// folders = append(folders, targetBase);
		uPTargets[unsearchedPaths[0]] = targetBase
		createFoldersInTarget = false
	} else {
		createFoldersInTarget = true
		var uP string
		for _, uP = range unsearchedPaths {
			uPTargets[uP] = targetBase
		}
	}
	if createFoldersInTarget {
		createFolders([]string{targetBase})
	}

	verbSearchStart()

	start = time.Now()

	go iteratePaths()
	copyLoop()

	printSummary()
}
