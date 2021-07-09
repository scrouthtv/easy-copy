package iterator

import (
	"easy-copy/flags"
	"easy-copy/tasks"
	"os"
	"path/filepath"
)

// Iterate initializes the task manager with the target provided by
// the flags package. Afterwards, it iterates all sources.
func Iterate() {
	tasks.Setup(flags.Current.Target(), shouldCreateFolders())

	for _, p := range flags.Current.Sources() {
		add(&tasks.Path{Base: p, Sub: ""})
	}
}

func shouldCreateFolders() bool {
	if len(flags.Current.Sources()) > 1 {
		return true
	}

	_, err := os.Open(flags.Current.Target())
	return err == nil // if target already exists, create folders inside it
}

func add(p *tasks.Path) error {
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
		add(&tasks.Path{Base: folder.Base, Sub: filepath.Join(folder.Sub, f)})
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
