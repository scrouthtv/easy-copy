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
	done_amount uint64 = 0
	full_amount uint64 = 0
	done_size   uint64 = 0
	full_size   uint64 = 0
)

var mode int = -1

const (
	ModeCopy = iota
	ModeMove
	ModeRemove
)

var doReadConfig bool = true

// Maybe these are too small:
// uint64 goes up to 18446744073709551615
// or 2097152 TB

var start time.Time

func iteratePaths() {
	filesLock.RLock()
	var uPlen int = len(unsearchedPaths)
	for uPlen > 0 {
		var next string = unsearchedPaths[0]
		unsearchedPaths = unsearchedPaths[1:] // discard first element
		filesLock.RUnlock()

		var err error
		var stat os.FileInfo
		stat, err = os.Lstat(next)
		// TODO don't exit on missing file, coreutils cp doesnt do that
		if err != nil {
			errMissingFile(err, next)
		}
		if stat.IsDir() {
			dir, err := os.Open(next)
			if err != nil {
				errMissingFile(err, next)
			}
			var names []string
			// TODO dont read all files at once, specify an amount of files to read
			//  and recall Readdirnames until io.EOF is returned (I guess)
			names, err = dir.Readdirnames(0)
			if err != nil {
				errMissingFile(err, next)
			}

			filesLock.Lock()
			if createFoldersInTarget {
				folders = append(folders,
					filepath.Join(uPTargets[next], filepath.Base(next)))
			}
			var fileInFolder string
			for _, fileInFolder = range names {
				unsearchedPaths = append(unsearchedPaths, filepath.Join(next, fileInFolder))
				if createFoldersInTarget {
					uPTargets[filepath.Join(next, fileInFolder)] = filepath.Join(uPTargets[next], filepath.Base(next))
				} else {
					uPTargets[filepath.Join(next, fileInFolder)] = uPTargets[next]
				}
			}
			// three possibilities:
			//  - only one *file* is passed and should be duplicated.
			//    we don't care about creating folders in the target
			//    as there aren't any to create anyways
			//  - only one *folder* is passed and should be duplicated.
			//    this if statement is the very first to be called and
			//    this variable is subsequently set to true
			//  - multiple files and folders are passed and should be copied
			//    into a target, this variable was already set to true in main
			createFoldersInTarget = true
			full_size += uint64(folder_size)
			filesLock.Unlock()
		} else if stat.Mode().IsRegular() {
			filesLock.Lock()
			fileOrder = append(fileOrder, next)
			targets[next] = uPTargets[next]
			filesLock.Unlock()
			full_size += uint64(stat.Size())
		} else if stat.Mode()&os.ModeDevice != 0 {
			warnBadFile(next)
		} else if stat.Mode()&os.ModeSymlink != 0 {
			filesLock.Lock()
			var nextTarget string = uPTargets[next]
			if followSymlinks == 1 {
				fileOrder = append(fileOrder, next)
				targets[next] = nextTarget
				full_size += uint64(symlink_size)
			} else if followSymlinks == 2 {
				nextResolved, err := os.Readlink(next)
				if err != nil {
					errReadingSymlink(err, next)
				}
				unsearchedPaths = append(unsearchedPaths, nextResolved)
				uPTargets[nextResolved] = nextTarget
			}
			filesLock.Unlock()
		} else {
			warnBadFile(next)
		}
		filesLock.RLock()
		uPlen = len(unsearchedPaths)
	}
	filesLock.RUnlock()
	iteratorDone = true
	full_amount = uint64(len(fileOrder))
	verbTargets()
	// as this function is forked anyways we can directly call this:
	drawLoop()
}

func main() {
	var err error

	color.Init(color.AutoColors())
	parseArgs()
	if doReadConfig {
		var kvs []string
		kvs, err = config.Load()
		if err == nil {
			var line string
			for _, line = range kvs {
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
	createFolders([]string{targetBase})

	verbSearchStart()

	start = time.Now()

	go iteratePaths()
	copyLoop()

	printSummary()
}
