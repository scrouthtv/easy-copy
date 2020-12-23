// +build !linux,!windows
// +build rawin

package input

func Getch() rune {
	panic("NOT SUPPORTED")
}
