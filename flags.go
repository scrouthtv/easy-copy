package main

import (
	"easy-copy/color"
	"easy-copy/config"
	"os"
	"strconv"
	"strings"
)

// ErrBadValue is used when an attempt to set a configuration value
// was unsuccessful.
type ErrBadValue struct {
	key   string
	value string
}

func (e *ErrBadValue) Error() string {
	return "bad value for " + e.key + ": " + e.value
}

// ErrBadConfigLine is used when there's a config line
// with too little or too many '='.
type ErrBadConfigLine struct {
	line string
}

func (e *ErrBadConfigLine) Error() string {
	return "couldn't parse this line: " + e.line
}

var verbose int = VerbNotice

const (
	// VerbQuiet indicates that no output should be written at all.
	VerbQuiet = iota

	// VerbCrit indicates that only critical messages should be written.
	VerbCrit

	// VerbNotice indicates that critical and helpful messages should be written.
	VerbNotice

	// VerbInfo indicates that additional info should be written.
	VerbInfo

	// VerbDebug should only be used for debugging.
	VerbDebug
)

var onExistingFile uint8 = Ask

const (
	// Skip existing files.
	Skip uint8 = iota

	// Overwrite existing files.
	Overwrite

	// Ask the user on existing files.
	Ask
)

// followSymlinks takes one of these values:
// 0 ignore symlinks
// 1 follow symlinks, copying them as links
// 2 fully dereference.
var followSymlinks uint8 = 1

var progressLSColors bool = false

var dryrun bool = false

// readConfig checks if the config file should be read
// (e.g. option no-config is not present),
// and reads it if we want to.
func readConfig() {
	for _, arg := range os.Args {
		if arg == "--no-config" {
			return
		} else if arg == "--" {
			break
		}
	}

	kvs, err := config.Load()
	if err != nil {
		warnConfig(err)
	}

	for _, line := range kvs {
		parseOption(line)
	}
}

func parseKeyValue(key string, value string) {
	// Trim away spaces and tabs:
	key = strings.ToLower(strings.Trim(key, " \t'\""))
	value = strings.ToLower(strings.Trim(value, " \t'\""))

	switch key {
	case "verbose":
		if configInterpretBoolean(value) {
			verbose = VerbInfo
		}
	case "quiet":
		if configInterpretBoolean(value) {
			verbose = VerbQuiet
		}
	case "overwrite":
		switch value {
		case "skip":
			onExistingFile = Skip
		case "overwrite":
			onExistingFile = Overwrite
		case "ask":
			onExistingFile = Ask
		}
	case "symlinks":
		switch value {
		case "ignore":
			followSymlinks = 0
		case "link":
			followSymlinks = 1
		case "dereference":
			followSymlinks = 2
		}
	case "extended-colors":
		if configInterpretBoolean(value) {
			progressLSColors = true
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
			setBuffersize(val)
		} else {
			warnConfig(&ErrBadValue{key: "buffersize", value: value})
		}
	default:
		warnBadConfigKey(key)
	}
}

func parseOption(line string) {
	kv := strings.Split(line, "=")
	if len(kv) != 2 {
		warnConfig(&ErrBadConfigLine{line: line})
		return
	}

	parseKeyValue(kv[0], kv[1])
}

func parseFlag(prefix string, flag string) {
	if strings.ContainsRune(flag, '=') {
		parseOption(flag)
		return
	}

	switch flag {
	case "h", "help":
		printHelp()
		os.Exit(0)
	case "v", "version":
		printVersion()
		os.Exit(0)
	case "V", "verbose":
		verbose = VerbInfo

		verbVerboseEnabled()
	case "q", "quiet":
		verbose = VerbQuiet
	case "copying":
		printCopying()
		os.Exit(0)
	case "warranty":
		printWarranty()
		os.Exit(0)
	case "e", "extended-colors":
		progressLSColors = true
	case "colortest":
		printColortest()
		os.Exit(0)
	case "n", "no-clobber", "no-overwrite":
		onExistingFile = Skip
	case "f", "force", "overwrite":
		onExistingFile = Overwrite
	case "i", "interactive":
		onExistingFile = Ask
	case "color":
		color.Init(true)
	case "dry":
		dryrun = true

		verbDryrun()
	default:
		errUnknownOption(prefix + flag)
	}
}

func configInterpretAutoOrBoolean(v string) int {
	switch v {
	case "never", "false", "no", "none":
		return 0
	case "auto":
		return 1
	case "always", "true", "yes", "all":
		return 2
	default:
		return -1
	}
}

func configInterpretBoolean(v string) bool {
	switch v {
	case "true", "on", "yes", "always":
		return true
	default:
		return false
	}
}
