package impl

import (
	"easy-copy/flags"
	"easy-copy/ui"
	"os"
)

type ErrMissingOperation struct{}

func (e *ErrMissingOperation) Error() string {
	return "missing operation"
}

type ErrUnsupportedOperation struct {
	Op string
}

func (e *ErrUnsupportedOperation) Error() string {
	return "unsupported operation: " + e.Op
}

func (s *settingsImpl) parseMode() {
	if len(os.Args) < 2 {
		ui.Error(&ErrMissingOperation{})
		return
	}

	switch os.Args[1] {
	case "cp":
		s.mode = flags.ModeCopy
	case "mv":
		s.mode = flags.ModeMove
	default:
		ui.Error(&ErrUnsupportedOperation{Op: os.Args[1]})
	}
}
