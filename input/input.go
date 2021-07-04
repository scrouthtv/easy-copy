package input

import (
	"os"
	"strings"
	"unicode"
)

const errReturnCode = 8

var buf []byte = make([]byte, 8)

// GetChoice takes a string of runes
// and lets the user choose one of them.
func GetChoice(choices string) rune {
	var producer func() rune = Getch

	err := Enter()
	if err != nil {
		producer = goin
	} else {
		defer Exit()
	}

	var in rune

	for {
		in = unicode.ToLower(producer())
		if strings.ContainsRune(choices, in) {
			return in
		}
	}
}

func goin() rune {
	os.Stdout.Write([]byte{' ', '>', ' '})

	_, err := os.Stdin.Read(buf)
	if err != nil {
		os.Exit(errReturnCode)
	}

	return rune(buf[0])
}
