package tasks

import (
	"fmt"
	"path/filepath"
)

var targetBase string
var createFoldersInTarget bool

var lock *LoggedLock = newLock()
var sources []Path
var pendingConflicts, solvedConflicts []Task
var folders []string

type Task struct {
	Source string
	Dest   string
}

// Setup is called once for setting the target base
// and whether to recreate the root folders inside the target.
func Setup(base string, cloneFolders bool) {
	targetBase = base
	createFoldersInTarget = cloneFolders

	AddFolder(base)
}

// PopTask pops the next available task from the task list.
func PopTask() *Task {
	lock.Lock()
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
	lock.Lock()
	sources = append(sources, *p)
	lock.Unlock()
}

func AddFolder(folder string) {
	if folder == "" {
		return
	}

	lock.Lock()
	folders = append(folders, folder)
	lock.Unlock()
}

func destForFolder(f string) string {
	if createFoldersInTarget {
		return filepath.Join(targetBase, f)
	} else {
		return filepath.Join(targetBase, removeFirst(f))
	}
}

func destFor(p *Path) string {
	if createFoldersInTarget {
		return filepath.Join(targetBase, p.Sub)
	} else {
		return filepath.Join(targetBase, removeFirst(p.Sub))
	}
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
		fmt.Printf(" - %s\n", destForFolder(folder))
	}

	lock.RUnlock()
}
