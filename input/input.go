package input

import "os"
import "unicode"
import "strings"

func GetChoice(choices string) rune {
	err := Enter()
	if err != nil {
		os.Exit(8)
	}
	defer Exit()

	var in rune

	for {
		in = unicode.ToLower(Getch())
		if strings.ContainsRune(choices, in) {
			return in
		}
	}
}

