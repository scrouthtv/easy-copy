package files

import (
	"easy-copy/flags"
	"easy-copy/progress"
	"easy-copy/ui/msg"
	"os"
)

// Syncdel deletes a list of files synchronously.
func Syncdel(files *[]string) {
	var err error

	for _, path := range *files {
		progress.CurrentTask = progress.TaskDelete
		progress.CurrentFile = path

		if !isnodelete(path) && !flags.Current.Dryrun() {
			err = os.RemoveAll(path)

			if err != nil {
				msg.ErrDeletingFile(path, err)
			}
		}
	}
}

func isnodelete(path string) bool {
	for _, p := range []string{} /* TODO nodelete */ {
		if p == path {
			return true
		}
	}

	return false
}
