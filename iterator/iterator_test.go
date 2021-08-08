package iterator

import (
	"easy-copy/flags"
	"easy-copy/flags/stub"
	"easy-copy/fs"
	"easy-copy/tasks"
	"flag"
	"log"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	flag.Parse()

	fs := fs.CreateFS([]string{
		"foo/a",
		"foo/b",
		"foo/sub/q",
		"quz/c",
		"quz/d",
		"quz/empty/",
		"bar.txt",
	})
	opener = fs

	cfg := stub.New()
	flags.Current = cfg

	if !testing.Verbose() {
		cfg.SetVerbosity(flags.VerbCrit)
	}

	exit := m.Run()

	if exit != 0 || testing.Verbose() {
		tree := fs.Tree()
		for _, l := range tree {
			log.Println(l)
		}
	}

	os.Exit(exit)
}

func TestIterateFolder(t *testing.T) {
	tasks.Setup("/baz", false)

	err := add(&tasks.Path{Base: "/foo", Sub: ""})
	if err != nil {
		t.Error(err)
	}

	if testing.Verbose() {
		tasks.PrintTasks()
	}

	shouldTasks := []tasks.Task{
		{Source: "/foo/sub/q", Dest: "/baz/sub/q"},
		{Source: "/foo/a", Dest: "/baz/a"},
		{Source: "/foo/b", Dest: "/baz/b"},
	}

	shouldFolders := []string{"/baz", "/baz/sub"}

	cmpTasks(t, shouldTasks, shouldFolders)
}

func TestIterateMultiFolders(t *testing.T) {
	tasks.Setup("/baz", true)

	err := add(&tasks.Path{Base: "/foo", Sub: ""})
	if err != nil {
		t.Error(err)
	}

	err = add(&tasks.Path{Base: "/quz", Sub: ""})
	if err != nil {
		t.Error(err)
	}

	tasks.PrintTasks()
}

func TestIterateFile(t *testing.T) {
	tasks.Setup("/baz", true)

	err := add(&tasks.Path{Base: "/bar.txt", Sub: ""})
	if err != nil {
		t.Error(err)
	}

	tasks.PrintTasks()
}
