package color

var FGColors colors
var BGColors colors
var Text textstyle

type colors struct {
	Default  string
	Black    string
	Red      string
	Green    string
	Yellow   string
	Blue     string
	Magenta  string
	Cyan     string
	LGray    string
	DGray    string
	LRed     string
	LGreen   string
	LYellow  string
	LBlue    string
	LMagenta string
	LCyan    string
	White    string
}

type textstyle struct {
	Reset      string
	Bold       string
	Dim        string
	Underlined string
	Blink      string
	Reverse    string
	Hidden     string
}
