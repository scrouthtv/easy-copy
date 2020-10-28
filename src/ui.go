package main

import "fmt"
import "strings"
import "os"

func printUsage() {
	fmt.Print(FGColors.LBlue)
	fmt.Print("Usage:")
	fmt.Println(Textstyle.Reset)

	fmt.Print(FGColors.LGray)
	fmt.Println("  ec [options] source target")
	fmt.Println("  ec [options] source ... directory")
	fmt.Print(FGColors.Default)
	fmt.Println("  ec --help")
	fmt.Print("  ec --version")
	fmt.Println(Textstyle.Reset)
}

func printHelp() {
	printVersion()
	fmt.Println()
	printUsage()
	fmt.Println()
	fmt.Print(FGColors.LBlue)
	fmt.Print("These options are supported:")
	fmt.Println()
	fmt.Print(Textstyle.Reset)
	const indent string = "                     "
	fmt.Println("  -f, --force      ", FGColors.LGray,
		"overwrite existing files without asking", FGColors.Default)
	fmt.Println("  -i, --interactive", FGColors.LGray,
		"ask before overwriting a file", FGColors.Default)
	fmt.Println("  -n, --no-clobber ", FGColors.LGray,
		"skip existing files", FGColors.Default)
	fmt.Println()
	fmt.Println("      --no-config  ", FGColors.LGray,
		"don't read any config file", Textstyle.Reset)
	fmt.Println("  -V, --verbose    ", FGColors.LGray,
		"verbose mode ", FGColors.Default)
	fmt.Println("      --color=WHEN ", FGColors.LGray,
		"whether to colorize the output.")
	fmt.Println(indent+"WHEN can be 'always', 'auto' (default) or 'never'",
		FGColors.Default)
	fmt.Println("")
	fmt.Println("  -h, --help       ", FGColors.LGray,
		"print this help and exit", FGColors.Default)
	fmt.Println("  -v, --version    ", FGColors.LGray,
		"print version information and exit", FGColors.Default)
	fmt.Print(Textstyle.Reset)
	fmt.Println()
	fmt.Print(FGColors.LBlue)
	fmt.Println("This is free software licensed under GNU GPL v3.0.")
	fmt.Println("You are welcome to redistribute it under certain conditions,")
	fmt.Print("run ")
	fmt.Print(FGColors.LGray + "ec --copying")
	fmt.Println(FGColors.LBlue + " for more information.")
	fmt.Println("This program is distributed with ABSOLUTELY NO WARRANTY,")
	fmt.Print("for details run ")
	fmt.Print(FGColors.LGray + "ec --warranty")
	fmt.Print(FGColors.LBlue + ".")
	fmt.Println(Textstyle.Reset)
	fmt.Println()
	fmt.Print(FGColors.LBlue)
	fmt.Print("Extensive documentation can be accessed through the manpages.")
	fmt.Println(Textstyle.Reset)
}

func printVersion() {
	fmt.Print(FGColors.Red)
	fmt.Print(EASYCOPY_NAME + " v" + EASYCOPY_VERSION)
	fmt.Println(Textstyle.Reset)
}

func printCopying() {
	runPager(infoCopying())
}

func printWarranty() {
	runPager(infoWarranty())
}

func printColortest() {
	fmt.Println(FGColors.Default + "Default")
	fmt.Println(FGColors.Black + "Black")
	fmt.Println(FGColors.Red + "Red")
	fmt.Println(FGColors.Green + "Green")
	fmt.Println(FGColors.Yellow + "Yellow")
	fmt.Println(FGColors.Blue + "Blue")
	fmt.Println(FGColors.Magenta + "Magenta")
	fmt.Println(FGColors.Cyan + "Cyan")
	fmt.Println(FGColors.LGray + "LGray")
	fmt.Println(FGColors.DGray + "DGray")
	fmt.Println(FGColors.LRed + "LRed")
	fmt.Println(FGColors.LGreen + "LGreen")
	fmt.Println(FGColors.LYellow + "LYellow")
	fmt.Println(FGColors.LBlue + "LBlue")
	fmt.Println(FGColors.LMagenta + "LMagenta")
	fmt.Println(FGColors.LCyan + "LCyan")
	fmt.Println(FGColors.White + "LWhite")
}

func verbVerboseEnabled() {
	fmt.Print(FGColors.Yellow + "Verbose mode enabled." + Textstyle.Reset)
}

func verbFlags() {
	if verbose {
		fmt.Printf(FGColors.Green)
		fmt.Println(" Verbose:", verbose)
		fmt.Println(" Overwrite Mode:", onExistingFile)
		fmt.Print(" Follow symlinks: ", followSymlinks)
		fmt.Println(Textstyle.Reset)
	}
}

func verbDisablingColors(shellname string) {
	// as this is called from the init function where verbose isn't set yet
	fmt.Println("Color support for", shellname, "is currently not ipmlemented.")
	fmt.Println("If your terminal does support colors, open an issue at")
	fmt.Println(" " + EASYCOPY_ISSUES)
}

func verbTargets() {
	if verbose {
		fmt.Print(FGColors.Yellow)
		fmt.Println("-------------------------")
		fmt.Println("these tasks will be done:")
		filesLock.RLock()
		var v string
		for _, v = range folders {
			fmt.Println("need to create folder", v)
		}
		for _, v = range fileOrder {
			var target string = targets[v]
			fmt.Println(v, "will be copied to", target+"/")
		}
		filesLock.RUnlock()
		fmt.Println("-------------------------")
		fmt.Print(Textstyle.Reset)
	}
}

func verbDoneIterating() {
	if verbose {
		fmt.Println(FGColors.Yellow + "All source files iterated." + Textstyle.Reset)
	}
}

func verbSearchStart() {
	if verbose {
		fmt.Print(FGColors.Yellow)
		fmt.Println("Have to search", unsearchedPaths)
		fmt.Print("Target is ", targetBase)
		fmt.Println(Textstyle.Reset)
	}
}

func verbCopyStart(sourcePath string, destPath string) {
	if verbose {
		fmt.Print(FGColors.Yellow)
		fmt.Print("src: ", sourcePath, " dest: ", destPath)
		fmt.Println(Textstyle.Reset)
	}
}

func verbCopyFinished(srcPath string, destPath string) {
	if verbose {
		fmt.Print(FGColors.Yellow)
		fmt.Print("finished copying ", srcPath, " to ", destPath)
		fmt.Println(Textstyle.Reset)
	}
}

func errCreatingFile(err error, file string) {
	fmt.Println("Could not create", file+":")
	fmt.Print(FGColors.Red)
	fmt.Print(err)
	fmt.Println(Textstyle.Reset)
	os.Exit(2)
}

func errCreatingLink(err error, source string, dest string) {
	fmt.Println("Error linking", source, "to", dest+":")
	fmt.Print(FGColors.Red)
	fmt.Print(err)
	fmt.Println(Textstyle.Reset)
	os.Exit(2)
}

func errMissingFile(err error, file string) {
	fmt.Println("Could not read", file+":")
	fmt.Print(FGColors.Red)
	fmt.Print(err)
	fmt.Println(Textstyle.Reset)
	os.Exit(2)
}

func errReadingSymlink(err error, link string) {
	fmt.Println("Could not resolve", link+":")
	fmt.Print(FGColors.Red)
	fmt.Print(err)
	fmt.Println(Textstyle.Reset)
	os.Exit(2)
}

func warnConfig(err error) {
	fmt.Println("Error while reading the config file:")
	fmt.Print(FGColors.LRed)
	fmt.Print(err)
	fmt.Println(Textstyle.Reset)
}

//func warnBadConfigValue(key string, given uint8, expected string) {
//fmt.Println("Error while reading the config file:");
//fmt.Print(FGColors.LRed);
//fmt.Print("Bad value for", key, "given", given, "but expected", expected);
//fmt.Println(Textstyle.Reset);
//fmt.Println("Reverting to default.");
//}

func warnBadConfigKey(key string) {
	fmt.Print(FGColors.LRed)
	fmt.Print("Unknown key ", key, " in the configuration file, skipping it.")
	fmt.Println(Textstyle.Reset)
}

func warnBadFile(file string) {
	fmt.Print(FGColors.LRed)
	fmt.Println(file, "is not a regular file, skipping it.")
	fmt.Println(Textstyle.Reset)
}

func verbReflinkFailed(sourcePath string, destPath string, err error) {
	if verbose {
		fmt.Print(FGColors.Yellow)
		fmt.Println("Error reflinking", sourcePath, "to", destPath+":")
		fmt.Println(err)
		fmt.Println("Going to copy it instead.")
		fmt.Print(Textstyle.Reset)
	}
}

func errReflinkFailed(sourcePath string, destPath string, err error) {
	fmt.Println("Error reflinking", sourcePath, "to", destPath+":")
	fmt.Print(FGColors.Red)
	fmt.Println(err)
	fmt.Print(Textstyle.Reset)
	os.Exit(2)
}

func errCopying(sourcePath string, destPath string, err error) {
	fmt.Println("Error copying", sourcePath, "to", destPath+":")
	fmt.Print(FGColors.Red)
	fmt.Print(err)
	fmt.Println(Textstyle.Reset)
	os.Exit(2)
}

func errUnknownOption(option string) {
	fmt.Print(FGColors.Red)
	fmt.Print("Unrecognized Option:", option)
	fmt.Println(Textstyle.Reset)
	printUsage()
	os.Exit(2)
}

func errEmptySource() {
	fmt.Print(FGColors.Red)
	fmt.Print("No sources specified.")
	fmt.Println(Textstyle.Reset)
	printUsage()
	os.Exit(2)
}

func errTargetNoDir(file string) {
	fmt.Print(FGColors.Red)
	fmt.Print(file, "is not a directory.")
	fmt.Println(Textstyle.Reset)
	os.Exit(2)
}

func errResolvingTarget(target string, err error) {
	fmt.Println("Cannot resolve", target, " as the target directory:")
	fmt.Print(FGColors.Red)
	fmt.Print(err)
	fmt.Println(Textstyle.Reset)
	os.Exit(2)
}

func warnCreatingConfig(err error) {
	fmt.Println("Could not create a default configuration file:")
	fmt.Print(FGColors.LRed)
	fmt.Print(err)
	fmt.Println(Textstyle.Reset)
}

func parseArgs() {
	args := os.Args[1:]
	var isFiles bool = false
	for _, arg := range args {
		if arg == "--" {
			isFiles = true
		} else if isFiles {
			unsearchedPaths = append(unsearchedPaths, arg)
		} else if strings.HasPrefix(arg, "--") {
			parseFlag("--", arg[2:])
		} else if strings.HasPrefix(arg, "-") {
			for i := 1; i < len(arg); i++ {
				parseFlag("-", arg[i:i+1])
			}
		} else {
			// TODO: clean & abs arg
			unsearchedPaths = append(unsearchedPaths, arg)
			uPTargets[arg] = targetBase
		}
	}
}
