package full_test

import (
	"easy-copy/common"
	"easy-copy/flags"
	"easy-copy/flags/impl"
	"easy-copy/fs"
	"easy-copy/iterator"
	"easy-copy/progress"
	"easy-copy/tasks"
	"easy-copy/ui"
	"testing"
)

type TestSetup struct {
	line   string
	is     *fs.MockFS
	should *fs.MockFS
}

func NewTest(line string) *TestSetup {
	return &TestSetup{line: line}
}

// Run runs the software with the arguments in t.line.
// It fails if a warning occurs.
func (test *TestSetup) Run(t *testing.T) {
	t.Helper()

	go func() {
		for !progress.CopyDone {
			select {
			case w := <-ui.Warns:
				t.Error(w)
			case i := <-ui.Infos:
				t.Log(i.Info())
			}
		}
	}()

	common.FileAdapter = test.is

	flags.Current = impl.New()

	flags.Current.ParseLine("ec " + test.line)

	iterator.Iterate()
	tasks.CopyLoop()
}
