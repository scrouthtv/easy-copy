package main

import (
	"easy-copy/lscolors"
	"errors"
	"io"
	"os"
	"path/filepath"
	"strconv"
)

var buffersize uint = 32768

var buf []byte = make([]byte, buffersize)

/**
 * Loops checking / waiting for any left work.
 */
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
			var id int = pendingConflicts[0]
			pendingConflicts = pendingConflicts[1:]
			var sourcePath string = fileOrder[id]
			var destPath string
			if createFoldersInTarget {
				destPath = filepath.Join(targets[sourcePath],
					filepath.Base(sourcePath))
			} else {
				destPath = targets[sourcePath]
			}

			filesLock.Unlock()

			os.Remove(destPath)
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
					os.Remove(destPath)
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
				if mode == ModeMove {
					syncdel(&fileOrder)
					syncdel(&sources)
				}

				done = true
			}
			filesLock.RUnlock()
		}
	}
}

/**
 * This function copies source to dest, while
 *  adding the progress to done_size
 * If source is a symlink that links to a file,
 *  dest will be created as a link that links to that file as well.
 */
func copyFilePath(sourcePath string, destPath string) {
	var err error
	if doReflinks > 0 {
		err = reflink(sourcePath, destPath, &doneSize)
		if err == nil {
			return
		} else {
			if doReflinks == 1 {
				verbReflinkFailed(sourcePath, destPath, err)
			} else {
				errReflinkFailed(sourcePath, destPath, err)
				// os.Exit(2);
			}
		}
	}

	stat, err := os.Lstat(sourcePath)
	if err != nil {
		errCopying(sourcePath, destPath, err)
	} else if stat.Mode().IsRegular() {
		currentTaskType = 1
		currentFile = sourcePath
		var source, dest *os.File
		source, err = os.OpenFile(sourcePath, os.O_RDONLY, 0o644)
		if err != nil {
			errMissingFile(err, sourcePath)
		}
		if progressLSColors {
			currentFile = "\033[" + lscolors.FormatFile(stat) + "m" +
				currentFile + "\033[" + lscolors.FormatType("re") + "m"
		}

		dest, err = os.OpenFile(destPath, os.O_CREATE|os.O_RDWR, stat.Mode().Perm())
		if err != nil {
			errCreatingFile(err, destPath)
		}
		err := copyFile(source, dest, &doneSize)
		if err != nil {
			errCopying(source.Name(), dest.Name(), err)
		}
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
		if onExistingFile == 1 {
			os.Remove(destPath)
		}
		err = os.Symlink(sourcePath, destPath)

		if err != nil {
			errCreatingLink(err, sourcePath, destPath)
		}
		doneSize += uint64(symlinkSize)
	}
}

/**
 * Create the folders specified in folders.
 * filesLock will not be locked.
 */
func createFolders(folders []string) {
	var folder string

	for _, folder = range folders {
		currentTaskType = 3
		currentFile = folder

		if progressLSColors {
			currentFile = "\033[" + lscolors.FormatType("di") + "m" + currentFile +
				"\033[" + lscolors.FormatType("rs") + "m"
		}

		var err error = os.MkdirAll(folder, 0o755)
		if err != nil {
			errCreatingFile(err, folder)
		}

		doneSize += uint64(folderSize)
	}
}

func copyFile(source *os.File, dest *os.File, progressStorage *uint64) error {
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

		writtenAmount, err = dest.Write(buf[:readAmount])
		if err != nil {
			return err
		}

		if readAmount != writtenAmount {
			return errors.New("couldn't write all the data: " +
				strconv.Itoa(readAmount) + " read, " +
				strconv.Itoa(writtenAmount) + "written")
		}

		*progressStorage += uint64(writtenAmount)
	}

	return nil
}
