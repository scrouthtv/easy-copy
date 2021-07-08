package ui

import (
	"easy-copy/color"
	"fmt"
)

func PrintUsage() {
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

func PrintHelp() {
	PrintVersion()
	fmt.Println()
	PrintUsage()
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

func PrintVersion() {
	fmt.Print(color.FGColors.Red)
	fmt.Print(EasyCopyName + " v" + EasyCopyVersion)
	fmt.Println(color.Text.Reset)
}
