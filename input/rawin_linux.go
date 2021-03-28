// +build linux freebsd openbsd netbsd dragonfly darwin
// +build !goin

package input

import (
	"os"

	"golang.org/x/sys/unix"
)

var in int
var oldMode unix.Termios
var buf []byte = make([]byte, 8)

func enter() error {
	in, err := unix.Open("/dev/tty", unix.O_RDONLY, 0)
	if err != nil {
		return err
	}

	mode, err := unix.IoctlGetTermios(in, reqGetTermios)
	if err != nil {
		unix.Close(in)
	}

	oldMode = *mode

	mode.Iflag &^= unix.BRKINT | unix.ICRNL | unix.INPCK | unix.ISTRIP | unix.IXON
	mode.Oflag &^= unix.OPOST
	mode.Lflag &^= unix.ECHO | unix.ECHONL | unix.ICANON | unix.ISIG | unix.IEXTEN
	mode.Cflag &^= unix.CSIZE | unix.PARENB
	mode.Cflag |= unix.CS8
	mode.Cc[unix.VMIN] = 1
	mode.Cc[unix.VTIME] = 0

	return unix.IoctlSetTermios(in, reqSetTermios, mode)
}

func exit() {
	unix.IoctlSetTermios(in, reqSetTermios, &oldMode)
	unix.Close(in)
}

func Getch() rune {
	err := enter()
	if err != nil {
		os.Exit(8)
	}

	n, err := unix.Read(in, buf)
	exit()

	if err != nil || n != 1 || buf[0] == 3 { // C-c
		os.Exit(8)
	}

	return rune(buf[0])
}
