package ui

import (
	"easy-copy/color"
	"fmt"
)

func ShowColortest() {
	fmt.Println(color.FGColors.Default + "Default")
	fmt.Println(color.FGColors.Black + "Black")
	fmt.Println(color.FGColors.Red + "Red")
	fmt.Println(color.FGColors.Green + "Green")
	fmt.Println(color.FGColors.Yellow + "Yellow")
	fmt.Println(color.FGColors.Blue + "Blue")
	fmt.Println(color.FGColors.Magenta + "Magenta")
	fmt.Println(color.FGColors.Cyan + "Cyan")
	fmt.Println(color.FGColors.LGray + "LGray")
	fmt.Println(color.FGColors.DGray + "DGray")
	fmt.Println(color.FGColors.LRed + "LRed")
	fmt.Println(color.FGColors.LGreen + "LGreen")
	fmt.Println(color.FGColors.LYellow + "LYellow")
	fmt.Println(color.FGColors.LBlue + "LBlue")
	fmt.Println(color.FGColors.LMagenta + "LMagenta")
	fmt.Println(color.FGColors.LCyan + "LCyan")
	fmt.Println(color.FGColors.White + "LWhite" + color.Text.Reset)
}
