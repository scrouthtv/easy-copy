package flags

import (
	"easy-copy/color"
	"fmt"
)

func VerbFlags() {
	if Current.Verbosity() >= VerbInfo {
		fmt.Print(color.FGColors.Green)
		fmt.Println(" Verbose:", Current.Verbosity())
		fmt.Println(" Overwrite Mode:", Current.OnConflict())
		fmt.Println(" Follow symlinks:", Current.OnSymlink())
		fmt.Println(" Dryrun:", Current.Dryrun())
		fmt.Println(" Parallel:", Current.Parallel(), color.Text.Reset)
	}
}
