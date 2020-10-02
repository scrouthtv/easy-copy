package main

import "fmt"
import "os"

var unsearchedPaths []string;
var pwd string;
var target string;

var targets map[string]string = make(map[string]string);
var fileOrder []string;
var filesLock = sync.RWMutex{}; // read/write exclusion lock

// all copying is done:
var done bool = false;

var done_amount uint64 = 0;
var full_amount uint64 = 0;
var done_size uint64 = 0;
var full_size uint64 = 0;
// Maybe these are too small:
// uint64 goes up to 18446744073709551615
// or 2097152 TB

// TODO: check if enough space is free
// settable buffer width
// dry run
// config file
// follow symlinks?
// backup
// sync

func iteratePaths() {
	for len(unsearchedPaths) > 0 {
		var next string = unsearchedPaths[0];
		unsearchedPaths = unsearchedPaths[1:]; // discard first element

		fmt.Println("search", next + ":");
		stat, err := os.Stat(next);
		if err != nil {
			fmt.Println(err);
			return;
		}
		if (os.IsNotExist(err)) {
			if verbose {
				fmt.Println(fmtFile(next), "does not exist");
			}
			// handle non existent file
		} else if (stat.IsDir()) {
			if verbose { fmt.Println(fmtFile(next), "is a dir"); }
			// handle directory
		} else if (stat.Mode().IsRegular()) {
			if verbose { fmt.Println(fmtFile(next), "is regular"); }
		} else if (stat.Mode() & os.ModeDevice != 0) {
			if verbose { fmt.Println(fmtFile(next), "is a device file"); }
		} else if (stat.Mode() & os.ModeSymlink != 0) {
			if verbose { fmt.Println(fmtFile(next), "is a symlink"); }
		} else {
			if verbose { fmt.Println(fmtFile(next), "is weird"); }
		}
	}
}

// copy works as follows:
// 1. open source for reading
// 2. stat target,
// 		if it is file, open it for writing (check if it exists & we want to overwrite)
//		if it is directory, create a new file in it with the same name as source
// 3. copy it over
// 4. eventually delete the source file

func main() {
	parseArgs();
	pwd, err := os.Getwd();
	if (err != nil) {
		errInvalidWD(err);
	}

	if verbose {
		printVersion();
		verboseFlags();
		fmt.Println("Working directory", pwd)
		fmt.Println("Have to search", unsearchedPaths);
	}

	if len(unsearchedPaths) < 2 {
		errEmptySource();
	}

	target = unsearchedPaths[len(unsearchedPaths) - 1];
	unsearchedPaths = unsearchedPaths[0:len(unsearchedPaths) - 1];

	if verbose {
		fmt.Println("Have to search", unsearchedPaths);
		fmt.Println("Target is", target);
	}

	iteratePaths();
	//copyFiles();
}
