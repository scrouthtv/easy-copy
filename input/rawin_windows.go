// +build windows
// +build !goin

package main

//#include <stdio.h>
//#include <conio.h>
import "C"

func Getch() rune {
	var ch int = int(C.getch())
	return rune(ch)
}
