package handler

import (
	"easy-copy/color"
	"easy-copy/progress"
	"easy-copy/ui"
)

func Handle() {
	for !progress.CopyDone {
		select {
		case w := <-ui.Warns:
			println(color.FGColors.Red + w.Error() + color.Text.Reset)
		case i := <-ui.Infos:
			println(color.FGColors.Cyan + i.Info() + color.Text.Reset)
		}
	}
}
