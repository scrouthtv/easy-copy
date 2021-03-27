// +build !nocolor

package color

import "runtime"

// save the evaluation of autoColors() to avoid this tedious calculation
var (
	autoColorsCache    bool
	autoColorsCacheSet bool = false
)

func AutoColors() bool {
	if autoColorsCacheSet {
		return autoColorsCache
	}
	if runtime.GOOS == "windows" {
		var ppname string
		var err error
		ppname, err = WindowsParentProcessName()
		if err != nil || (ppname != "pwsh.exe" && ppname != "powershell.exe") {
			autoColorsCacheSet = true
			autoColorsCache = false
			return false
		}
	} else if runtime.GOOS == "linux" || runtime.GOOS == "darwin" {
		if LinuxIsPiped() {
			autoColorsCacheSet = true
			autoColorsCache = false
			return false
		}
	}
	autoColorsCacheSet = true
	autoColorsCache = true
	return true
}
