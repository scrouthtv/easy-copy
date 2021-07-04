package main

import (
	"bufio"
	"errors"
	"os"
	"os/exec"
	"strconv"
)

const (
	bfactor  = 1024 // bytes in a kb
	mfactor  = 60   // minutes in an hour
	hfactor  = 24   // hours in a day
	nicesize = 300  // highest file size before choosing the bigger unit
	nicetime = 100  // highest time before choosing the bigge runit
)

var errNoPager = errors.New("no suitable pager found")

type errOpeningPager struct {
	err error
}

func (err *errOpeningPager) Unwrap() error {
	return err.err
}

func (err *errOpeningPager) Error() string {
	return "error opening pager: " + err.err.Error()
}

// runPager tries to find a suitable pager.
// If one's found, the specified text is displayed via it.
// If none could be opened, an error is returned.
func runPager(text string) error {
	pager, ok := os.LookupEnv("PAGER")
	if !ok {
		_, err := exec.LookPath("less")
		if err == nil {
			pager = "less"
		} else {
			_, err = exec.LookPath("more")
			if err == nil {
				pager = "more"
			} else {
				return &errOpeningPager{err}
			}
		}
	}

	cmd := exec.Command(pager)

	out, err := cmd.StdinPipe()
	if err != nil {
		return &errOpeningPager{err}
	}

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err = cmd.Start()
	if err != nil {
		return &errOpeningPager{err}
	}

	writer := bufio.NewWriter(out)

	_, err = writer.WriteString(text)
	if err != nil {
		out.Close()
		return &errOpeningPager{err}
	}

	writer.Flush()
	out.Close()

	err = cmd.Wait()
	if err != nil {
		return &errOpeningPager{err}
	}

	return nil
}

func formatSeconds(seconds float64) string {
	if seconds < nicetime {
		return strconv.FormatFloat(seconds, 'f', 0, 32) + " s"
	} else if seconds < nicetime*mfactor {
		return strconv.FormatFloat(seconds/mfactor, 'f', 1, 32) + " m"
	} else if seconds < nicetime*mfactor*mfactor {
		return strconv.FormatFloat(seconds/mfactor/mfactor, 'f', 2, 32) + " h"
	} else {
		return strconv.FormatFloat(seconds/mfactor/mfactor/hfactor, 'f', 2, 32) + " days"
	}
}

// sizeAutoUnit determines which unit this size in bytes
// prints the nicest in
// 0: best print in bytes
// 1: best print in kb
// 2: best print in mb
// 3: best print in gb
// 4: best print in tb
func sizeAutoUnit(size float64) int {
	if size < nicesize {
		return 0 //  B up to 300
	} else if size < nicesize*bfactor {
		return 1 // kB up to 300
	} else if size < nicesize*bfactor*bfactor {
		return 2 // MB up to 300
	} else if size < nicesize*bfactor*bfactor*bfactor {
		return 3 // GB up to 300
	} else {
		return 4 // no one copies more than 300 tb
	}
}

func formatSize(size float64, unit int) string {
	switch unit {
	case 0:
		return strconv.FormatFloat(size, 'f', 0, 32) + " b"
	case 1:
		return strconv.FormatFloat(size/bfactor, 'f', 1, 32) + " kB"
	case 2:
		return strconv.FormatFloat(size/bfactor/bfactor, 'f', 2, 32) + " MB"
	case 3:
		return strconv.FormatFloat(size/bfactor/bfactor/bfactor, 'f', 2, 32) + " GB"
	default:
		return strconv.FormatFloat(size/bfactor/bfactor/bfactor/bfactor, 'f', 2, 32) + " TB"
	}
}

func shrinkPath(path string, length int) string {
	if len(path) > length {
		return path[0:length-8] + "..." + path[len(path)-5:]
	} else {
		return path
	}
}
