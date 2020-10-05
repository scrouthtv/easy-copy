package main;

import "io";
import "os";
import "errors";
import "strconv";
import "path/filepath";

const BUFFERSIZE uint = 1024;

var buf []byte = make([]byte, BUFFERSIZE);

/**
 * Loops checking / waiting for any left work.
 */
func copyLoop() {
	var i int = 0;
	for !done {
		filesLock.Lock();
		if len(folders) > 0 {
			var localFolders []string = folders;
			folders = nil;
			filesLock.Unlock();
			createFolders(localFolders);
		} else if len(pendingConflicts) > 0 {
			// TODO
			filesLock.Unlock();
		} else if i < len(fileOrder) {
			var sourcePath string = fileOrder[i];
			var destPath string = filepath.Join(targets[sourcePath],
				filepath.Base(sourcePath));
			filesLock.Unlock();

			// check if file already exists and we even care about that:
			var doCopy bool = true;
			if onExistingFile != 1 {
				stat, _ := os.Lstat(destPath);
				// TODO error handling
				if stat != nil {
					doCopy = false;
					// file exists
					if onExistingFile == 2 {
						// save it to the conflicts:
						filesLock.Lock();
						piledConflicts = append(piledConflicts, i);
						drawAskOverwriteDialog = true;
						filesLock.Unlock();
					}
				}
			}
			if doCopy {
				copyFilePath(sourcePath, destPath);
			}
			i += 1;
			done_amount += 1;
		} else {
			filesLock.Unlock();
		}

		if iteratorDone {
			// all sources have been iterated, no more work will come later on
			filesLock.RLock();
			if len(folders) == 0 && len(fileOrder) == i &&
				len(piledConflicts) == 0 && len(pendingConflicts) == 0 {
					// 1. all folders have been created
					// 2. we've tried to copy all files so far
					// 3. all conflicts we had to ask the user are resolved
					// 4. all conflicts the user already answered have been dealt with
				done = true;
			}
			filesLock.RUnlock();
		}
	}
}

func copyFilePath(sourcePath string, destPath string) {
	verbCopyStart(sourcePath, destPath);
	var source, dest *os.File;
	var err error;
	source, err = os.OpenFile(sourcePath, os.O_RDONLY, 0644);
	if err != nil { errMissingFile(err, sourcePath); }
	dest, err = os.OpenFile(destPath, os.O_CREATE | os.O_APPEND | os.O_WRONLY, 0644);
	if err != nil { errCreatingFile(err, destPath); }
	copyFile(source, dest, &done_size);
}

/**
 * Create the folders specified in folders.
 * filesLock will not be locked.
 */
func createFolders(folders []string) {
	verbCreatingFolders();
	var folder string;
	for _, folder = range folders {
		var err error = os.MkdirAll(folder, 0755);
		if err != nil { errCreatingFile(err, folder); }
	}
}

/**
 * This function copies source to dest, while
 *  adding the progress in bytes to progressStorage
 * The outer function still has to open (and create) source and dest
 *  and handle returned errors.
 */
func copyFile(source *os.File, dest *os.File, progressStorage *uint64) error {
	var readAmount, writtenAmount int;
	var err error;
	for {
		readAmount, err = source.Read(buf);
		if err != nil && err != io.EOF {
			errCopying(source.Name(), dest.Name(), err);
		}
		if readAmount == 0 {
			// when the file is fully read
			break;
		}
		writtenAmount, err = dest.Write(buf[:readAmount]);
		if err != nil {
			return err;
		}
		if readAmount != writtenAmount {
			return errors.New("couldn't write all the data: " + strconv.Itoa(readAmount) +
				" read, " + strconv.Itoa(writtenAmount) + "written");
		}
		*progressStorage += uint64(writtenAmount);
	}
	verbCopyFinished(source.Name(), dest.Name());
	return nil;
}
