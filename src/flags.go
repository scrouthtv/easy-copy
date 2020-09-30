package main

import "fmt"
import "os"

var verbose bool = false;

var onExistingFile uint8 = 2;
// 0 skip
// 1 overwrite
// 2 ask every time
// 3 

func verboseFlags() {
	fmt.Println("Verbose: ", verbose);
	fmt.Println("Overwrite Mode: ", onExistingFile);
}

func parseFlag(prefix string, flag string) {
	switch(flag) {
	case "h": case "help":
		printHelp();
		os.Exit(0);
	case "v": case "version":
		printVersion();
		os.Exit(0);
	case "V": case "verbose":
		verbose = true;
		break;
	case "f": case "force":
		onExistingFile = 1;
		break;
	case "i": case "interactive":
		onExistingFile = 2;
		break;
	case "n": case "no-clobber": //case "no-overwrite":
		onExistingFile = 0;
		break;
	default:
		errUnknownOption(prefix + flag);
	}
}
