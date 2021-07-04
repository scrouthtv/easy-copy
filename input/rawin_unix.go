// +build linux freebsd openbsd netbsd dragonfly darwin
// +build !goin

package input

import (
	"os"

	"golang.org/x/sys/unix"
)

var (
	in      int
	oldMode unix.Termios
	isRaw   bool = false
)

// Enter enters raw mode where every character is directly
// forwarded to us.
func Enter() error {
	if isRaw {
		return nil
	}

	in, err := unix.Open("/dev/tty", unix.O_RDONLY, 0)
	if err != nil {
		return err //nolint:wrapcheck // for now, we only care that an error occurred
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

	isRaw = true

	return unix.IoctlSetTermios(in, reqSetTermios, mode)
}

// Exit resets the terminal to the old mode.
func Exit() {
	if !isRaw {
		return
	}

	unix.IoctlSetTermios(in, reqSetTermios, &oldMode) //nolint:errcheck // nothing we can do about it
	unix.Close(in)

	isRaw = false
}

// Getch requires the terminal to be in raw mode and reads a single rune.
// If C-c is encountered, the program exits with code 8.
func Getch() rune {
	n, err := unix.Read(in, buf)

	if err != nil || n != 1 || buf[0] == 3 { // C-c
		os.Exit(errReturnCode)
	}

	return rune(buf[0])
}
