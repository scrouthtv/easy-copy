package main

import "os"
import "strings"
import "errors"
import "strconv"

import "github.com/scrouthtv/easy-copy/color"

var verbose int = VERB_NOTICE

const (
	VERB_QUIET = iota
	VERB_CRIT
	VERB_NOTICE
	VERB_INFO
	VERB_DEBUG
)

// 0 skip
// 1 overwrite
// 2 ask
var onExistingFile uint8 = 2

// 0 ignore symlinks
// 1 follow symlinks, copying them as links
// 2 fully dereference
var followSymlinks uint8 = 1

// 0 never  -> no reflinks
// 1 auto   -> attempt reflink, if that fails simply copy
// 2 always -> attempt reflink, if that fails, fail
var doReflinks uint8 = 0

var progressLSColors bool = false

func parseKeyValue(key string, value string) {
	key = strings.ToLower(strings.Trim(key, " \t'\""))
	value = strings.ToLower(strings.Trim(value, " \t'\""))
	switch key {
	case "verbose":
		if configInterpretBoolean(value) {
			verbose = VERB_INFO
		}
	case "quiet":
		if configInterpretBoolean(value) {
			verbose = VERB_QUIET
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
		var val int
		var err error
		val, err = strconv.Atoi(value)
		if err == nil {
			BUFFERSIZE = uint(val)
			buf = make([]byte, BUFFERSIZE)
		}
	default:
		warnBadConfigKey(key)
	}
}

func parseOption(line string) {
	var kv []string = strings.Split(line, "=")
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
		verbose = VERB_INFO
		verbVerboseEnabled()
	case "q", "quiet":
		verbose = VERB_QUIET
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
	case "n", "no-clobber": //case "no-overwrite":
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
