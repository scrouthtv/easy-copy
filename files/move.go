package files

import (
	"easy-copy/common"
)

func Move(source string, dest string) error {
	err := common.FileAdapter.Rename(source, dest)
	if err == nil { // native move successful
		return nil
	}

	// move "manually": copy + delete
	s, err := common.FileAdapter.Open(source)
	if err != nil {
		return err
	}

	d, err := common.FileAdapter.Create(dest)
	if err != nil {
		return err
	}

	CopyFile(s, d)
	Syncdel(&[]string{source})

	return nil
}
