package main;

import "fmt";
import "math";
import "time";
import "path/filepath";
import "bufio";
import "os";

const BAR_WIDTH int = 60;

var drawBar bool = true;

var reader *bufio.Reader;

// contains ids to files that should be recopied after the
//  dialog whether to overwrite files has been answered.
// once their respective dialogs have been answered, they are either
//  added to pendingOverwrites or simply removed from piledOverwrites.
var piledConflicts []int;
var pendingConflicts []int;

var currentTask string = "";

var lines int = 0;

func drawLoop() {
	go speedLoop();
	reader = bufio.NewReader(os.Stdin);
	for !done {

		var i int;
		for i = 0; i < lines; i++ {
			fmt.Print("\033[2K\033[1A");
		}

		lines = 0;
		// up one line to overwrite the previous bar
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
			lines++;
			if verbose {
				fmt.Print("\033[2K");
				fmt.Print(currentTask);
				fmt.Print(" @ ");
				var kbPerSecond float32 = sizePerSecond / 1024;
				var mbPerSecond float32 = kbPerSecond / 1024;
				if mbPerSecond > 2 {
					fmt.Print(mbPerSecond);
					fmt.Println(" MB/s");
				} else {
					fmt.Print(kbPerSecond);
					fmt.Println(" kB/s");
				}
				lines++;
			}
		}

		if len(piledConflicts) > 0 {
			filesLock.RLock();
			var conflictID int = piledConflicts[0];
			var conflict string = fileOrder[conflictID];
			var cTarget string = filepath.Join(targets[conflict], 
				filepath.Base(conflict));
			filesLock.RUnlock();
			fmt.Println();
			lines++;
			fmt.Print(FGColors.Yellow, Textstyle.Bold);
			fmt.Print(conflict);
			fmt.Print(Textstyle.Reset, FGColors.Magenta);
			fmt.Print(" already exists in ");
			fmt.Print(FGColors.Yellow, Textstyle.Bold);
			fmt.Print(cTarget + "/.");
			fmt.Println(Textstyle.Reset + FGColors.Magenta);
			lines++;
			fmt.Println("[S]kip | Skip [A]ll | [O]verwrite | O[v]erwrite All");
			lines++;
			fmt.Print("[I]nfo | [D]iff | [R]ename | [E]dit target | [Q]uit");
			fmt.Println(Textstyle.Reset);
			lines++;
			var in rune = getChoice("soavidreq");
			switch in {
				case 's':
					filesLock.Lock();
					piledConflicts = piledConflicts[1:];
					filesLock.Unlock();
				case 'o':
					filesLock.Lock();
					pendingConflicts = append(pendingConflicts, conflictID);
					piledConflicts = piledConflicts[1:];
					filesLock.Unlock();
				case 'a':
					fmt.Println("skip all");
				case 'v':
					fmt.Println("overwrite all");
				case 'i':
					fmt.Println("info");
				case 'd':
					fmt.Println("diff");
				case 'r':
					fmt.Println("rename");
				case 'e':
					fmt.Println("edit");
				case 'q':
					os.Exit(0);
			}
		}

		time.Sleep(100 * time.Millisecond);
	}
}
