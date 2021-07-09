package tasks

import (
	"fmt"
	"path/filepath"
	"sync"
)

var targetBase string
var createFoldersInTarget bool

var lock *sync.RWMutex = &sync.RWMutex{}
var sources []Path
var pendingConflicts, solvedConflicts []Path
var folders []string

type Task struct {
	Source string
	Dest   string
}

var Done = false

// Setup is called once for setting the target base
// and whether to recreate the root folders inside the target.
func Setup(base string, cloneFolders bool) {
	targetBase = base
	createFoldersInTarget = cloneFolders
}

// PopTask pops the next available task from the task list.
func PopTask() Task {
	lock.Lock()
	next := sources[0]
	sources = sources[1:]
	lock.Unlock()

	return Task{Source: next.AsAbs(), Dest: destFor(&next)}
}

func PopPendingConflict() Task {
	// TODO impl
	return Task{}
}

func PopSolvedConflict() *Task {
	lock.Lock()
	pop := solvedConflicts[0]
	solvedConflicts = solvedConflicts[1:]
	lock.Unlock()

	return &Task{Source: pop.AsAbs(), Dest: destFor(&pop)}
}

func AddTask(p *Path) {
	lock.Lock()
	sources = append(sources, *p)
	lock.Unlock()
}

func AddFolder(folder string) {
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
