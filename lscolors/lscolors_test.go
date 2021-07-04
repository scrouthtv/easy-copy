// +build !nolscolors

package lscolors

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
	"testing"
)

func TestLoadLSC(t *testing.T) {
	ReloadLsColors()

	for k, v := range lsc.types {
		t.Logf("%s => %s\n", k, v)
	}

	t.Logf("%d extensions", len(lsc.exts))
}

func TestFormatSingleFile(t *testing.T) {
	f, err := ioutil.TempFile("", "*.txt")
	if err != nil {
		t.Fatal(err)
	}

	info, err := os.Lstat(f.Name())
	if err != nil {
		t.Fatal(err)
	}

	format := FormatFile(info)

	t.Logf("formatting with %s:\n", format)
	t.Logf("\033["+format+"m%s\033[0m", f.Name())

	f.Close()
	os.Remove(f.Name())
}

func TestFormatSingleFolder(t *testing.T) {
	home, err := os.UserHomeDir()
	if err != nil {
		t.Fatal(err)
	}

	info, err := os.Lstat(home)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("%t", info.IsDir())
	t.Logf(FormatFile(info))
}

// Simulate an ls output on ~
// Compare to ls -A -w 81 -x --color=auto ~.
func TestFormatHomeFolder(t *testing.T) {
	home, err := os.UserHomeDir()
	if err != nil {
		t.Fatal(err)
	}

	hf, err := os.Open(home)
	if err != nil {
		t.Fatal(err)
	}

	names, err := hf.Readdirnames(0)
	if err != nil {
		t.Fatal(err)
	}

	var longest int = 0
	for _, name := range names {
		if len(name) > longest {
			longest = len(name)
		}
	}

	sort.Slice(names, func(i int, j int) bool {
		return strings.Compare(
			strings.ToLower(strings.TrimLeft(names[i], ".")),
			strings.ToLower(strings.TrimLeft(names[j], ".")),
		) == -1
	})

	// entries per line
	var epl int = 81 / longest

	var rows int = len(names) / epl
	var row, col int
	var path string
	var info os.FileInfo
	var line string

	for row = 0; row < rows; row++ {
		line = ""

		for col = 0; col < epl; col++ {
			path = filepath.Join(home, names[row*epl+col])
			info, err = os.Lstat(path)

			if err != nil {
				t.Fatal(err)
			}

			line += fmt.Sprintf("\033[%sm%-"+strconv.Itoa(longest+1)+"s\033[0m",
				FormatFile(info), names[row*epl+col])
		}
		fmt.Println(line)
	}
}
