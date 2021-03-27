// +build !linux,!windows goin

package input

import (
	"bufio"
	"os"
)

func Getch() rune {
	var rdr *bufio.Reader = bufio.NewReader(os.Stdin)
	r, _, _ := rdr.ReadRune()
	return r
}
