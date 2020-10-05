package main;

import "fmt";
import "strings";
import "os";

func printUsage() {
	fmt.Print(FGColors.LBlue);
	fmt.Print("Usage:");
	fmt.Println(Textstyle.Reset);

	fmt.Print(FGColors.LGray);
	fmt.Println("  ec [options] source target");
	fmt.Println("  ec [options] source ... directory");
	fmt.Print(FGColors.Default);
	fmt.Println("  ec --help");
	fmt.Print("  ec --version");
	fmt.Println(Textstyle.Reset);
}

func printHelp() {
	fmt.Print(FGColors.LMagenta);
	fmt.Print("This is ", EASYCOPY_NAME, " v", EASYCOPY_VERSION + ".");
	fmt.Println(Textstyle.Reset);
	fmt.Println();
	printUsage();
	fmt.Println();
	fmt.Print(FGColors.LBlue);
	fmt.Print("These options are supported:");
	fmt.Println();
	fmt.Print(Textstyle.Reset);
	fmt.Println("  -f, --force      ", FGColors.LGray,
		"overwrite existing files without asking", FGColors.Default);
	fmt.Println("  -i, --interactive", FGColors.LGray,
		"ask before overwriting a file", FGColors.Default);
	fmt.Println("  -n, --no-clobber ", FGColors.LGray,
		"skip existing files", FGColors.Default);
	fmt.Println("");
	fmt.Println("  -V, --verbose    ", FGColors.LGray,
		"verbose mode ", FGColors.Default);
	fmt.Println("");
	fmt.Println("  -h, --help       ", FGColors.LGray,
		"print this help and exit", FGColors.Default);
	fmt.Println("  -v, --version    ", FGColors.LGray,
		"print version information and exit", FGColors.Default);
	fmt.Print(Textstyle.Reset);
}

func printVersion() {
	fmt.Print(FGColors.Red);
	fmt.Print(EASYCOPY_NAME, Textstyle.Bold);
	fmt.Print(" v" + EASYCOPY_VERSION);
	fmt.Println(Textstyle.Reset);
}

func verboseVerboseEnabled() {
	fmt.Print(FGColors.Yellow);
	fmt.Print("Verbose mode enabled.");
	fmt.Println(Textstyle.Reset);
}

func verboseDisablingColors(shellname string) {
	// as this is called from the init function where verbose isn't set yet
	fmt.Println("Color support for", shellname, "is currently not ipmlemented.");
	fmt.Println("If your terminal does support colors, open an issue at");
	fmt.Println(" " + EASYCOPY_ISSUES);
}

func verboseTargets() {
	if verbose {
		fmt.Print(FGColors.Yellow);
		fmt.Println("-------------------------");
		fmt.Println("these tasks will be done:");
		filesLock.RLock();
		var v string;
		for _, v = range folders {
			fmt.Println("need to create folder", v);
		}
		for _, v = range fileOrder {
			var target string = targets[v];
			fmt.Println(v, "will be copied to", target);
		}
		filesLock.RUnlock();
		fmt.Println("-------------------------");
		fmt.Print(Textstyle.Reset);
	}
}

func errCreatingFile(err error, file string) {
	fmt.Println("Could not create", file + ":");
	fmt.Print(FGColors.Red);
	fmt.Print(err);
	fmt.Println(Textstyle.Reset);
	os.Exit(2);
}

func errMissingFile(err error, file string) {
	fmt.Println("Could not read", file + ":");
	fmt.Print(FGColors.Red);
	fmt.Print(err);
	fmt.Println(Textstyle.Reset);
	os.Exit(2);
}

func errReadingSymlink(err error, link string) {
	fmt.Println("Could not resolve", link + ":");
	fmt.Print(FGColors.Red);
	fmt.Print(err);
	fmt.Println(Textstyle.Reset);
	os.Exit(2);
}

func warnConfig(err error) {
	fmt.Println("Error while reading the config file:");
	fmt.Print(FGColors.LRed);
	fmt.Print(err);
	fmt.Println(Textstyle.Reset);
}

func warnBadConfig(key string, given uint8, expected string) {
	fmt.Println("Error while reading the config file:");
	fmt.Print(FGColors.LRed);
	fmt.Print("Bad value for", key, "given", given, "but expected", expected);
	fmt.Println(Textstyle.Reset);
	fmt.Println("Reverting to default.");
}

func warnBadFile(file string) {
	fmt.Print(FGColors.LRed);
	fmt.Println(file, "is not a regular file, skipping it.");
	fmt.Println(Textstyle.Reset);
}

func errUnknownOption(option string) {
	fmt.Print(FGColors.Red);
	fmt.Print("Unrecognized Option:", option);
	fmt.Println(Textstyle.Reset);
	printUsage();
	os.Exit(2);
}

func errEmptySource() {
	fmt.Print(FGColors.Red);
	fmt.Print("No sources specified.")
	fmt.Println(Textstyle.Reset);
	printUsage();
	os.Exit(2);
}

func errTargetNoDir(file string) {
	fmt.Print(FGColors.Red);
	fmt.Print(file, "is not a directory.");
	fmt.Println(Textstyle.Reset);
	os.Exit(2);
}

func errInvalidWD(err error) {
	fmt.Println("The current directory is invalid:");
	fmt.Print(FGColors.Red);
	fmt.Print(err);
	fmt.Println(Textstyle.Reset);
	os.Exit(2);
}

func errResolvingTarget(target string, err error) {
	fmt.Println("Cannot resolve", target, " as the target directory:")
	fmt.Print(FGColors.Red);
	fmt.Print(err);
	fmt.Println(Textstyle.Reset);
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
