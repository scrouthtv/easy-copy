package handler

import (
	"easy-copy/color"
	"easy-copy/ui"
	"fmt"
	"os"
)

func ErrCreatingFile(err error, file string) {
	fmt.Println("Could not create", file+":")
	fmt.Print(color.FGColors.Red)
	fmt.Print(err)
	fmt.Println(color.Text.Reset)
	os.Exit(2)
}

func ErrCreatingLink(err error, source string, dest string) {
	fmt.Println("Error linking", source, "to", dest+":")
	fmt.Print(color.FGColors.Red)
	fmt.Print(err)
	fmt.Println(color.Text.Reset)
	os.Exit(2)
}

func ErrMissingFile(err error, file string) {
	fmt.Println("Could not read", file+":")
	fmt.Print(color.FGColors.Red)
	fmt.Print(err)
	fmt.Println(color.Text.Reset)
	os.Exit(2)
}

func ErrReadingSymlink(err error, link string) {
	fmt.Println("Could not resolve", link+":")
	fmt.Print(color.FGColors.Red)
	fmt.Print(err)
	fmt.Println(color.Text.Reset)
	os.Exit(2)
}

func WarnConfig(err error) {
	fmt.Println("Error while reading the config file:")
	fmt.Print(color.FGColors.LRed)
	fmt.Print(err)
	fmt.Println(color.Text.Reset)
}

func WarnBadConfigKey(key string) {
	fmt.Print(color.FGColors.LRed)
	fmt.Print("Unknown key ", key, " in the configuration file, skipping it.")
	fmt.Println(color.Text.Reset)
}

func WarnBadFile(file string) {
	fmt.Print(color.FGColors.LRed)
	fmt.Print(file, " is not a regular file, skipping it.")
	fmt.Println(color.Text.Reset)
}

func WarnWormhole(src string) {
	fmt.Print(color.FGColors.LRed)
	fmt.Print("Cannot copy ", src, " into itself, skipping it.")
	fmt.Println(color.Text.Reset)
}

func ErrCopying(sourcePath string, destPath string, err error) {
	fmt.Println("Error copying", sourcePath, "to", destPath+":")
	fmt.Print(color.FGColors.Red)
	fmt.Print(err)
	fmt.Println(color.Text.Reset)
	os.Exit(2)
}

func ErrUnknownOption(option string) {
	fmt.Print(color.FGColors.Red)
	fmt.Print("Unrecognized Option: ", option)
	fmt.Println(color.Text.Reset)
	ui.PrintUsage()
	os.Exit(2)
}

func ErrMissingOperation() {
	fmt.Print(color.FGColors.Red)
	fmt.Print("No operation specified")
	fmt.Println(color.Text.Reset)
	ui.PrintUsage()
	os.Exit(2)
}

func ErrEmptySource() {
	fmt.Print(color.FGColors.Red)
	fmt.Print("No sources specified.")
	fmt.Println(color.Text.Reset)
	ui.PrintUsage()
	os.Exit(2)
}

func ErrTargetNoDir(file string) {
	fmt.Print(color.FGColors.Red)
	fmt.Print(file, " is not a directory.")
	fmt.Println(color.Text.Reset)
	os.Exit(2)
}

func ErrResolvingTarget(target string, err error) {
	fmt.Println("Cannot resolve", target, " as the target directory:")
	fmt.Print(color.FGColors.Red)
	fmt.Print(err)
	fmt.Println(color.Text.Reset)
	os.Exit(2)
}

func ErrInvalidMode(given string, expected string) {
	fmt.Print(color.FGColors.Red)
	fmt.Println("Invalid mode", given+", expected one of")
	fmt.Print(expected)
	fmt.Println(color.Text.Reset)
	os.Exit(2)
}

func ErrDeletingFile(path string, err error) {
	fmt.Print(color.FGColors.Red)
	fmt.Println("Error deleting", path+":")
	fmt.Print(err)
	fmt.Println(color.Text.Reset)
	os.Exit(2)
}

func ErrStall() {
	fmt.Print(color.FGColors.Red)
	fmt.Print("Aborting because less than 8 b have been transferred in 1 minute")
	fmt.Println(color.Text.Reset)
	os.Exit(2)
}
