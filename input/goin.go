// +build !windows,!linux,!freebsd,!openbsd,!netbsd,!dragonfly,!darwin
// +build !goin

package input

import "bufio"

func Getch() rune {
	rdr := bufio.NewReader(os.Stdin)
}
