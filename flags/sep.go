package flags

import "fmt"

// Sep separates the line into arguments by splitting at spaces.
// Spaces inside quotes ("" and '') are kept,
// as well as those escaped by a backslash \.
// The quotes and the backslash are consumed.
func Sep(line string) ([]string, error) {
	r := []rune(line)
	args := []string{}
	j := 0
	args = append(args, "")

	for i := 0; i < len(r); i++ {
		if r[i] == '"' || r[i] == '\'' { // parse quoted string
			rq := r[i] // right quote = left quote for now
			lp := i // left pos
			i++

			for r[i] != rq { // search for right quote
				i++
				if i >= len(r) {
					return nil, &ErrMissingClosingQuotes{rq, i, line}
				}
			}

			args[j] = string(r[lp+1:i]) // append entire qoted string consuming quotes
			j++
			args = append(args, "")
		} else if r[i] == '\\' {
			i++ // consume backslash
			if i < len(r) { // ignore trailing backslash
				args[j] += string(r[i]) // append next character as-is
			}
		} else if r[i] == '\r' {
			// ignore
		} else if r[i] == ' ' || r[i] == '\t' || r[i] == '\n' {
			if args[j] != "" { // don't create empty arguments
				j++
				args = append(args, "")
			}
		} else {
			args[j] += string(r[i]) // append other characters as-is
		}
	}

	if args[j] == "" {
		args = args[:j]
	}

	return args, nil
}

type ErrMissingClosingQuotes struct {
	LeftQuote rune
	LeftPos int
	Line string
}

func (e *ErrMissingClosingQuotes) Error() string {
	return fmt.Sprintf("missing closing quote for opening %q at %d: %s",
				e.LeftQuote, e.LeftPos, e.Line)
}

