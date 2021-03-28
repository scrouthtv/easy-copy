// +build linux
// +build !goin

package input

import "golang.org/x/sys/unix"

const (
	reqGetTermios = unix.TCGETS
	reqSetTermios = unix.TCSETS
)
