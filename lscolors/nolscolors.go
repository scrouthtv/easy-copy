// +build nocolor nolscolors

package lscolors

import "os"

// FormatType returns an empty string if lscolors are disabled.
func FormatType(t string) string {
	return ""
}

// FormatFile returns an empty string if lscolors are disabled.
func FormatFile(info os.FileInfo) string {
	return ""
}

// ReloadLsColors does nothing if lscolors are disabled.
func ReloadLsColors() {
}
