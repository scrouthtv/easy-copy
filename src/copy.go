package main;

import "io";
import "fmt";
import "os";
import "errors";
import "strconv";
import "time";

const BUFFERSIZE uint = 1024;
var buf []byte = make([]byte, BUFFERSIZE);

func copyFiles() {
	var i int = 0;
	for !done {
		filesLock.RLock();
		if (len(folders) > 0) {
			if verbose { fmt.Println("creating folders"); }
			var localFolders []string = append([]string(nil), folders...);
			folders = nil;
			filesLock.RUnlock();
			var folder string;
			for _, folder = range localFolders {
				var folderInTarget string = rebasePathOntoTarget(folder)
				if verbose { fmt.Println("creating", folderInTarget + ":"); }
				var err error = os.Mkdir(folderInTarget, 0755);
				if err != nil { errCreatingFile(err, folderInTarget); }
			}
		} else {
			filesLock.RUnlock();
		}
		filesLock.RLock();
		for len(fileOrder) <= i {
			filesLock.RUnlock();
			if iteratorDone {
				if verbose { fmt.Println("done iterating, no more files"); }
				return;
			}
			time.Sleep(100 * time.Millisecond);
			filesLock.RLock();
		}
		var sourcePath string = fileOrder[i];
		var destPath string = targets[sourcePath];
		filesLock.RUnlock();

		if verbose { fmt.Println("src: ", sourcePath, "dest: ", destPath); }

		var source, dest *os.File;
		var err error;
		source, err = os.OpenFile(sourcePath, os.O_RDONLY, 0755);
		if err != nil { errMissingFile(err, sourcePath); }
		dest, err = os.OpenFile(destPath, os.O_WRONLY | os.O_CREATE, 0755);
		copyFile(source, dest, &done_size);
		i += 1;
		done_amount += 1;

		// check if we are done:
		if iteratorDone {
			filesLock.RLock();
			if len(folders) == 0 && len(fileOrder) == 0 {
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
		if err != nil && err != io.EOF {
			return err;
		}
		if readAmount == 0 {
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
	return nil;
}
