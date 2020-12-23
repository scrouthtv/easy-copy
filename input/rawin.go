// +build !linux,!windows
// +build !goin

package input

func Getch() rune {
	panic("NOT SUPPORTED")
}
