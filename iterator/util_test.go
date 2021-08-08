package iterator

import (
	"easy-copy/tasks"
	"testing"
)

func cmpTasks(t *testing.T, ts []tasks.Task, fs []string) {
	t.Helper()

	task := tasks.PopTask()
	for task != nil {
		if ts == nil {
			t.Error("Unexpected extra task: ", *task)
			return
		}

		var ok bool
		ts, ok = removeTask(ts, task)
		if !ok {
			t.Error("Couldn't find this task: ", *task)
		}

		task = tasks.PopTask()
	}

	if len(ts) != 0 {
		t.Error("Missing tasks: ", ts)
	}
}

func removeTask(ts []tasks.Task, t *tasks.Task) ([]tasks.Task, bool) {
	if len(ts) == 0 {
		return nil, false
	}

	var idx int = -1
	for i, task := range ts {
		if task.Source == t.Source && task.Dest == t.Dest {
			idx = i
			break
		}
	}

	if idx == -1 {
		return ts, false
	}

	if idx == 0 {
		return ts[1:], true
	} else if idx == len(ts)-1 {
		return ts[:idx], true
	} else {
		return append(ts[:idx], ts[idx+1:]...), true
	}
}
