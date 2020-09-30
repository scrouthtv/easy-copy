package main

import "fmt"
import "os"

var unsearchedPaths []string;
var files map[string]string = make(map[string]string);
var pwd string;
var target string;
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

func iteratePaths() {
	for len(unsearchedPaths) >= 0 {
		var next string = unsearchedPaths[0];
		unsearchedPaths = unsearchedPaths[1:]; // discard first element

		fmt.Println("search ", next, ":");
		file, err := os.Stat(next);
		fmt.Printf("%T\n", file);
		if err != nil {
			fmt.Println(err);
			return;
		}
		if (os.IsNotExist(err)) {
			fmt.Println(file, " does not exist");
		} else if (file.IsDir()) {
			fmt.Println(file, " is a dir");
		} else if(file.Mode().IsRegular()) {
			fmt.Println(file, " is regular");
		} else {
			fmt.Println(file, " is weird");
		}

		fmt.Println("end of loop");
	}
}

func main() {
	parseArgs();

	if verbose {
		printVersion();
		verboseFlags();
	}

	if len(unsearchedPaths) < 2 {
		errEmptySource();
	}

	target = unsearchedPaths[len(unsearchedPaths) - 1];
	unsearchedPaths = unsearchedPaths[0:len(unsearchedPaths) - 2];

	if verbose {
		fmt.Println("Have to search ", unsearchedPaths);
		fmt.Println("Target is ", target);
	}

	go iteratePaths();
	//copyFiles();
}
