package main

import (
	"bufio"
	"easy-copy/input"
	"errors"
	"io"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"unicode"
)

var errNoPager error = errors.New("error: no suitable pager found")

type errOpeningPager struct {
	err error
}

func (err *errOpeningPager) Unwrap() error {
	return err.err
}

func (err *errOpeningPager) Error() string {
	return "error opening pager: " + err.err.Error()
}

/**
* Tries to find a pager in $PAGER or defaults to less or more.
* If none of those are available, runPager returns false and noPagerError.
 */
func runPager(text string) (bool, error) {
	var pager string
	var ok bool
	var err error
	pager, ok = os.LookupEnv("PAGER")
	if !ok {
		_, err = exec.LookPath("less")
		if err == nil {
			pager = "less"
		} else {
			_, err = exec.LookPath("more")
			if err == nil {
				pager = "more"
			} else {
				return false, &errOpeningPager{err}
			}
		}
	}
	cmd := exec.Command(pager)
	var out io.WriteCloser
	out, err = cmd.StdinPipe()
	if err != nil {
		return false, &errOpeningPager{err}
	}
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Start()
	if err != nil {
		return false, &errOpeningPager{err}
	}
	writer := bufio.NewWriter(out)
	writer.WriteString(text)
	writer.Flush()
	out.Close()
	cmd.Wait()

	return true, nil
}

func getChoice(choices string) rune {
	var in rune

	for {
		in = unicode.ToLower(input.Getch())
		if strings.ContainsRune(choices, in) {
			return in
		}
	}
}

func formatSeconds(seconds float64) string {
	if seconds < 100 {
		return strconv.FormatFloat(seconds, 'f', 0, 32) + " s"
	} else if seconds < 60*100 {
		return strconv.FormatFloat(seconds/60, 'f', 1, 32) + " m"
	} else if seconds < 60*60*72 {
		return strconv.FormatFloat(seconds/60/60, 'f', 2, 32) + " h"
	} else {
		return strconv.FormatFloat(seconds/60/60/24, 'f', 2, 32) + " days"
	}
}

/**
 * 0: best print in bytes
 * 1: best print in kb
 * 2: best print in mb
 * 3: best print in gb
 * 4: best print in tb
 */
func sizeAutoUnit(size float64) int {
	if size < 1000 {
		return 0
	} else if size < 1000*1000 {
		return 1
	} else if size < 1000*1000*1000 {
		return 2
	} else if size < 1000*1000*1000*1000 {
		return 3
	} else {
		return 4
	}
}

func formatSize(size float64, unit int) string {
	switch unit {
	case 0:
		return strconv.FormatFloat(size, 'f', 0, 32) + " b"
	case 1:
		return strconv.FormatFloat(size/1000, 'f', 0, 32) + " kB"
	case 2:
		return strconv.FormatFloat(size/1000/1000, 'f', 1, 32) + " MB"
	case 3:
		return strconv.FormatFloat(size/1000/1000/1000, 'f', 2, 32) + " GB"
	default:
		return strconv.FormatFloat(size/1000/1000/1000/1000, 'f', 2, 32) + " TB"
	}
}

func shrinkPath(path string, length int) string {
	if len(path) > length {
		return path[0:length-8] + "..." + path[len(path)-5:]
	} else {
		return path
	}
}
