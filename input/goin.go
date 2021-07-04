// +build goin !windows,!linux,!freebsd,!openbsd,!netbsd,!dragonfly,!darwin

package input

// Enter does nothing if rawin is not used.
func Enter() error {
	return nil
}

// Exit does nothing if rawin is not used.
func Exit() {
}

// Getch uses the fallback goin implementation by default
// if rawin is not used.
func Getch() rune {
	return goin()
}
