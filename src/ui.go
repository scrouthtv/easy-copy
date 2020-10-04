package main

import "fmt"
import "strings"
import "os"

func printUsage() {
	fmt.Println("Usage:");
	fmt.Println("  ec [options] source target");
	fmt.Println("  ec --help");
	fmt.Println("  ec --version");
}

func printHelp() {
	fmt.Println("This is", EASYCOPY_NAME, "v" + EASYCOPY_VERSION + ".");
	fmt.Println("");
	printUsage();
	fmt.Println("");
	fmt.Println("These options are supported:");
	fmt.Println("  -f, --force        overwrite existing files without asking");
	fmt.Println("  -i, --interactive  ask before overwriting a file");
	fmt.Println("  -n, --no-clobber   skip existing files");
	fmt.Println("");
	fmt.Println("  -V, --verbose      verbose mode ")
	fmt.Println("");
	fmt.Println("  -h, --help         print this help and exit");
	fmt.Println("  -v, --version      print version information and exit");
}

func printVersion() {
	fmt.Println(EASYCOPY_NAME, "v" + EASYCOPY_VERSION);
}

func verboseTargets() {
	var v string;
	for _, v = range folders {
		fmt.Println("need to create folder", v);
	}
	for _, v = range fileOrder {
		var target string = targets[v];
		fmt.Println(v, "will be copied to", target);
	}
}

func errMissingFile(err error, file string) {
	fmt.Println("Could not read", file + ":");
	fmt.Println(err);
	os.Exit(2);
}

func errReadingSymlink(err error, link string) {
	fmt.Println("Could not resolve", link + ":");
	fmt.Println(err);
	os.Exit(2);
}

func warnConfig(err error) {
	fmt.Println("Error while reading the config file:", err)
}

func warnBadConfig(key string, given uint8, expected string) {
	fmt.Println("Error while reading the config file:");
	fmt.Println("Bad value for", key, "given", given, "but expected", expected);
}

func warnBadFile(file string) {
	fmt.Println(file, "is not a regular file, skipping it.");
}

func errUnknownOption(option string) {
	fmt.Println("Unrecognized Option:", option);
	printUsage();
	os.Exit(2);
}

func errEmptySource() {
	fmt.Println("No sources specified.")
	printUsage();
	os.Exit(2);
}

func errTargetNoDir(file string) {
	fmt.Println(file, "is not a directory.");
	os.Exit(2);
}

func errInvalidWD(err error) {
	fmt.Println("The current directory is invalid:");
	fmt.Println(err);
	os.Exit(2);
}

func errResolvingTarget(target string, err error) {
	fmt.Println("Cannot resolve", target, " as the target directory:")
	fmt.Println(err);
	os.Exit(2);
}

func parseArgs() {
	args := os.Args[1:];
	var isFiles bool = false;
	for _, arg := range args {
		if (arg == "--") {
			isFiles = true;
		} else if (isFiles) {
			unsearchedPaths = append(unsearchedPaths, arg);
		} else if (strings.HasPrefix(arg, "--")) {
			parseFlag("--", arg[2:len(arg)]);
		} else if (strings.HasPrefix(arg, "-")) {
			for i := 1; i < len(arg); i++ {
				parseFlag("-", arg[i:i+1]);
			}
		} else {
			unsearchedPaths = append(unsearchedPaths, arg);
		}
	}
}
