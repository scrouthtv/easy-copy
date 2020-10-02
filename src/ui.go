package main

import "fmt"
import "strings"
import "os"
import "path/filepath"

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

func fmtFile(file string) string {
	path, err := filepath.Abs(file);
	if (err != nil) {
		fmt.Println("Can't reda file:", err);
		os.Exit(8);
	}
	return path;
}

func errUnknownOption(option string) {
	fmt.Println("unrecognized option:", option);
	printUsage();
	os.Exit(2);
}

func errEmptySource() {
	fmt.Println("No sources specified.")
	printUsage();
	os.Exit(1);
}

func errInvalidWD(err error) {
	fmt.Println("The current directory is invalid:", err);
	os.Exit(4);
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
