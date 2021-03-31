// +build windows
// +build !goin

package input

//#include <stdio.h>
//#include <conio.h>
import "C"

func Enter() error {
	return nil
}

func Exit() {
	// nothing to do
}

func Getch() rune {
	var ch int = int(C.getch())
	return rune(ch)
}
