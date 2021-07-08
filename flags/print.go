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
		fmt.Print(" Follow symlinks: ", "todo")
		fmt.Println(color.Text.Reset)
	}
}
