package tasks

import (
	"easy-copy/flags"
	"easy-copy/lscolors"
	"easy-copy/progress"
	"easy-copy/ui/msg"
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
	case len(sources) > 0:
		lock.Unlock()
		t := PopTask()
	}
}

func createFolders(f []string) {
	progress.CurrentTask = progress.TaskFolder

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
				msg.ErrCreatingFile(err, folder)
			}
		}

		progress.DoneSize += progress.FolderSize

	}
}
