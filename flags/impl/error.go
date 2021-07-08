package impl

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
