package files

import (
	"easy-copy/flags"
	"easy-copy/progress"
	"easy-copy/ui/msg"
	"errors"
	"io"
	"os"
)

var buf []byte = make([]byte, 32678)

func SetBuffersize(size int) {
	buf = make([]byte, size)
	msg.VerbSetBuffersize(size)
}

func Copy(source string, dest string) error {
	s, err := os.Open(source)
	if err != nil {
		return err
	}

	defer s.Close()

	d, err := os.Create(dest)
	if err != nil {
		return err
	}

	defer d.Close()

	copyFile(s, d)

	return nil
}

// copyFile copies the openend source file to the already
// created dest file.
func copyFile(source *os.File, dest *os.File) error {
	var readAmount, writtenAmount int
	var err error

	for {
		readAmount, err = source.Read(buf)

		if err != nil && !errors.Is(err, io.EOF) {
			return err
		}

		if readAmount == 0 {
			// when the file is fully read
			break
		}

		if !flags.Current.Dryrun() {
			writtenAmount, err = dest.Write(buf[:readAmount])
			if err != nil {
				return err
			}

			if readAmount != writtenAmount {
				return &ErrWritingData{read: readAmount, written: writtenAmount}
			}
		}

		progress.DoneSize += uint64(writtenAmount)
	}

	return nil
}
