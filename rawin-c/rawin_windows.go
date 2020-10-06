// +build windows
package main;

//#include <stdio.h> 
//#include <conio.h> 
import "C";

func enableRaw() {
	// nothing to do on windows!
}

func disableRaw() {
	// nothing to do on windows!
}

func getch() rune {
	var ch int = int(C.getch());
	return rune(ch);
}
