package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func iteratePaths() {
	filesLock.RLock()
	uPlen := len(unsearchedPaths)

	for uPlen > 0 {
		next := unsearchedPaths[0]
		unsearchedPaths = unsearchedPaths[1:] // discard that element
		filesLock.RUnlock()

		stat, err := os.Lstat(next)
		// TODO don't exit on missing file, coreutils cp doesnt do that
		if err != nil {
			errMissingFile(err, next)
		}

		switch {
		case stat.IsDir():
			wormholeCheck(next)

			dir, err := os.Open(next)
			if err != nil {
				errMissingFile(err, next)
			}

			// TODO dont read all files at once, specify an amount of files to read
			//  and recall Readdirnames until io.EOF is returned (I guess)
			names, err := dir.Readdirnames(0)
			if err != nil {
				errMissingFile(err, next)
			}

			filesLock.Lock()
			if createFoldersInTarget {
				folders = append(folders,
					filepath.Join(uPTargets[next], filepath.Base(next)))
			}

			for _, fileInFolder := range names {
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
			//    createFoldersInTarget is subsequently set to true
			//  - multiple files and folders are passed and should be copied
			//    into a target, createFoldersInTarget was already
			//    set to true in main
			createFoldersInTarget = true
			fullSize += uint64(folderSize)
			filesLock.Unlock()
		case stat.Mode().IsRegular():
			filesLock.Lock()
			fileOrder = append(fileOrder, next)
			targets[next] = uPTargets[next]
			filesLock.Unlock()
			fullSize += uint64(stat.Size())
		case stat.Mode()&os.ModeDevice != 0:
			warnBadFile(next)
		case stat.Mode()&os.ModeSymlink != 0:
			filesLock.Lock()
			var nextTarget string = uPTargets[next]
			if followSymlinks == 1 {
				fileOrder = append(fileOrder, next)
				targets[next] = nextTarget
				fullSize += uint64(symlinkSize)
			} else if followSymlinks == 2 {
				nextResolved, err := os.Readlink(next)
				if err != nil {
					errReadingSymlink(err, next)
				}
				unsearchedPaths = append(unsearchedPaths, nextResolved)
				uPTargets[nextResolved] = nextTarget
			}
			filesLock.Unlock()
		default:
			warnBadFile(next)
		}

		filesLock.RLock()
		uPlen = len(unsearchedPaths)
	}

	// close the l
	filesLock.RUnlock()

	iteratorDone = true
	fullAmount = uint64(len(fileOrder))

	verbDoneIterating()
	verbTargets()

	// as this function is forked anyways we can directly call this:
	drawLoop()
}

// wormholeCheck detectes attempts to copy folders into themselves:
func wormholeCheck(src string) {
	return
	fmt.Println("\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n")
	fmt.Println("is src:", src, "targetBase:", targetBase, "a wormhole?")
	fmt.Println("\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n")
}
