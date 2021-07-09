package tasks

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

}
