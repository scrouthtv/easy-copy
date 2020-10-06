package main

import (
	"fmt"
	"strings"
)

func tbprint(x, y int, fg, bg Attribute, msg string) {
	for _, c := range msg {
		SetCell(x, y, c, fg, bg)
		x++
	}
}

var current string
var curev Event

func mod_str(m Modifier) string {
	var out []string
	if m&ModAlt != 0 {
		out = append(out, "ModAlt")
	}
	if m&ModMotion != 0 {
		out = append(out, "ModMotion")
	}
	return strings.Join(out, " | ")
}

func redraw_all() {
	const coldef = ColorDefault
	Clear(coldef, coldef)
	tbprint(0, 0, ColorMagenta, coldef, "Press 'q' to quit")
	tbprint(0, 1, coldef, coldef, current)
	switch curev.Type {
	case EventKey:
		tbprint(0, 2, coldef, coldef,
			fmt.Sprintf("EventKey: k: %d, c: %c, mod: %s", curev.Key, curev.Ch, mod_str(curev.Mod)))
	}
	tbprint(0, 3, coldef, coldef, fmt.Sprintf("%d", curev.N))
	Flush()
}

func main() {
	err := Init()
	if err != nil {
		panic(err)
	}
	defer Close()
	SetInputMode(InputAlt);
	redraw_all()

	data := make([]byte, 0, 64)
mainloop:
	for {
		if cap(data)-len(data) < 32 {
			newdata := make([]byte, len(data), len(data)+32)
			copy(newdata, data)
			data = newdata
		}
		beg := len(data)
		d := data[beg : beg+32]
		switch ev := PollRawEvent(d); ev.Type {
		case EventRaw:
			data = data[:beg+ev.N]
			current = fmt.Sprintf("%q", data)
			if current == `"q"` {
				break mainloop
			}

			for {
				ev := ParseEvent(data)
				if ev.N == 0 {
					break
				}
				curev = ev
				copy(data, data[curev.N:])
				data = data[:len(data)-curev.N]
			}
		case EventError:
			panic(ev.Err)
		}
		redraw_all()
	}
}
