package impl

import (
	"easy-copy/flags"
	"easy-copy/ui/msg"
	"os"
)

func (s *settingsImpl) parseMode() {
	if len(os.Args) < 2 {
		msg.ErrMissingOperation()
	}

	switch os.Args[1] {
	case "copy":
		s.mode = flags.ModeCopy
	case "move":
		s.mode = flags.ModeMove
	default:
		panic("mode not supported")
	}
}
