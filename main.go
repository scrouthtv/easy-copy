package main

import "os"
import "sync"
import "strings"
import "path/filepath"
import "time"

var createFoldersInTarget bool

var unsearchedPaths []string
var uPTargets map[string]string = make(map[string]string)
var targetBase string

var fileOrder []string
var folders []string
var targets map[string]string = make(map[string]string)

// read/write exclusion lock for the three arrays above
var filesLock = sync.RWMutex{}

var iteratorDone, done bool = false, false

var sources []string

var done_amount uint64 = 0
var full_amount uint64 = 0
var done_size uint64 = 0
var full_size uint64 = 0

var mode int = -1

const (
	MODE_CP = iota
	MODE_MV
	MODE_RM
)

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
	initColors(autoColors())
	parseArgs()
	readConfig()

	if verbose >= VERB_INFO {
		printVersion()
		verbFlags()
	}

	if len(unsearchedPaths) < 3 {
		errEmptySource()
	}

	switch strings.ToLower(unsearchedPaths[0]) {
	case "cp":
		mode = MODE_CP
	case "mv":
		mode = MODE_MV
	case "rm":
		mode = MODE_RM
		panic("This mode is not implemented (yet).")
	default:
		errInvalidMode(strings.ToLower(unsearchedPaths[0]), "cp, mv")
	}
	unsearchedPaths = unsearchedPaths[1:]

	var err error
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
		//folders = append(folders, targetBase);
		uPTargets[unsearchedPaths[0]] = targetBase
		createFoldersInTarget = false
	} else {
		createFoldersInTarget = true
		var uP string
		for _, uP = range unsearchedPaths {
			uPTargets[uP] = targetBase
		}
	}

	verbSearchStart()

	start = time.Now()

	go iteratePaths()
	copyLoop()

	printSummary()
}
