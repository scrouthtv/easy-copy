package main;

import "os";
import "os/exec";
import "bytes";
import "strconv";
import "strings";
import "path/filepath";
import "runtime";
import "bufio";
import "errors";
import "io";
import "unicode";

var noPagerError error = errors.New("No suitable pager found.");

/**
* Tries to find a pager in $PAGER or defaults to less or more.
* If none of those are available, runPager returns false and noPagerError.
*/
func runPager(text string) (bool, error) {
	var pager string;
	var ok bool;
	var err error;
	pager, ok = os.LookupEnv("PAGER")
	if !ok {
		_, err = exec.LookPath("less");
		if err == nil {
			pager = "less";
		} else {
			_, err = exec.LookPath("more");
			if err == nil {
				pager = "more";
			} else {
				return false, noPagerError;
			}
		}
	}
	var cmd *exec.Cmd;
	cmd = exec.Command(pager)
	var out io.WriteCloser;
	out, err = cmd.StdinPipe()
	if err != nil { return false, err; }
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Start();
	if err != nil { return false, err; }
	writer := bufio.NewWriter(out);
	writer.WriteString(text);
	writer.Flush();
	out.Close();
	cmd.Wait();
	return true, nil;
}

func getChoice(choices string) rune {
	var in rune;
	for {
		in = unicode.ToLower(getch());
		if strings.ContainsRune(choices, in) {
			return in;
		}
	}
}

func LinuxIsPiped() bool {
	fi, _ := os.Stdout.Stat();

	return (fi.Mode() & os.ModeCharDevice) == 0;
}

func WindowsParentProcessName() (string, error) {
	cmd := exec.Command("tasklist");//, "/fi \"pid eq " + strconv.Itoa(ppid) + "\" /nh", "") does not work
	var out bytes.Buffer;
	cmd.Stdout = &out;
	err := cmd.Run();
	if err != nil {
		return "", err;
	} else {
		var tasklist []string = strings.Split(out.String(), "\n");
		var header string = tasklist[2];
		var colpidstart int = strings.Index(header, " ");
		var colpidend int = colpidstart + 1 + strings.Index(header[colpidstart + 1:len(header)], " ");

		var paddedPID string = strconv.Itoa(os.Getppid());
		for len(paddedPID) < colpidend - colpidstart {
			paddedPID = " " + paddedPID;
		}

		var i int;
		for i = 3; i < len(tasklist) - 1; i++ {
			var row string = tasklist[i];
			if len(row) >= colpidend {
				if row[colpidstart:colpidend] == paddedPID {
					// process names are technically paths so we can use this to strip trailing whitespaces:
					return filepath.Clean(row[0:colpidstart - 1]), nil;
				}
			}
		}
	}
	return "", nil;
}


var autoColorsCache bool;
var autoColorsCacheSet bool = false;
// save the evaluation of autoColors() to avoid this tedious calculation

func autoColors() bool {
	if autoColorsCacheSet {
		return autoColorsCache;
	}
	if (runtime.GOOS == "windows") {
		var ppname string;
		var err error;
		ppname, err = WindowsParentProcessName();
		if err != nil || (ppname != "pwsh.exe" && ppname != "powershell.exe") {
			autoColorsCacheSet = true;
			autoColorsCache = false;
			return false;
		}
	} else if (runtime.GOOS == "linux" || runtime.GOOS == "darwin") {
		if LinuxIsPiped() {
			autoColorsCacheSet = true;
			autoColorsCache = false;
			return false;
		}
	}
	autoColorsCacheSet = true;
	autoColorsCache = true;
	return true;
}
