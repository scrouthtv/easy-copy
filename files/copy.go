package files

import (
	"easy-copy/flags"
	"easy-copy/progress"
	"easy-copy/ui"
	"errors"
	"fmt"
	"io"
	"os"
)

var buf []byte = make([]byte, 32678)

type InfoSetBuffersize struct {
	Size int
}

func (i *InfoSetBuffersize) Info() string {
	return fmt.Sprintf("buffersize set to %d b", i.Size) // TODO auto format size
}

func SetBuffersize(size int) {
	buf = make([]byte, size)
	ui.Infos <- &InfoSetBuffersize{size}
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

	return copyFile(s, d)
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
