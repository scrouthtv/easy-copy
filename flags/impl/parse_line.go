package impl

import (
	"easy-copy/ui/msg"
	"os"
	"strings"
)

func (s *settingsImpl) ParseLine() {
	s.parseMode()

	isFiles := false

	for _, arg := range os.Args[2:] {
		if isFiles {
			s.sources = append(s.sources, arg)
		} else if arg == "--" {
			isFiles = true
		} else if strings.HasPrefix(arg, "--") {
			s.parseFlag("--", arg[2:])
		} else if strings.HasPrefix(arg, "-") {
			for i := 1; i < len(arg); i++ {
				s.parseFlag("-", arg[i:i+1])
			}
		} else {
			s.sources = append(s.sources, arg)
		}
	}

	if len(s.sources) < 2 {
		msg.ErrEmptySource()
	}

	s.target = s.sources[len(s.sources)-1]
	s.sources = s.sources[:len(s.sources)-1]
}
