// +build freebsd openbsd netbsd dragonfly darwin
// +build !goin

package input

import "golang.org/x/sys/unix"

const (
	reqGetTermios = unix.TIOCGETA
	reqSetTermios = unix.TIOCSETA
)
