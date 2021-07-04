package main

import (
	"easy-copy/lscolors"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

var buf []byte = make([]byte, 32678)

// ErrWritingData is returned by the copy loop
// if not all data could be written,
// but no other is returned.
type ErrWritingData struct {
	read    int
	written int
}

func (e *ErrWritingData) Error() string {
	return fmt.Sprintf("could only write %d b out of %d b", e.written, e.read)
}

func setBuffersize(size int) {
	buf = make([]byte, size)
	verbSetBuffersize(size)
}

// copyLoop checks for any work and does it synchronously.
func copyLoop() {
	var i int = 0

	for !done {
		filesLock.Lock()
		if len(folders) > 0 {
			// Create all folders:
			localFolders := folders
			folders = nil

			filesLock.Unlock()
			createFolders(localFolders)
		} else if len(pendingConflicts) > 0 {
			// Work resolved conflicts:
			id := pendingConflicts[0]
			pendingConflicts = pendingConflicts[1:]

			sourcePath := fileOrder[id]
			var destPath string
			if createFoldersInTarget {
				destPath = filepath.Join(targets[sourcePath],
					filepath.Base(sourcePath))
			} else {
				destPath = targets[sourcePath]
			}

			filesLock.Unlock()

			if !dryrun {
				os.Remove(destPath)
			}
			copyFilePath(sourcePath, destPath)
			doneAmount += 1
		} else if i < len(fileOrder) {
			// Copy normal files:
			var sourcePath string = fileOrder[i]
			var destPath string

			if createFoldersInTarget {
				destPath = filepath.Join(targets[sourcePath],
					filepath.Base(sourcePath))
			} else {
				destPath = targets[sourcePath]
			}

			filesLock.Unlock()

			// check if file already exists and we even care about that:
			var doCopy bool = true
			stat, err := os.Lstat(destPath)
			if err == nil && stat != nil {
				// file exists
				switch onExistingFile {
				case Skip:
					doCopy = false
				case Overwrite:
					if !dryrun {
						os.Remove(destPath)
					}
				case Ask:
					// save it to the conflicts:
					filesLock.Lock()
					piledConflicts = append(piledConflicts, i)
					filesLock.Unlock()
					doCopy = false
				default:
					// better safe than sorry
					doCopy = false
				}
			}

			if doCopy {
				copyFilePath(sourcePath, destPath)
				doneAmount += 1
			}

			i += 1
		} else {
			filesLock.Unlock()
		}

		if iteratorDone {
			// all sources have been iterated, no more work will come later on
			filesLock.RLock()
			if len(folders) == 0 && len(fileOrder) == i &&
				len(piledConflicts) == 0 && len(pendingConflicts) == 0 {
				// 1. all folders have been created
				// 2. we've tried to copy all files so far
				// 3. all conflicts we had to ask the user are resolved
				// 4. all conflicts the user already answered have been dealt with
				done = true

				if mode == ModeMove {
					syncdel(&fileOrder)
					syncdel(&sources)
				}
			}
			filesLock.RUnlock()
		}
	}
}

// copyFilePath clones the file (!) at sourcePath to
// destPath while adding the progress to done_size.
// If source is a symlink that links to a file,
// dest will be created as a link that links to the
// original file.
func copyFilePath(sourcePath string, destPath string) {
	stat, err := os.Lstat(sourcePath)
	if err != nil {
		errCopying(sourcePath, destPath, err)
	} else if stat.Mode().IsRegular() {
		if mode == ModeMove && !dryrun {
			// first attempt native move
			if isSameDevice(sourcePath, targetBase) {
				err := os.Rename(sourcePath, destPath)
				if err == nil { // yes, this should be == nil
					return
				}

				verbNativeMoveFailed(sourcePath, destPath, err)
			}
		}

		currentTaskType = 1
		currentFile = sourcePath
		var source, dest *os.File
		source, err = os.OpenFile(sourcePath, os.O_RDONLY, 0)
		if err != nil {
			errMissingFile(err, sourcePath)
		}
		if progressLSColors {
			currentFile = "\033[" + lscolors.FormatFile(stat) + "m" +
				currentFile + "\033[" + lscolors.FormatType("re") + "m"
		}

		if !dryrun {
			dest, err = os.OpenFile(destPath, os.O_CREATE|os.O_RDWR, stat.Mode().Perm())
			if err != nil {
				errCreatingFile(err, destPath)
			}
		}

		copyFile(source, dest, &doneSize)
		source.Close()
		dest.Close()
	} else if stat.Mode()&os.ModeSymlink != 0 {
		currentTaskType = 2
		currentFile = sourcePath
		if progressLSColors {
			currentFile = "\033[" + lscolors.FormatType("li") + "m" +
				currentFile + "\033[" + lscolors.FormatType("re") + "m"
		}

		var resolvedSourcePath string
		resolvedSourcePath, err = os.Readlink(sourcePath)
		if err != nil {
			errMissingFile(err, sourcePath)
		}
		sourcePath = resolvedSourcePath

		if !dryrun {
			if onExistingFile == 1 {
				os.Remove(destPath)
			}
			err = os.Symlink(sourcePath, destPath)

			if err != nil {
				errCreatingLink(err, sourcePath, destPath)
			}
		}

		doneSize += uint64(symlinkSize)
	}
}

func createFolders(folders []string) {
	var folder string

	for _, folder = range folders {
		currentTaskType = 3
		currentFile = folder

		if progressLSColors {
			currentFile = "\033[" + lscolors.FormatType("di") + "m" + currentFile +
				"\033[" + lscolors.FormatType("rs") + "m"
		}

		if !dryrun {
			var err error = os.MkdirAll(folder, 0o755)
			if err != nil {
				errCreatingFile(err, folder)
			}
		}

		doneSize += uint64(folderSize)
	}
}

// copyFile copies the openend source file to the already
// created dest file. Any error is handled over to
// errCopying().
func copyFile(source *os.File, dest *os.File, progressStorage *uint64) {
	var readAmount, writtenAmount int
	var err error

	for {
		readAmount, err = source.Read(buf)

		if err != nil && !errors.Is(err, io.EOF) {
			errCopying(source.Name(), dest.Name(), err)
		}

		if readAmount == 0 {
			// when the file is fully read
			break
		}

		if !dryrun {
			writtenAmount, err = dest.Write(buf[:readAmount])
			if err != nil {
				errCopying(source.Name(), dest.Name(), err)
			}

			if readAmount != writtenAmount {
				errCopying(source.Name(), dest.Name(),
					&ErrWritingData{read: readAmount, written: writtenAmount})
			}
		}

		*progressStorage += uint64(writtenAmount)
	}
}
