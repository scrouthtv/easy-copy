package main;

import "os";
import "syscall";

// #include <linux/fs.h>
import "C";
// this defines C.FICLONE

/**
 * Takes two paths as arguments. Attempts to reflink dst to src.
 * If that did not work, an error is returned and the file has to
 * be copied manually
 */
func Reflink(srcPath string, dstPath string) error {
	var err error;
	var src, dst *os.File;
	src, err = os.OpenFile(srcPath, os.O_RDONLY, 0644);
	if err != nil { return err; }
	dst, err = os.OpenFile(dstPath, os.O_WRONLY | os.O_CREATE, 0644);
	if err != nil { return err; }

	err = reflinkInternal(dst, src);
	//if (err != nil) && fallback {
		// seek both src & dst at beginning
		//src.Seek(0, io.SeekStart)
		//dst.Seek(0, io.SeekStart)
		//dst.Truncate(0) // assuming any error in trucate will result in copy error
		//_, err = io.Copy(dst, src)
	//}
	return err;
}

func reflinkInternal(d, s *os.File) error {
	var ss, sd syscall.RawConn;
	var err error;
	ss, err = s.SyscallConn();
	if err != nil { return err; }
	sd, err = d.SyscallConn()
	if err != nil { return err; }

	var err2, err3 error;

	err = sd.Control(func(dfd uintptr) {
		err2 = ss.Control(func(sfd uintptr) {
			// int ioctl(int dest_fd, FICLONE, int src_fd);
			_, _, errno := syscall.Syscall(syscall.SYS_IOCTL, dfd, C.FICLONE, sfd);
			if errno != 0 {
				err3 = errno;
			}
		});
	});

	if err != nil {
		// sd.Control failed
		return err
	}
	if err2 != nil {
		// ss.Control failed
		return err2
	}

	// err3 is ioctl() response
	return err3
}
