package main;

import "fmt";
import "os";

var verbose bool;
var onExistingFile uint8 = 2;
// 0 skip
// 1 overwrite
// 2 ask
var followSymlinks uint8 = 1;
// 0 ignore symlinks
// 1 follow symlinks, copying them as links
// 2 fully dereference

func verboseFlags() {
	fmt.Printf(FGColors.Green)
	fmt.Println(" Verbose:", verbose);
	fmt.Println(" Overwrite Mode:", onExistingFile);
	fmt.Print(" Follow symlinks: ", followSymlinks);
	fmt.Println(Textstyle.Reset);
}

func parseOption(key string, value string) {
	switch(key) {
		case "verbose":
			verbose = configInterpretBoolean(value);
		case "overwrite":
			switch (value) {
				case "skip": onExistingFile = 0;
				case "overwrite": onExistingFile = 1;
				case "ask": onExistingFile = 2;
			}
		case "symlinks":
			switch (value) {
				case "ignore": followSymlinks = 0;
				case "link": followSymlinks = 1;
				case "dereference": followSymlinks = 2;
			}
		default:
			fmt.Println("Unknown config key", key);
	}
}

func parseFlag(prefix string, flag string) {
	switch(flag) {
	case "h", "help":
		printHelp();
		os.Exit(0);
	case "v", "version":
		printVersion();
		os.Exit(0);
	case "V", "verbose":
		verbose = true;
		verboseVerboseEnabled();
		break;
	case "f", "force":
		onExistingFile = 1;
		break;
	case "i", "interactive":
		onExistingFile = 2;
		break;
	case "n", "no-clobber": //case "no-overwrite":
		onExistingFile = 0;
		break;
	default:
		errUnknownOption(prefix + flag);
	}
}

func configInterpretBoolean(v string) bool {
	switch(v) {
		case "true", "on", "yes": return true;
		default: return false;
	}
}
