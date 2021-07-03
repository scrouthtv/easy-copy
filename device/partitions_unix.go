// +build linux freebsd openbsd netbsd dragonfly darwin

package device

import "os"
import "bufio"
import "strings"
import "strconv"

import "golang.org/x/sys/unix"

type partTable map[uint64]partition

type partition struct {
	dev string /* device path */
	mnt string /* where the partition is mounted */
}

var currentPartTable partTable

func reloadDevices() {
	currentPartTable = make(partTable)

	f, err := os.Open("/proc/" + strconv.Itoa(os.Getpid()) + "/mountinfo")
	if err != nil {
		pushError(err)
		return
	}
	defer f.Close()

	s := bufio.NewScanner(f)

	for s.Scan() {
		reloadDevice(s.Text())
	}
}

func reloadDevice(line string) {
	split := strings.Split(line, " ")

	if len(split) > 10 {
		majmin := split[2]
		mnt := split[4]
		dev := split[sepDash(split) + 2]

		// fix spaces, maybe other characters are escaped too?
		mnt = strings.ReplaceAll(mnt, "\\040", " ")
		dev = strings.ReplaceAll(dev, "\\040", " ")

		c := strings.Index(majmin, ":")
		maj, err := strconv.ParseUint(majmin[:c], 10, 32)
		if err != nil {
			return
		}

		min, err := strconv.ParseUint(majmin[c+1:], 10, 32)
		if err != nil {
			return
		}

		id := unix.Mkdev(uint32(maj), uint32(min))
		currentPartTable[id] = partition{dev: dev, mnt: mnt}
	}
}

func sepDash(split []string) int {
	for i, s := range split {
		if s == "-" {
			return i
		}
	}

	return -1
}

func name(dev uint64) string {
	if (currentPartTable == nil) {
		reloadDevices()
	}

	part, ok := currentPartTable[dev]
	if ok {
		return part.mnt
	} else {
		return "<not found>"
	}
}
