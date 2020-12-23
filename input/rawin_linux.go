// +build linux
// +build rawin

package input

import "os"

//#include <ctype.h>
//#include <errno.h>
//#include <stdio.h>
//#include <stdlib.h>
//#include <termios.h>
//#include <unistd.h>
//
//struct termios orig_termios;
//
//void disableRawMode() {
//  if (tcsetattr(STDIN_FILENO, TCSAFLUSH, &orig_termios) == -1)
//    printf("tcsetattr");
//}
//
//void enableRawMode() {
//  if (tcgetattr(STDIN_FILENO, &orig_termios) == -1) printf("tcgetattr"); // TODO die
//
//  struct termios raw = orig_termios;
//  raw.c_iflag &= ~(BRKINT | ICRNL | INPCK | ISTRIP | IXON);
//  raw.c_oflag &= ~(OPOST);
//  raw.c_cflag |= (CS8);
//  raw.c_lflag &= ~(ECHO | ICANON | IEXTEN | ISIG);
//  raw.c_cc[VMIN] = 0;
//  raw.c_cc[VTIME] = 1;
//
//  if (tcsetattr(STDIN_FILENO, TCSAFLUSH, &raw) == -1) printf("tcsetattr");
//}
//
//char getOneKey() {
//	enableRawMode();
//
//	while (1) {
//		char c = '\0';
//		if (read(STDIN_FILENO, &c, 1) == -1 && errno != EAGAIN) printf("read");
//		if (c != '\0') {
//			disableRawMode();
//			return c;
//		}
//	}
//
//	return 0;
//}
import "C"

func Getch() rune {
	var in rune = rune(C.getOneKey())
	if in == rune(3) { // C-c
		os.Exit(8)
	}
	return in
}