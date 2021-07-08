package impl

import (
	"easy-copy/flags"
	"os"
)

func (s *settingsImpl) parseMode() {
	switch os.Args[1] {
	case "copy":
		s.mode = flags.ModeCopy
	case "move":
		s.mode = flags.ModeMove
	default:
		panic("mode not supported")
	}
}
