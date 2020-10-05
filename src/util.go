package main;

import "os";
import "os/exec";
import "bytes";
import "strconv";
import "strings";
import "path/filepath";

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
