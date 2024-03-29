package tasks

import (
	"log"
	"easy-copy/common"
	"easy-copy/files"
	"easy-copy/flags"
	"easy-copy/lscolors"
	"easy-copy/progress"
	"easy-copy/ui"
)

func CopyLoop() {
	for !progress.CopyDone {
		lock.Lock()
		switch {
		case len(folders) > 0:
			f := folders
			folders = make([]string, 0)
			lock.Unlock()

			if flags.Current.Verbosity() >= flags.VerbDebug {
				log.Println("creating folders:", f)
			}

			createFolders(f)
		case len(solvedConflicts) > 0:
			lock.Unlock()
			t := PopSolvedConflict()

			if flags.Current.Verbosity() >= flags.VerbDebug {
				log.Println("solved conflict:", t)
			}

			work(t, flags.ConflictOverwrite)
		case len(sources) > 0:
			lock.Unlock()
			t := PopTask()

			if flags.Current.Verbosity() >= flags.VerbDebug {
				log.Println("basic task:", t)
			}

			work(t, flags.Current.OnConflict())
		default:
			if len(pendingConflicts) > 0 {
				lock.Unlock()
				continue
			}

			lock.Unlock()
			if progress.IteratorDone {
				progress.CopyDone = true
			}
		}
	}
}

type ErrMissingFile struct { // FIXME basically the same def as in iterator, maybe move all errors to a spearate package
	Path string
	Err  error
}

func (e *ErrMissingFile) Error() string {
	return "missing " + e.Path + ": " + e.Err.Error()
}

func (e *ErrMissingFile) Unwrap() error {
	return e.Err
}

type ErrCreatingFile struct {
	Path string
	Err  error
}

func (e *ErrCreatingFile) Error() string {
	return "creating " + e.Path + ": " + e.Err.Error()
}

func (e *ErrCreatingFile) Unwrap() error {
	return e.Err
}

func work(t *Task, onconflict flags.Conflict) {
	source, err := common.FileAdapter.Open(t.Source)
	if err != nil {
		ui.Error(&ErrMissingFile{t.Source, err})
	}

	defer source.Close()

	_, err = common.FileAdapter.Lstat(t.Dest)
	if err == nil {
		// dest exists
		if onconflict == flags.ConflictOverwrite {
			files.Syncdel(&[]string{t.Dest})
		} else if onconflict == flags.ConflictAsk {
			PushPendingConflict(*t)
			return
		} else if onconflict == flags.ConflictSkip {
			return
		}
	}

	switch flags.Current.Mode() {
	case flags.ModeCopy:
		dest, err := common.FileAdapter.Create(t.Dest) // TODO perms
		if err != nil {
			ui.Error(&ErrCreatingFile{t.Dest, err})
		}

		files.CopyFile(source, dest)
	case flags.ModeMove:
		err := files.Move(t.Source, t.Dest)
		if err != nil {
			ui.Error(&ErrCreatingFile{t.Dest, err})
		}
	case flags.ModeRemove:
		files.Syncdel(&[]string{t.Source})
	}

	progress.DoneAmount++
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
			err := common.FileAdapter.MkdirAll(folder, 0o755)
			if err != nil {
				ui.Warns <- &ErrCreatingFolder{folder, err}
			}
		}

		progress.DoneSize += uint64(progress.FolderSize)

	}
}
