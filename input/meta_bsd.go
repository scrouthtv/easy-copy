//go:build !goin && (freebsd || openbsd || netbsd || dragonfly || darwin)

package input

import "golang.org/x/sys/unix"

const (
	reqGetTermios = unix.TIOCGETA
	reqSetTermios = unix.TIOCSETA
)
