package main

import (
	"easy-copy/color"
	"errors"
	"fmt"
	"os"
	"strings"
)

func printUsage() {
	fmt.Print(color.FGColors.LBlue)
	fmt.Print("Usage:")
	fmt.Println(color.Text.Reset)

	fmt.Print(color.FGColors.LGray)
	fmt.Println("  ec OPERATION [options] source target")
	fmt.Println("  ec OPERATION [options] source ... directory")
	fmt.Print(color.FGColors.Default)
	fmt.Println("  ec --help")
	fmt.Print("  ec --version")
	fmt.Println(color.Text.Reset)
}

func printHelp() {
	printVersion()
	fmt.Println()
	printUsage()
	fmt.Println()
	fmt.Print(color.FGColors.LBlue)
	fmt.Print("These options are supported:")
	fmt.Println()
	fmt.Print(color.Text.Reset)
	const indent = "                     "

	fmt.Println("  -f, --force      ", color.FGColors.LGray,
		"overwrite existing files without asking", color.FGColors.Default)
	fmt.Println("  -i, --interactive", color.FGColors.LGray,
		"ask before overwriting a file", color.FGColors.Default)
	fmt.Println("  -n, --no-clobber ", color.FGColors.LGray,
		"skip existing files", color.FGColors.Default)
	fmt.Println()
	fmt.Println("      --no-config  ", color.FGColors.LGray,
		"don't read any config file", color.Text.Reset)
	fmt.Println("  -V, --verbose    ", color.FGColors.LGray,
		"verbose mode", color.FGColors.Default)
	fmt.Println("  -q, --quiet      ", color.FGColors.LGray,
		"quiet mode", color.FGColors.Default)
	fmt.Println("      --color=WHEN ", color.FGColors.LGray,
		"whether to colorize the output.")
	fmt.Println(indent+"WHEN can be 'always', 'auto' (default) or 'never'",
		color.FGColors.Default)
	fmt.Println("")
	fmt.Println("  -h, --help       ", color.FGColors.LGray,
		"print this help and exit", color.FGColors.Default)
	fmt.Println("  -v, --version    ", color.FGColors.LGray,
		"print version information and exit", color.FGColors.Default)
	fmt.Print(color.Text.Reset)
	fmt.Println()
	fmt.Print(color.FGColors.LBlue)
	fmt.Println("This is free software licensed under GNU GPL v3.0.")
	fmt.Println("You are welcome to redistribute it under certain conditions,")
	fmt.Print("run ")
	fmt.Print(color.FGColors.LGray + "ec --copying")
	fmt.Println(color.FGColors.LBlue + " for more information.")
	fmt.Println("This program is distributed with ABSOLUTELY NO WARRANTY,")
	fmt.Print("for details run ")
	fmt.Print(color.FGColors.LGray + "ec --warranty")
	fmt.Print(color.FGColors.LBlue + ".")
	fmt.Println(color.Text.Reset)
	fmt.Println()
	fmt.Print(color.FGColors.LBlue)
	fmt.Print("Extensive documentation can be accessed through the manpages.")
	fmt.Println(color.Text.Reset)
}

func printVersion() {
	fmt.Print(color.FGColors.Red)
	fmt.Print(EasyCopyName + " v" + EasyCopyVersion)
	fmt.Println(color.Text.Reset)
}

func printCopying() {
	_, err := runPager(infoCopying())
	if errors.Is(err, errNoPager) {
		fmt.Println(infoCopying())
	} else if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func printWarranty() {
	_, err := runPager(infoWarranty())
	if errors.Is(err, errNoPager) {
		fmt.Println(infoCopying())
	} else if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func printColortest() {
	fmt.Println(color.FGColors.Default + "Default")
	fmt.Println(color.FGColors.Black + "Black")
	fmt.Println(color.FGColors.Red + "Red")
	fmt.Println(color.FGColors.Green + "Green")
	fmt.Println(color.FGColors.Yellow + "Yellow")
	fmt.Println(color.FGColors.Blue + "Blue")
	fmt.Println(color.FGColors.Magenta + "Magenta")
	fmt.Println(color.FGColors.Cyan + "Cyan")
	fmt.Println(color.FGColors.LGray + "LGray")
	fmt.Println(color.FGColors.DGray + "DGray")
	fmt.Println(color.FGColors.LRed + "LRed")
	fmt.Println(color.FGColors.LGreen + "LGreen")
	fmt.Println(color.FGColors.LYellow + "LYellow")
	fmt.Println(color.FGColors.LBlue + "LBlue")
	fmt.Println(color.FGColors.LMagenta + "LMagenta")
	fmt.Println(color.FGColors.LCyan + "LCyan")
	fmt.Println(color.FGColors.White + "LWhite")
}

func verbVerboseEnabled() {
	fmt.Println(color.FGColors.Yellow + "Verbose mode enabled." + color.Text.Reset)
}

func verbSetBuffersize(size int) {
	if verbose >= VerbDebug {
		fmt.Print(color.FGColors.Yellow + "Set buffersize to ")
		fmt.Print(formatSize(float64(size), sizeAutoUnit(float64(size))))
		fmt.Println(color.Text.Reset)
	}
}

func verbNativeMoveFailed(sourcePath string, destPath string, err error) {
	if verbose >= VerbDebug {
		fmt.Println(color.FGColors.Yellow + "Native moving")
		fmt.Print(sourcePath, "to", destPath, "failed:")
		fmt.Println(color.FGColors.Red + err.Error() + color.Text.Reset)
	}
}

func verbFlags() {
	if verbose >= VerbInfo {
		fmt.Printf(color.FGColors.Green)
		fmt.Println(" Verbose:", verbose)
		fmt.Println(" Overwrite Mode:", onExistingFile)
		fmt.Print(" Follow symlinks: ", followSymlinks)
		fmt.Println(color.Text.Reset)
	}
}

func verbTargets() {
	if verbose >= VerbInfo {
		fmt.Print(color.FGColors.Yellow)
		fmt.Println("-------------------------")
		fmt.Println("these tasks will be done:")
		filesLock.RLock()

		for _, v := range folders {
			fmt.Println("need to create folder", v)
		}

		for _, v := range fileOrder {
			fmt.Println(v, "will be copied into", targets[v]+"/")
		}

		filesLock.RUnlock()
		fmt.Println("-------------------------")
		fmt.Print(color.Text.Reset)
	}
}

func verbDoneIterating() {
	if verbose >= VerbInfo {
		fmt.Println(color.FGColors.Yellow + "All source files iterated." + color.Text.Reset)
	}
}

func verbSearchStart() {
	if verbose >= VerbDebug {
		fmt.Print(color.FGColors.Yellow)
		fmt.Println("Have to search", unsearchedPaths)
		fmt.Print("Target is ", targetBase)
		fmt.Println(color.Text.Reset)
	}
}

func errCreatingFile(err error, file string) {
	fmt.Println("Could not create", file+":")
	fmt.Print(color.FGColors.Red)
	fmt.Print(err)
	fmt.Println(color.Text.Reset)
	os.Exit(2)
}

func errCreatingLink(err error, source string, dest string) {
	fmt.Println("Error linking", source, "to", dest+":")
	fmt.Print(color.FGColors.Red)
	fmt.Print(err)
	fmt.Println(color.Text.Reset)
	os.Exit(2)
}

func errMissingFile(err error, file string) {
	fmt.Println("Could not read", file+":")
	fmt.Print(color.FGColors.Red)
	fmt.Print(err)
	fmt.Println(color.Text.Reset)
	os.Exit(2)
}

func errReadingSymlink(err error, link string) {
	fmt.Println("Could not resolve", link+":")
	fmt.Print(color.FGColors.Red)
	fmt.Print(err)
	fmt.Println(color.Text.Reset)
	os.Exit(2)
}

func warnConfig(err error) {
	fmt.Println("Error while reading the config file:")
	fmt.Print(color.FGColors.LRed)
	fmt.Print(err)
	fmt.Println(color.Text.Reset)
}

func warnBadConfigKey(key string) {
	fmt.Print(color.FGColors.LRed)
	fmt.Print("Unknown key ", key, " in the configuration file, skipping it.")
	fmt.Println(color.Text.Reset)
}

func warnBadFile(file string) {
	fmt.Print(color.FGColors.LRed)
	fmt.Println(file, "is not a regular file, skipping it.")
	fmt.Println(color.Text.Reset)
}

func errCopying(sourcePath string, destPath string, err error) {
	fmt.Println("Error copying", sourcePath, "to", destPath+":")
	fmt.Print(color.FGColors.Red)
	fmt.Print(err)
	fmt.Println(color.Text.Reset)
	os.Exit(2)
}

func errUnknownOption(option string) {
	fmt.Print(color.FGColors.Red)
	fmt.Print("Unrecognized Option: ", option)
	fmt.Println(color.Text.Reset)
	printUsage()
	os.Exit(2)
}

func errMissingOperation() {
	fmt.Print(color.FGColors.Red)
	fmt.Print("No operation specified")
	fmt.Println(color.Text.Reset)
	printUsage()
	os.Exit(2)
}

func errEmptySource() {
	fmt.Print(color.FGColors.Red)
	fmt.Print("No sources specified.")
	fmt.Println(color.Text.Reset)
	printUsage()
	os.Exit(2)
}

func errTargetNoDir(file string) {
	fmt.Print(color.FGColors.Red)
	fmt.Print(file, " is not a directory.")
	fmt.Println(color.Text.Reset)
	os.Exit(2)
}

func errResolvingTarget(target string, err error) {
	fmt.Println("Cannot resolve", target, " as the target directory:")
	fmt.Print(color.FGColors.Red)
	fmt.Print(err)
	fmt.Println(color.Text.Reset)
	os.Exit(2)
}

func errInvalidMode(given string, expected string) {
	fmt.Print(color.FGColors.Red)
	fmt.Println("Invalid mode", given+", expected one of")
	fmt.Print(expected)
	fmt.Println(color.Text.Reset)
	os.Exit(2)
}

func errDeletingFile(path string, err error) {
	fmt.Print(color.FGColors.Red)
	fmt.Println("Error deleting", path+":")
	fmt.Print(err)
	fmt.Println(color.Text.Reset)
	os.Exit(2)
}

func parseMode() {
	if len(os.Args) < 2 {
		errMissingOperation()
	}

	switch strings.ToLower(os.Args[1]) {
	case "cp":
		mode = ModeCopy
	case "mv":
		mode = ModeMove
	case "rm":
		mode = ModeRemove

		errInvalidMode(strings.ToLower(os.Args[1]), "cp, mv")
	default:
		errInvalidMode(strings.ToLower(os.Args[1]), "cp, mv")
	}
}

func parseArgs() {
	parseMode()

	args := os.Args[2:]
	isFiles := false

	for _, arg := range args {
		if isFiles {
			unsearchedPaths = append(unsearchedPaths, arg)
		} else if arg == "--" {
			isFiles = true
		} else if strings.HasPrefix(arg, "--") {
			parseFlag("--", arg[2:])
		} else if strings.HasPrefix(arg, "-") {
			for i := 1; i < len(arg); i++ {
				parseFlag("-", arg[i:i+1])
			}
		} else {
			unsearchedPaths = append(unsearchedPaths, arg)
			uPTargets[arg] = targetBase
		}
	}
}
