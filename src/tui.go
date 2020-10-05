package main;

import "fmt";
import "math";
import "time";

const BAR_WIDTH int = 60;

var drawBar bool = true;
var drawAskOverwriteDialog bool = false;

// contains ids to files that should be recopied after the
//  dialog whether to overwrite files has been answered.
// once their respective dialogs have been answered, they are either
//  added to pendingOverwrites or simply removed from piledOverwrites.
var piledConflicts []int;
var pendingConflicts []int;

func drawLoop() {
	fmt.Println();
	for !done {
		//fmt.Print("\033[4A"); // up one line to overwrite the previous bar
		if drawBar {
			var BAR_FILLED int;
			if full_size == 0 {
				BAR_FILLED = BAR_WIDTH / 2;
			} else {
				BAR_FILLED = int(math.Round(float64(BAR_WIDTH) * float64(done_size) / float64(full_size)));
			}

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

		if drawAskOverwriteDialog {
			fmt.Println("start dialog");
			filesLock.RLock();
			var conflictID int = piledConflicts[0];
			var conflict string = fileOrder[conflictID];
			var cTarget string = targets[conflict];
			filesLock.RUnlock();
			fmt.Println();
			fmt.Print(FGColors.Magenta);
			fmt.Print(conflict, " already exists in ", cTarget, ".");
			fmt.Print("Do you want to [S]kip or [O]verwrite?");
			fmt.Println(Textstyle.Reset);
		} else {
			fmt.Println("hewo");
		}

		time.Sleep(100 * time.Millisecond);
	}
}
