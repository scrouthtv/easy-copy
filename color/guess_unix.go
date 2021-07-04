// +build linux freebsd openbsd netbsd dragonfly darwin
// +build !nocolor

package color

import (
	"os"
)

var (
	autoColorsCacheSet = false
	autoColorsCache    = false
)

func isPiped() bool {
	fi, _ := os.Stdout.Stat()

	return (fi.Mode() & os.ModeCharDevice) == 0
}

// AutoColors determines whether colors should be enabled
// by checking if the output is piped into a file.
func AutoColors() bool {
	if autoColorsCacheSet {
		return autoColorsCache
	}

	if isPiped() {
		autoColorsCacheSet = true
		autoColorsCache = false

		return false
	}

	autoColorsCacheSet = true
	autoColorsCache = true

	return true
}
