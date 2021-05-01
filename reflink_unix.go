// +build linux freebsd netbsd openbsd darwin dragonfly

package main

import (
	"os"
	"syscall"
)

// #include <linux/fs.h>
import "C"

// this defines C.FICLONE

// reflink takes two paths as arguments. Attempts to reflink dst to
// src. If that did not work, an error is returned and the file has to
// be copied manually
func reflink(srcPath string, dstPath string, progressStorage *uint64) error {
	var err error
	var src, dst *os.File
	src, err = os.OpenFile(srcPath, os.O_RDONLY, 0o644)
	if err != nil {
		return err
	}
	dst, err = os.OpenFile(dstPath, os.O_WRONLY|os.O_CREATE, 0o644)
	if err != nil {
		return err
	}

	var ss, sd syscall.RawConn
	ss, err = src.SyscallConn()
	if err != nil {
		return err
	}
	sd, err = dst.SyscallConn()
	if err != nil {
		return err
	}

	var err2, err3 error

	err = sd.Control(func(dfd uintptr) {
		err2 = ss.Control(func(sfd uintptr) {
			_, _, errno := syscall.Syscall(syscall.SYS_IOCTL, dfd, C.FICLONE, sfd)
			if errno != 0 {
				err3 = errno
			}
		})
	})

	if err != nil {
		// sd.Control failed
		return err
	}
	if err2 != nil {
		// ss.Control failed
		return err2
	}

	if err3 != nil {
		// write progress:
		stat, _ := dst.Stat() //nolint:errcheck // this must work at this point
		*progressStorage += uint64(stat.Size())
	}

	// err3 is ioctl() response
	return err3
}