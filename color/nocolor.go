//go:build nocolor || (!windows && !linux && !freebsd && !openbsd && !netbsd && !dragonfly && !darwin)

package color

// Init initializes no colors if built with nocolor.
func Init(value bool) {
	FGColors = colors{}
	BGColors = colors{}
	Text = textstyle{}
}

// AutoColors returns false if built with nocolor.
func AutoColors() bool {
	return false
}
