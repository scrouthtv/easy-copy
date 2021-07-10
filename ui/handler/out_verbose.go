package handler

import (
	"easy-copy/color"
	"easy-copy/flags"
	"easy-copy/tasks"
	"easy-copy/ui"
	"fmt"
)

func VerbVerboseEnabled() {
	fmt.Println(color.FGColors.Yellow + "Verbose mode enabled." + color.Text.Reset)
}

func VerbDryrun() {
	fmt.Print(color.FGColors.Yellow)
	fmt.Print("Dry run - nothing on the disk will be changed.")
	fmt.Println(color.Text.Reset)
}

func VerbSetBuffersize(size int) {
	if flags.Current.Verbosity() >= flags.VerbDebug {
		fmt.Print(color.FGColors.Yellow + "Set buffersize to ")
		fmt.Print(ui.FormatSize(float64(size), ui.SizeAutoUnit(float64(size))))
		fmt.Println(color.Text.Reset)
	}
}

func VerbNativeMoveFailed(sourcePath string, destPath string, err error) {
	if flags.Current.Verbosity() >= flags.VerbDebug {
		fmt.Println(color.FGColors.Yellow + "Native moving")
		fmt.Print(sourcePath, "to", destPath, "failed:")
		fmt.Println(color.FGColors.Red + err.Error() + color.Text.Reset)
	}
}

func VerbTargets() {
	if flags.Current.Verbosity() >= flags.VerbInfo {
		fmt.Print(color.FGColors.Yellow)
		fmt.Println("-------------------------")
		fmt.Println("these tasks will be done:")

		tasks.PrintTasks()

		fmt.Println("-------------------------")
		fmt.Print(color.Text.Reset)
	}
}

func VerbDoneIterating() {
	if flags.Current.Verbosity() >= flags.VerbInfo {
		fmt.Println(color.FGColors.Yellow + "All source files iterated." + color.Text.Reset)
	}
}
