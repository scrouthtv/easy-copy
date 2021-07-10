package impl

import (
	"easy-copy/color"
	"easy-copy/files"
	"easy-copy/flags"
	"easy-copy/ui"
	"strconv"
	"strings"
)

type InfoVerboseEnabled struct{}

func (i *InfoVerboseEnabled) Info() string {
	return "verbose mode enabled"
}

type InfoDryrun struct{}

func (i *InfoDryrun) Info() string {
	return "dry run mode enabled"
}

type ErrUnknownOption struct {
	option string
}

func (e *ErrUnknownOption) Error() string {
	return "unknown option: " + e.option
}

type ErrBadConfigKey struct {
	Key string
}

func (e *ErrBadConfigKey) Error() string {
	return "invalid config key '" + e.Key + "'"
}

func (s *settingsImpl) parseFlag(prefix string, flag string) {
	if strings.ContainsRune(flag, '=') {
		s.parseOption(flag)
		return
	}

	switch flag {
	case "debug":
		s.verbosity = flags.VerbDebug

		ui.Infos <- &InfoVerboseEnabled{}
	case "V", "verbose":
		s.verbosity = flags.VerbInfo

		ui.Infos <- &InfoVerboseEnabled{}
	case "q", "quiet":
		s.verbosity = flags.VerbQuiet
	case "e", "extended-colors":
		s.doLScolors = true
	case "n", "no-clobber", "no-overwrite":
		s.onConflict = flags.ConflictSkip
	case "f", "force", "overwrite":
		s.onConflict = flags.ConflictOverwrite
	case "i", "interactive":
		s.onConflict = flags.ConflictAsk
	case "color":
		color.Init(true)
	case "dry":
		s.dryrun = true

		ui.Infos <- &InfoDryrun{}
	default:
		ui.Warns <- &ErrUnknownOption{prefix + flag}
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
			ui.Warns <- &ErrBadValue{key: "buffersize", value: value}
		}
	default:
		ui.Warns <- &ErrBadConfigKey{key}
	}
}
