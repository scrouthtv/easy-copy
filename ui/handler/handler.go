package handler

import (
	"easy-copy/tasks"
	"easy-copy/ui"
	"fmt"
)

func Handle() {
	for !tasks.Done {
		select {
		case w := <-ui.Warns:
			fmt.Println(w)
		case i := <-ui.Infos:
			fmt.Println(i)
		}
	}
}
