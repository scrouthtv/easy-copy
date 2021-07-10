package tasks

import (
	"easy-copy/color"
	"runtime/debug"
	"sync"
)

type LoggedLock struct {
	lock  sync.RWMutex
	dolog bool
}

func newLock() *LoggedLock {
	return &LoggedLock{dolog: false}
}

func (l *LoggedLock) Lock() {
	l.lock.Lock()
	if l.dolog {
		println(color.BGColors.Green + "Lock: Locked" + color.Text.Reset)
		debug.PrintStack()
	}
}

func (l *LoggedLock) Unlock() {
	l.lock.Unlock()
	if l.dolog {
		println(color.BGColors.Green + "Lock: Unlocked" + color.Text.Reset)
		debug.PrintStack()
	}
}

func (l *LoggedLock) RLock() {
	l.lock.RLock()
	if l.dolog {
		println(color.BGColors.Green + "Lock: RLock" + color.Text.Reset)
		debug.PrintStack()
	}
}

func (l *LoggedLock) RUnlock() {
	l.lock.RUnlock()
	if l.dolog {
		println(color.BGColors.Green + "Lock: RUnlock" + color.Text.Reset)
		debug.PrintStack()
	}
}
