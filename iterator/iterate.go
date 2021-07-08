package iterator

import (
	"easy-copy/tasks"
	"os"
	"path/filepath"
)

func Add(p *tasks.Path) error {
	info, err := os.Lstat(p.AsAbs())
	if err != nil {
		return err
	}

	switch {
	case info.IsDir():
		return addAllInFolder(p)
	case info.Mode().IsRegular():
		tasks.AddTask(p)
	default:
		return &ErrBadType{p.AsAbs(), info}
	}

	return nil
}

// addAllInFolder creates tasks for all files in the specified folder.
func addAllInFolder(folder *tasks.Path) error {
	f, err := os.Open(folder.AsAbs())
	if err != nil {
		return err
	}

	defer f.Close()

	return walk(f, func(f string) {
		Add(&tasks.Path{Base: folder.Base, Sub: filepath.Join(folder.Sub, f)})
	})
}

// walk runs the given function for all files in the folder.
func walk(base *os.File, consumer func(string)) error {
	files, err := base.Readdirnames(0)
	if err != nil {
		return err
	}

	for _, f := range files {
		consumer(f)
	}

	return nil
}
