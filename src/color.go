package main;

var FGColors colors;
var BGColors colors;
var Textstyle textstyle;

/**
 * Value determines whether colors should be enabled.
 * true  - colors are enabled
 * false - colors are disabled
 */
func initColors(value bool) {
	if value {
		FGColors = colors{"\033[39m",
			"\033[30m", "\033[31m", "\033[32m", "\033[33m",
			"\033[34m", "\033[35m", "\033[36m", "\033[37m",
			"\033[90m", "\033[91m", "\033[92m", "\033[93m",
			"\033[94m", "\033[95m", "\033[96m", "\033[97m" };
		BGColors = colors{"\033[49m",
			"\033[40m", "\033[41m", "\033[42m", "\033[43m",
			"\033[44m", "\033[45m", "\033[46m", "\033[47m",
			"\033[100m", "\033[101m", "\033[102m", "\033[103m",
			"\033[104m", "\033[105m", "\033[106m", "\033[107m" };
		Textstyle = textstyle{
			"\033[0m", "\033[1m", "\033[2m",
			"\033[4m", "\033[5m", "\033[7m", "\033[8m" };
	} else {
		FGColors = colors{};
		BGColors = colors{};
		Textstyle = textstyle{};
	}
}

type colors struct {
	Default string;
	Black string;
	Red string;
	Green string;
	Yellow string;
	Blue string;
	Magenta string;
	Cyan string;
	LGray string;
	DGray string;
	LRed string;
	LGreen string;
	LYellow string;
	LBlue string;
	LMagenta string;
	LCyan string;
	White string;
}

type textstyle struct {
	Reset string;
	Bold string;
	Dim string;
	Underlined string;
	Blink string;
	Reverse string;
	Hidden string;
}

