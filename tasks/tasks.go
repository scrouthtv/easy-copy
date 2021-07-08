package tasks

import (
	"fmt"
	"sync"
)

var targetBase string

var lock *sync.RWMutex = &sync.RWMutex{}
var sources []Path

type Path struct {
	Base string
	Sub  string
}

type Task struct {
	Source string
	Dest   string
}

func SetBase(base string) {
	targetBase = base
}

func NextTask() Task {
	panic("not impl")
	return Task{}
}

func NextConflict() Task {
	panic("not impl")
	return Task{}
}

func AddTask(p Path) {
	lock.Lock()
	sources = append(sources, p)
	lock.Unlock()
}

func PrintTasks() {
	lock.RLock()

	for _, source := range sources {
		fmt.Printf(" - %s will be copied to %s/%s\n", source.Base, targetBase, source.Sub)
	}

	lock.RUnlock()
}
