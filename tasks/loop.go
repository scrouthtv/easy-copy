package tasks

import (
	"easy-copy/flags"
	"easy-copy/lscolors"
	"easy-copy/progress"
	"easy-copy/ui"
	"os"
)

func CopyLoop() {
	lock.Lock()
	switch {
	case len(folders) > 0:
		f := folders
		folders = make([]string, 0)
		lock.Unlock()
		createFolders(f)
	case len(solvedConflicts) > 0:
		lock.Unlock()
		t := PopSolvedConflict()
		work(t)
	case len(sources) > 0:
		lock.Unlock()
		t := PopTask()
		work(t)
	}
}

func work(t *Task) {
	// TODO
}

type ErrCreatingFolder struct {
	Path string
	Err  error
}

func (e *ErrCreatingFolder) Error() string {
	return "creating folder " + e.Path + ": " + e.Err.Error()
}

func (e *ErrCreatingFolder) Unwrap() error {
	return e.Err
}

func createFolders(f []string) {
	progress.CurrentTask = progress.TaskMkdir

	for _, folder := range f {
		if flags.Current.DoLSColors() {
			progress.CurrentFile = "\033[" + lscolors.FormatType("di") + "m" +
				folder + "\033[" + lscolors.FormatType("rs") + "m"
		} else {
			progress.CurrentFile = folder
		}

		if !flags.Current.Dryrun() {
			err := os.MkdirAll(folder, 0o755)
			if err != nil {
				ui.Warns <- &ErrCreatingFolder{folder, err}
			}
		}

		progress.DoneSize += uint64(progress.FolderSize)

	}
}
