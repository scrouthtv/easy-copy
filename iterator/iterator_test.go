package iterator

import (
	"easy-copy/common"
	"easy-copy/flags"
	"easy-copy/flags/stub"
	"easy-copy/fs"
	"easy-copy/tasks"
	"easy-copy/ui"
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
	common.FileAdapter = fs

	cfg := stub.New()
	flags.Current = cfg

	if !testing.Verbose() {
		cfg.SetVerbosity(flags.VerbCrit)
	}

	errExit := 0

	go func() {
		for {
			err := <-ui.Warns
			errExit = 125
			log.Println("Got an error:", err)
		}
	}()

	exit := m.Run()

	if exit|errExit != 0 || testing.Verbose() {
		tree := fs.Tree()
		for _, l := range tree {
			log.Println(l)
		}
	}

	os.Exit(exit | errExit)
}

func TestIterateFolder(t *testing.T) {
	common.FileAdapter.(*fs.MockFS).Rewind()

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
	common.FileAdapter.(*fs.MockFS).Rewind()

	tasks.Setup("/baz", true)

	err := add(&tasks.Path{Base: "/foo", Sub: ""})
	if err != nil {
		t.Error(err)
	}

	err = add(&tasks.Path{Base: "/quz", Sub: ""})
	if err != nil {
		t.Error(err)
	}

	if testing.Verbose() {
		for _, l := range common.FileAdapter.(*fs.MockFS).Tree() {
			log.Println(l)
		}

		tasks.PrintTasks()
	}

	shouldTasks := []tasks.Task{
		{Source: "/foo/sub/q", Dest: "/baz/foo/sub/q"},
		{Source: "/foo/a", Dest: "/baz/foo/a"},
		{Source: "/foo/b", Dest: "/baz/foo/b"},
		{Source: "/quz/c", Dest: "/baz/quz/c"},
		{Source: "/quz/d", Dest: "/baz/quz/d"},
	}

	shouldFolders := []string{
		"/baz",
		"/baz/foo", "/baz/foo/sub",
		"/baz/quz", "/baz/quz/empty",
	}

	cmpTasks(t, shouldTasks, shouldFolders)
}

func TestIterateFile(t *testing.T) {
	tasks.Setup("/baz", true)

	err := add(&tasks.Path{Base: "/bar.txt", Sub: ""})
	if err != nil {
		t.Error(err)
	}

	if testing.Verbose() {
		tasks.PrintTasks()
	}

	shouldTasks := []tasks.Task{
		{Source: "/bar.txt", Dest: "/baz/bar.txt"},
	}

	shouldFolders := []string{"/baz"}

	cmpTasks(t, shouldTasks, shouldFolders)
}

func TestIterateEmptyFolder(t *testing.T) {
	common.FileAdapter.(*fs.MockFS).Rewind()

	tasks.Setup("/baz", true)

	err := add(&tasks.Path{Base: "/quz/empty", Sub: ""})
	if err != nil {
		t.Error(err)
	}

	if testing.Verbose() {
		tasks.PrintTasks()
	}

	shouldTasks := []tasks.Task{}

	shouldFolders := []string{"/baz", "/baz/empty"}

	cmpTasks(t, shouldTasks, shouldFolders)
}
