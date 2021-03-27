package main

import (
	"easy-copy/color"
	"errors"
	"os"
	"strconv"
	"strings"
)

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

// onExistingFile takes one of these values:
// 0 skip
// 1 overwrite
// 2 ask.
var onExistingFile uint8 = 2

// followSymlinks takes one of these values:
// 0 ignore symlinks
// 1 follow symlinks, copying them as links
// 2 fully dereference.
var followSymlinks uint8 = 1

// doReflinks takes one of these values:
// 0 never  -> no reflinks
// 1 auto   -> attempt reflink, if that fails simply copy
// 2 always -> attempt reflink, if that fails, fail.
var doReflinks uint8 = 0

var progressLSColors bool = false

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
			onExistingFile = 0
		case "overwrite":
			onExistingFile = 1
		case "ask":
			onExistingFile = 2
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
		switch value {
		case "never", "false", "no", "none":
			color.Init(false)
		case "auto":
			color.Init(color.AutoColors())
		case "always", "true", "yes", "all":
			color.Init(true)
		}
	case "reflink":
		switch value {
		case "never", "false", "no", "none":
			doReflinks = 0
		case "auto":
			doReflinks = 1
		case "always", "true", "yes", "all":
			doReflinks = 2
		}
	case "buffersize":
		val, err := strconv.Atoi(value)
		if err == nil {
			buffersize = uint(val)
			buf = make([]byte, buffersize)
		} else {
			warnConfig(errors.New("bad value for buffersize: " + value))
		}
	default:
		warnBadConfigKey(key)
	}
}

func parseOption(line string) {
	kv := strings.Split(line, "=")
	if len(kv) != 2 {
		warnConfig(errors.New("missing '=' or too many '=' : " + line))
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
	case "f", "force":
		onExistingFile = 1
	case "i", "interactive":
		onExistingFile = 2
	case "no-config":
		doReadConfig = false
	case "color":
		color.Init(true)
	case "reflink":
		doReflinks = 2
	case "n", "no-clobber": // case "no-overwrite":
		onExistingFile = 0
	default:
		errUnknownOption(prefix + flag)
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
