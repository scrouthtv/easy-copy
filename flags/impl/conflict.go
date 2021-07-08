package impl

import "easy-copy/flags"

var conflict flags.Conflict = 0

func OnConflict() flags.Conflict {
	return conflict
}
