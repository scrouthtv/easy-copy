package main;

import "fmt";
import "math";
import "time";

const BAR_WIDTH int = 60;

func drawLoop() {
	for !done {
		var BAR_FILLED int;
		if full_size == 0 {
			BAR_FILLED = BAR_WIDTH / 2;
		} else {
			BAR_FILLED = int(math.Round(float64(BAR_WIDTH) * float64(done_amount) / float64(full_size)));
		}

		fmt.Print("\r"); // up one line to overwrite the previous bar
		fmt.Print("  [");
		var i int;
		for i = 0; i < BAR_FILLED - 1; i++ { fmt.Print("="); }
		fmt.Print(">");
		for i = BAR_FILLED; i < BAR_WIDTH; i++ { fmt.Print(" "); }
		fmt.Println("]")


	}
}
