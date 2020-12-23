// +build goin

package input

import "os"
import "bufio"

func Getch() rune {
	var rdr *bufio.Reader = bufio.NewReader(os.Stdin)
	r, _, _ := rdr.ReadRune()
	return r
}
