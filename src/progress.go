package main;

import "fmt";
import "math";
import "time";

const BAR_WIDTH int = 60;

var drawBar bool = false;

func drawLoop() {
	fmt.Println();
	for !done {
		if drawBar {
			var BAR_FILLED int;
			if full_size == 0 {
				BAR_FILLED = BAR_WIDTH / 2;
			} else {
				BAR_FILLED = int(math.Round(float64(BAR_WIDTH) * float64(done_size) / float64(full_size)));
			}

			fmt.Print("\033[1A"); // up one line to overwrite the previous bar
			fmt.Print("  [");
			var i int;
			for i = 0; i < BAR_FILLED - 1; i++ { fmt.Print("="); }
			if BAR_FILLED == BAR_WIDTH {
				fmt.Print("=");
			} else {
				fmt.Print(">");
			}
			for i = BAR_FILLED; i < BAR_WIDTH; i++ { fmt.Print(" "); }
			fmt.Print("] ")
			fmt.Print(done_size / 1024);
			fmt.Print("k / ");
			fmt.Print(full_size / 1024);
			fmt.Println("k");
		}

		time.Sleep(10 * time.Millisecond);
	}
}
