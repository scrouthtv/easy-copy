package ui

import (
	"bufio"
	"errors"
	"os"
	"os/exec"
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
