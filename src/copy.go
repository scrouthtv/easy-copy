package main;

import "io";
import "os";
import "errors";
import "strconv";
import "time";
import "path/filepath";
import "fmt";

const BUFFERSIZE uint = 1024;
var buf []byte = make([]byte, BUFFERSIZE);

func createFolders() {
	filesLock.RLock();
	if (len(folders) > 0) {
		verbCreatingFolders();
		var localFolders []string = append([]string(nil), folders...);
		folders = nil;
		filesLock.RUnlock();
		var folder string;
		for _, folder = range localFolders {
			var err error = os.MkdirAll(folder, 0755);
			if err != nil { errCreatingFile(err, folder); }
		}
	}
}

/**
 * At the beginning, a counter is set up to point at the next
 * file to copy ("i").
 * In a loop, the following will be done:
 *  1. If the iterator isn't done yet, retrieve
 *     fileOrder[i] as next or wait for a new file in fileOrder
 *  2. Any pending folders are created.
 *  3. It is tested whether the file already exists.
 *     If it does, and the specified overwrite strategy does not allow
 *     overwriting, it is skipped and the loop reruns with i+1.
 *  4. The destination file is created and the file is transferred whilst
 *     saving the progress to done_size.
 *  5. When done copying, i and done_amount is incremented by one.
 *  6. It is evaluated whether we are done:
 *     If the iterator is not finished, rerun the loop.
 *     If the iterator is finished and i == len(files), e. g. all files are
 *     copied, exit.
 */
func copyFiles() {
	var i int = 0;
	for !done {
		filesLock.RLock();
		for len(fileOrder) <= i {
			filesLock.RUnlock();
			if iteratorDone {
				verbDoneIterating();
				return;
			}
			time.Sleep(100 * time.Millisecond);
			filesLock.RLock();
		}
		var sourcePath string = fileOrder[i];
		var destPath string = filepath.Join(targets[sourcePath],
			filepath.Base(sourcePath));
		filesLock.RUnlock();

		verbCopyStart(sourcePath, destPath);
		var source, dest *os.File;
		var err error;
		source, err = os.OpenFile(sourcePath, os.O_RDONLY, 0755);
		if err != nil { errMissingFile(err, sourcePath); }
		createFolders();

		// check if file exists:
		if onExistingFile != 1 {
			stat, _ := os.Lstat(destPath);
			if stat != nil {
				// file exists
				if onExistingFile == 2 {
					// ask
					fmt.Println("ask");
					piledOverwrites = append(piledOverwrites, i);
				}
				i += 1;
				fmt.Println("next file");
				continue; // rerun loop
			} else {
				// TODO
			}
		}
		dest, err = os.OpenFile(destPath, os.O_WRONLY | os.O_CREATE, 0755);
		if err != nil { errCreatingFile(err, destPath); }
		copyFile(source, dest, &done_size);
		i += 1;
		done_amount += 1;

		// check if we are done:
		if iteratorDone {
			filesLock.RLock();
			// all folders are created & we copied all files up to this point
			if len(folders) == 0 && len(fileOrder) == i {
				done = true;
			}
			filesLock.RUnlock();
		}
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
