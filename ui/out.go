package ui

import (
	"easy-copy/color"
	"fmt"
	"os"
)

type Info interface {
	Info() string
}

var Infos = make(chan Info, 8)
var Warns = make(chan error, 8)

func Error(err error) {
	fmt.Println(color.FGColors.Red + "error: " + err.Error() + color.Text.Reset)
	PrintUsage()
	os.Exit(2)
}
