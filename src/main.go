package main;

import "os";
import "sync";
import "path/filepath";
import "fmt";

var unsearchedPaths []string;
var uPTargets map[string]string = make(map[string]string);
var targetBase string;

var fileOrder []string;
var folders []string;
var targets map[string]string = make(map[string]string);
var filesLock = sync.RWMutex{};
// read/write exclusion lock for the three arrays above

var iteratorDone, done bool = false, false;

var done_amount uint64 = 0;
var full_amount uint64 = 0;
var done_size uint64 = 0;
var full_size uint64 = 0;
// Maybe these are too small:
// uint64 goes up to 18446744073709551615
// or 2097152 TB

func iteratePaths() {
	filesLock.RLock();
	var uPlen int = len(unsearchedPaths);
	for uPlen > 0 {
		var next string = unsearchedPaths[0];
		unsearchedPaths = unsearchedPaths[1:]; // discard first element
		filesLock.RUnlock();

		var err error;
		var stat os.FileInfo;
		stat, err = os.Lstat(next);
		if err != nil { errMissingFile(err, next); }
		if (os.IsNotExist(err)) {
			errMissingFile(err, next);
			// TODO don't exit on missing file, coreutils cp doesnt do that
		} else if (stat.IsDir()) {
			dir, err := os.Open(next);
			if err != nil { errMissingFile(err, next); }
			var names []string;
			// TODO dont read all files at once, specify an amount of files to read
			//  and recall Readdirnames until io.EOF is returned (I guess)
			names, err = dir.Readdirnames(0);
			if err != nil { errMissingFile(err, next); }
			// merge target + folder name + file in folder name

			filesLock.Lock();
			folders = append(folders,
				filepath.Join(uPTargets[next], filepath.Base(next)));
			var fileInFolder string;
			for _, fileInFolder = range names {
				unsearchedPaths = append(unsearchedPaths, filepath.Join(next, fileInFolder));
				uPTargets[filepath.Join(next, fileInFolder)] = filepath.Join(uPTargets[next], filepath.Base(next));
			}
			full_size += uint64(folder_size);
			filesLock.Unlock();
		} else if (stat.Mode().IsRegular()) {
			filesLock.Lock();
			fileOrder = append(fileOrder, next);
			targets[next] = uPTargets[next];
			filesLock.Unlock();
			full_size += uint64(stat.Size());
		} else if (stat.Mode() & os.ModeDevice != 0) {
			warnBadFile(next);
		} else if (stat.Mode() & os.ModeSymlink != 0) {
			var nextTarget string = uPTargets[next];

			filesLock.Lock();
			if followSymlinks == 1 {
				fileOrder = append(fileOrder, next);
				targets[next] = nextTarget;
				full_size += uint64(symlink_size);
			} else if followSymlinks == 2 {
				// TODO what if the symlink points to a directory:
				//  don't add this to file order but to unsearched paths
				nextResolved, err := os.Readlink(next);
				if err != nil { errReadingSymlink(err, next); }
				fileOrder = append(fileOrder, nextResolved);
				targets[nextResolved] = nextTarget;
			}
			filesLock.Unlock();
		} else {
			warnBadFile(next);
		}
		filesLock.RLock();
		uPlen = len(unsearchedPaths);
	}
	filesLock.RUnlock();
	iteratorDone = true;
	full_amount = uint64(len(fileOrder));
	verbTargets();
	// as this function is forked anyways we can directly call this:
	drawLoop();
}

// copy works as follows:
// 1. open source for reading
// 2. stat target,
// 		if it is file, open it for writing (check if it exists & we want to overwrite)
//		if it is directory, create a new file in it with the same name as source
// 3. copy it over
// 4. eventually delete the source file

func main() {
	initColors(autoColors());
	parseArgs();
	readConfig();

	if verbose {
		printVersion();
		verbFlags();
	}

	if len(unsearchedPaths) < 2 {
		errEmptySource();
	}

	var err error;
	targetBase, err = filepath.Abs(unsearchedPaths[len(unsearchedPaths) - 1]);
	if err != nil {
		errResolvingTarget(unsearchedPaths[len(unsearchedPaths) - 1], err);
	}
	unsearchedPaths = unsearchedPaths[0:len(unsearchedPaths) - 1];
	if len(unsearchedPaths) > 1 {
		stat, err := os.Stat(targetBase);
		if err != nil {
			errMissingFile(err, targetBase);
		}
		if os.IsNotExist(err) {
			errMissingFile(err, targetBase);
		} else if !stat.IsDir() {
			errTargetNoDir(targetBase);
		}
	}

	verbSearchStart();

	go iteratePaths();
	copyFiles();
}
