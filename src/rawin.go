// +build !linux,!windows

package main

func getch() rune {
	panic(FGColors.Red + "NOT SUPPORTED" + Textstyle.Reset)
}
