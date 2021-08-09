package tasks

import (
	"easy-copy/flags"
	"easy-copy/progress"
	"fmt"
	"path/filepath"
)

var (
	targetBase            string
	createFoldersInTarget bool
)

var (
	lock                              *LoggedLock = newLock()
	sources                           []Path
	pendingConflicts, solvedConflicts []Task
	folders                           []string
)

type Task struct {
	Source string
	Dest   string
}

// Setup is called once for setting the target base
// and whether to recreate the root folders inside the target.
func Setup(base string, cloneFolders bool) {
	targetBase = base
	createFoldersInTarget = cloneFolders

	sources = nil
	pendingConflicts = nil
	solvedConflicts = nil
	folders = nil

	addFolder(base)

	if flags.Current.Verbosity() >= flags.VerbDebug {
		fmt.Println("create folders in target:", cloneFolders)
	}
}

// PopTask pops the next available task from the task list.
func PopTask() *Task {
	lock.Lock()
	if len(sources) == 0 {
		lock.Unlock()
		return nil
	}

	next := sources[0]
	sources = sources[1:]
	lock.Unlock()

	return &Task{Source: next.AsAbs(), Dest: destFor(&next)}
}

func PopPendingConflict() *Task {
	lock.Lock()
	defer lock.Unlock()
	if len(pendingConflicts) == 0 {
		return nil
	}

	pop := pendingConflicts[0]
	pendingConflicts = pendingConflicts[1:]
	return &pop
}

func ReadPendingConflict() *Task {
	lock.Lock()
	defer lock.Unlock()
	if len(pendingConflicts) == 0 {
		return nil
	}

	read := pendingConflicts[0]
	return &read
}

func PushPendingConflict(t Task) {
	lock.Lock()
	pendingConflicts = append(pendingConflicts, t)
	lock.Unlock()
}

func ClearPendingConflicts() {
	lock.Lock()
	pendingConflicts = nil
	lock.Unlock()
}

// SolveAllConflicts marks all conflicts as solved,
// which means that all conflicts that have been found
// so far will be overwritten.
func SolveAllConflicts() {
	lock.Lock()
	solvedConflicts = append(solvedConflicts, pendingConflicts...)
	pendingConflicts = nil
	lock.Unlock()
}

func PopSolvedConflict() *Task {
	lock.Lock()
	defer lock.Unlock()

	if len(solvedConflicts) == 0 {
		return nil
	}

	pop := solvedConflicts[0]
	solvedConflicts = solvedConflicts[1:]
	return &pop
}

// PushSolvedConflict marks the conflict as solved
// which means that the already existing file will
// be overwritten.
func PushSolvedConflict(t Task) {
	lock.Lock()
	solvedConflicts = append(solvedConflicts, t)
	lock.Unlock()
}

func AddTask(p *Path) {
	progress.FullAmount++

	lock.Lock()
	sources = append(sources, *p)
	lock.Unlock()
}

func AddFolder(folder string) {
	if folder == "" {
		return
	}

	if !createFoldersInTarget && filepath.Dir(folder) == "." {
		// don't recreate root folders
		return
	}

	if !createFoldersInTarget {
		folder = removeFirst(folder)
	}

	folder = filepath.Join(targetBase, folder)

	addFolder(folder)
}

func addFolder(f string) {
	lock.Lock()
	folders = append(folders, f)
	lock.Unlock()
}

func destFor(p *Path) string {
	if createFoldersInTarget {
		return filepath.Join(targetBase, filepath.Base(p.Base), p.Sub)
	} else {
		return filepath.Join(targetBase, p.Sub)
	}
}

func Flen() int {
	lock.RLock()
	defer lock.RUnlock()
	return len(folders)
}

func PrintTasks() {
	lock.RLock()

	for _, source := range sources {
		fmt.Printf(" - %s will be copied to %s\n",
			filepath.Clean(source.AsAbs()),
			filepath.Clean(destFor(&source)))
	}

	fmt.Println("Need to create these folders:")
	for _, folder := range folders {
		fmt.Printf(" - %s\n", folder)
	}

	lock.RUnlock()
}
