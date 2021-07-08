package impl

import (
	"easy-copy/color"
	"easy-copy/files"
	"easy-copy/flags"
	"easy-copy/ui"
	"easy-copy/ui/msg"
	"os"
	"strconv"
	"strings"
)

func (s *settingsImpl) parseFlag(prefix string, flag string) {
	if strings.ContainsRune(flag, '=') {
		s.parseOption(flag)
		return
	}

	switch flag {
	case "h", "help":
		ui.PrintHelp()
		os.Exit(0)
	case "v", "version":
		ui.PrintVersion()
		os.Exit(0)
	case "debug":
		s.verbosity = flags.VerbDebug

		msg.VerbVerboseEnabled()
	case "V", "verbose":
		s.verbosity = flags.VerbInfo

		msg.VerbVerboseEnabled()
	case "q", "quiet":
		s.verbosity = flags.VerbQuiet
	case "copying":
		ui.ShowCopying()
		os.Exit(0)
	case "warranty":
		ui.ShowWarranty()
		os.Exit(0)
	case "e", "extended-colors":
		lscolors = true
	case "colortest":
		ui.ShowColortest()
		os.Exit(0)
	case "n", "no-clobber", "no-overwrite":
		s.onConflict = flags.ConflictSkip
	case "f", "force", "overwrite":
		conflict = flags.ConflictOverwrite
	case "i", "interactive":
		conflict = flags.ConflictAsk
	case "color":
		color.Init(true)
	case "dry":
		s.dryrun = true

		msg.VerbDryrun()
	default:
		msg.ErrUnknownOption(prefix + flag)
	}
}

func (s *settingsImpl) parseKeyValue(key string, value string) {
	// Trim away spaces and tabs:
	key = strings.ToLower(strings.Trim(key, " \t'\""))
	value = strings.ToLower(strings.Trim(value, " \t'\""))

	switch key {
	case "verbose":
		if configInterpretBoolean(value) {
			s.verbosity = flags.VerbInfo
		}
	case "quiet":
		if configInterpretBoolean(value) {
			s.verbosity = flags.VerbQuiet
		}
	case "overwrite":
		switch value {
		case "skip":
			conflict = flags.ConflictSkip
		case "overwrite":
			conflict = flags.ConflictOverwrite
		case "ask":
			conflict = flags.ConflictAsk
		}
	case "symlinks":
		switch value {
		case "ignore":
			s.onSymlink = flags.SymlinkIgnore
		case "link":
			s.onSymlink = flags.SymlinkLink
		case "dereference":
			s.onSymlink = flags.SymlinkDeref
		}
	case "extended-colors":
		if configInterpretBoolean(value) {
			s.doLScolors = true
		}
	case "color":
		switch configInterpretAutoOrBoolean(value) {
		case 0:
			color.Init(false)
		case 1:
			color.Init(color.AutoColors())
		case 2:
			color.Init(true)
		}
	case "buffersize":
		val, err := strconv.Atoi(value)
		if err == nil {
			files.SetBuffersize(val)
		} else {
			msg.WarnConfig(&ErrBadValue{key: "buffersize", value: value})
		}
	default:
		msg.WarnBadConfigKey(key)
	}
}
