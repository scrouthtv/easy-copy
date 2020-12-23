// +build windows

package main

//#include <stdio.h>
//#include <conio.h>
import "C"

func getch() rune {
	var ch int = int(C.getch())
	return rune(ch)
}
