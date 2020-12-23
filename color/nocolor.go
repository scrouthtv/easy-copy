// +build nocolor

package color

/**
 * Value determines whether colors should be enabled.
 * true  - colors are enabled
 * false - colors are disabled
 */
func Init(value bool) {
	FGColors = colors{}
	BGColors = colors{}
	Text = textstyle{}
}
