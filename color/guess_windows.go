//go:build !nocolor && windows

package color

import (
	"bytes"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
)

// save the evaluation of autoColors() to avoid re-guessing.
var (
	autoColorsCache    bool
	autoColorsCacheSet bool = false
)

// AutoColors determines whether colors should be enabled
// by checking whether we are running in powershell.
func AutoColors() bool {
	if autoColorsCacheSet {
		return autoColorsCache
	}

	ppname, err := parentProcessName()
	if err != nil {
		autoColorsCacheSet = true
		autoColorsCache = false

		return false
	}

	if ppname == "pwsh.exe" || ppname == "powershell.exe" {
		autoColorsCacheSet = true
		autoColorsCache = false

		return true
	}

	autoColorsCacheSet = true
	autoColorsCache = false

	return false
}

func parentProcessName() (string, error) {
	cmd := exec.Command("tasklist")
	var out bytes.Buffer
	cmd.Stdout = &out

	err := cmd.Run()
	if err != nil {
		return "", err
	}

	tasklist := strings.Split(out.String(), "\n")
	header := tasklist[2]
	colpidstart := strings.Index(header, " ")
	colpidend := colpidstart + 1 + strings.Index(header[colpidstart+1:], " ")

	var paddedPID string = strconv.Itoa(os.Getppid())
	for len(paddedPID) < colpidend-colpidstart {
		paddedPID = " " + paddedPID
	}

	for i := 3; i < len(tasklist)-1; i++ {
		var row string = tasklist[i]
		if len(row) >= colpidend {
			if row[colpidstart:colpidend] == paddedPID {
				// process names are technically paths so we can use this to strip trailing whitespaces:
				return filepath.Clean(row[0 : colpidstart-1]), nil
			}
		}
	}

	return "", nil
}
