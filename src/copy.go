package main;

import "io";
import "os";
import "errors";
import "strconv";
import "time";
import "path/filepath";

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
