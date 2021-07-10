package impl

import (
	"easy-copy/ui"
	"os"
	"strings"
)

type ErrEmptySource struct{}

func (e *ErrEmptySource) Error() string {
	return "empty source"
}

func (s *settingsImpl) searchStopFlag() {
	if len(os.Args) == 1 {
		return
	}

	for _, arg := range os.Args[1:] {
		if arg == "--" {
			return
		} else if strings.HasPrefix(arg, "--") {
			s.isStopFlag(arg[2:])
		} else if strings.HasPrefix(arg, "-") {
			s.isStopFlag(arg[1:])
		}
	}
}

func (s *settingsImpl) isStopFlag(arg string) {
	switch arg {
	case "h", "help":
		ui.PrintHelp()
		os.Exit(0)
	case "v", "version":
		ui.PrintVersion()
		os.Exit(0)
	case "copying":
		ui.ShowCopying()
		os.Exit(0)
	case "warranty":
		ui.ShowWarranty()
		os.Exit(0)
	case "colortest":
		ui.ShowColortest()
		os.Exit(0)
	}
}

func (s *settingsImpl) ParseLine() {
	s.searchStopFlag()

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
		ui.Error(&ErrEmptySource{})
	}

	s.target = s.sources[len(s.sources)-1]
	s.sources = s.sources[:len(s.sources)-1]
}
