package iterator

import (
	"easy-copy/flags"
	"easy-copy/progress"
	"easy-copy/tasks"
	"easy-copy/ui"
	"os"
	"path/filepath"
)

type ErrMissingFile struct {
	Path string
	Err  error
}

func (e *ErrMissingFile) Error() string {
	return "missing " + e.Path + ": " + e.Err.Error()
}

func (e *ErrMissingFile) Unwrap() error {
	return e.Err
}

// Iterate initializes the task manager with the target provided by
// the flags package. Afterwards, it iterates all sources.
func Iterate() {
	tasks.Setup(flags.Current.Target(), shouldCreateFolders())

	for _, p := range flags.Current.Sources() {
		err := add(&tasks.Path{Base: p, Sub: ""})
		if err != nil {
			ui.Warns <- &ErrMissingFile{p, err}
		}
	}

	if flags.Current.Verbosity() >= flags.VerbInfo {
		tasks.PrintTasks()
	}
	progress.IteratorDone = true
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
		tasks.AddFolder(p.Sub)
		progress.FullSize += uint64(progress.FolderSize)
		return addAllInFolder(p)
	case info.Mode().IsRegular():
		progress.FullSize += uint64(info.Size())
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
