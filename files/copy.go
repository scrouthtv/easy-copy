package files

import (
	"easy-copy/common"
	"easy-copy/flags"
	"easy-copy/progress"
	"easy-copy/ui"
	"errors"
	"fmt"
	"io"
)

var buf []byte = make([]byte, 32678)

type InfoSetBuffersize struct {
	Size int
}

func (i *InfoSetBuffersize) Info() string {
	return fmt.Sprintf("buffersize set to %d b", i.Size) // TODO auto format size
}

func (i *InfoSetBuffersize) Required() flags.Verbose {
	return flags.VerbInfo
}

func SetBuffersize(size int) {
	buf = make([]byte, size)
	ui.Infos <- &InfoSetBuffersize{size}
}

type InfoStartCopy struct {
	Source string
	Dest   string
}

func (i *InfoStartCopy) Info() string {
	return fmt.Sprintf("copy %s to %s", i.Source, i.Dest)
}

func (i *InfoStartCopy) Required() flags.Verbose {
	return flags.VerbDebug
}

// copyFile copies the openend source file to the already
// created dest file.
func CopyFile(source common.File, dest common.File) error {
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
