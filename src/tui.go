package main;

import "fmt";
import "math";
import "time";
import "path/filepath";
import "bufio";
import "os";
import "strings";

const BAR_WIDTH int = 60;

var drawBar bool = true;

var reader *bufio.Reader;

// contains ids to files that should be recopied after the
//  dialog whether to overwrite files has been answered.
// once their respective dialogs have been answered, they are either
//  added to pendingOverwrites or simply removed from piledOverwrites.
var piledConflicts []int;
var pendingConflicts []int;

func drawLoop() {
	reader = bufio.NewReader(os.Stdin);
	fmt.Println();
	fmt.Println();
	fmt.Println();
	fmt.Println();
	fmt.Println();
	fmt.Println();
	fmt.Println();
	fmt.Println();
	fmt.Println();
	fmt.Println();
	fmt.Println();
	fmt.Println();
	fmt.Println();
	fmt.Println();
	fmt.Println();
	for !done {
		//fmt.Print("\033[4A"); // up one line to overwrite the previous bar
		if drawBar {
			var BAR_FILLED int;
			if full_size == 0 {
				// unneeded as this is only called after the iterator is done
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

		if len(piledConflicts) > 0 {
			fmt.Println("start dialog");
			filesLock.RLock();
			var conflictID int = piledConflicts[0];
			var conflict string = fileOrder[conflictID];
			var cTarget string = filepath.Join(targets[conflict], 
				filepath.Base(conflict));
			fmt.Println(targets);
			fmt.Println(conflict);
			filesLock.RUnlock();
			fmt.Println();
			fmt.Print(FGColors.Yellow, Textstyle.Bold);
			fmt.Print(conflict);
			fmt.Print(Textstyle.Reset, FGColors.Magenta);
			fmt.Print(" already exists in ");
			fmt.Print(FGColors.Yellow, Textstyle.Bold);
			fmt.Print(cTarget + "/.");
			fmt.Println(Textstyle.Reset + FGColors.Magenta);
			fmt.Print("Do you want to [S]kip or [O]verwrite?");
			fmt.Println(Textstyle.Reset);
			text, _ := reader.ReadString('\n');
			text = strings.ToLower(text);
			if strings.HasPrefix(text, "s") {
				filesLock.Lock();
				piledConflicts = piledConflicts[1:];
				filesLock.Unlock();
			} else if strings.HasPrefix(text, "o") {
				filesLock.Lock();
				pendingConflicts = append(pendingConflicts, conflictID);
				piledConflicts = piledConflicts[1:];
				filesLock.Unlock();
			}
		}

		time.Sleep(100 * time.Millisecond);
	}
}
