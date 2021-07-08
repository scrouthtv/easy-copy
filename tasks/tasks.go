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
var conflicts []Path

type Task struct {
	Source string
	Dest   string
}

// Setup is called once for setting the target base
// and whether to recreate the root folders inside the target.
func Setup(base string, createFolders bool) {
	targetBase = base
	createFoldersInTarget = createFolders
}

// PopTask pops the next available task from the task list.
func PopTask() Task {
	lock.Lock()
	next := sources[0]
	sources = sources[1:]
	lock.Unlock()

	return Task{Source: next.AsAbs(), Dest: destFor(&next)}
}

func NextConflict() Task {
	// TODO impl
	return Task{}
}

func AddTask(p *Path) {
	lock.Lock()
	sources = append(sources, *p)
	lock.Unlock()
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

	lock.RUnlock()
}
