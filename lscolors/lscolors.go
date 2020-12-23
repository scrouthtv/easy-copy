// +build !nolscolors

package lscolors

import "os"
import "strings"

type lscolors struct {
	types map[string]string
	exts  map[string]string
}

var types []string = []string{
	"no", "fi", "rs", "di", "ln", "mh", "pi",
	"so", "do", "bd", "cd", "or", "mi", "su",
	"sg", "ca", "tw", "ow", "st", "ex",
}

var lsc_loaded bool = false
var lsc lscolors = lscolors{
	make(map[string]string), make(map[string]string),
}

func FormatType(t string) string {
	return lsc.types[t]
}

func FormatFile(info os.FileInfo) string {
	if !lsc_loaded {
		ReloadLsColors()
	}
	if info == nil {
		return ""
	}
	//                      rwxrwxrwx
	if info.Mode().Perm()&0b001000000 != 0 && info.Mode().IsRegular() {
		return lsc.types["ex"]
	}
	switch info.Mode() & os.ModeType {
	case os.ModeDir:
		return lsc.types["di"]
	case os.ModeSymlink:
		return lsc.types["ln"]
	case os.ModeDevice:
		return lsc.types["bd"]
	case os.ModeNamedPipe:
		return lsc.types["pi"]
	case os.ModeSocket:
		return lsc.types["so"]
	case os.ModeSetuid:
		return lsc.types["su"]
	case os.ModeSetgid:
		return lsc.types["sg"]
	case os.ModeCharDevice:
		return lsc.types["cd"]
	case os.ModeSticky:
		//                      rwxrwxrwx
		if info.Mode().Perm()&0b000000010 == 0 {
			return lsc.types["ex"]
		} else {
			return lsc.types["st"]
		}
	case os.ModeIrregular:
		return lsc.types["mi"]
	default:
		return formatByExtension(info.Name())
	}
}

func formatByExtension(name string) string {
	// the coreutils ls implementation of LS_COLORS matching only works
	// if there is a single * in the beginning, so I'm going to do the same thing:
	name = strings.ToLower(name)
	var ext, format string
	for ext, format = range lsc.exts {
		if strings.HasSuffix(name, ext) {
			return format
		}
	}
	return lsc.types["fi"]
}

func ReloadLsColors() {
	var lscvar string
	var ok bool
	lscvar, ok = os.LookupEnv("LS_COLORS")
	if !ok {
		lscvar, ok = os.LookupEnv("LS_COLOURS")
	}
	if !ok {
		return
	}

	var colors []string
	var eqsym int
	colors = strings.Split(lscvar, ":")
	var clr string
	for _, clr = range colors {
		eqsym = strings.IndexRune(clr, '=')
		if eqsym == -1 {
			continue
		}
		if lscIsType(clr[0:eqsym]) {
			lsc.types[clr[:eqsym]] = clr[eqsym+1:]
		} else {
			// assume the text starts with an asterisk:
			// also lowercase it, since ls also does that
			lsc.exts[clr[1:eqsym]] = strings.ToLower(clr[eqsym+1:])
		}
	}
	lsc_loaded = true
}

func lscIsType(key string) bool {
	var t string
	for _, t = range types {
		if key == t {
			return true
		}
	}
	return false
}
