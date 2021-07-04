// +build windows linux freebsd openbsd netbsd dragonfly darwin
// +build !nocolor

package color

// Init initializes the available colors.
func Init(value bool) {
	if value {
		FGColors = colors{
			Default:  "\033[39m",
			Black:    "\033[30m",
			Red:      "\033[31m",
			Green:    "\033[32m",
			Yellow:   "\033[33m",
			Blue:     "\033[34m",
			Magenta:  "\033[35m",
			Cyan:     "\033[36m",
			LGray:    "\033[37m",
			DGray:    "\033[90m",
			LRed:     "\033[91m",
			LGreen:   "\033[92m",
			LYellow:  "\033[93m",
			LBlue:    "\033[94m",
			LMagenta: "\033[95m",
			LCyan:    "\033[96m",
			White:    "\033[97m",
		}
		BGColors = colors{
			Default:  "\033[49m",
			Black:    "\033[40m",
			Red:      "\033[41m",
			Green:    "\033[42m",
			Yellow:   "\033[43m",
			Blue:     "\033[44m",
			Magenta:  "\033[45m",
			Cyan:     "\033[46m",
			LGray:    "\033[47m",
			DGray:    "\033[100m",
			LRed:     "\033[101m",
			LGreen:   "\033[102m",
			LYellow:  "\033[103m",
			LBlue:    "\033[104m",
			LMagenta: "\033[105m",
			LCyan:    "\033[106m",
			White:    "\033[107m",
		}
		Text = textstyle{
			Reset:      "\033[0m",
			Bold:       "\033[1m",
			Dim:        "\033[2m",
			Underlined: "\033[4m",
			Blink:      "\033[5m",
			Reverse:    "\033[7m",
			Hidden:     "\033[8m",
		}
	} else {
		FGColors = colors{}
		BGColors = colors{}
		Text = textstyle{}
	}
}
