package ui

import (
	"errors"
	"fmt"
	"os"
)

func ShowCopying() {
	err := runPager(infoCopying())
	if errors.Is(err, errNoPager) {
		fmt.Println(infoCopying())
	} else if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func ShowWarranty() {
	err := runPager(infoWarranty())
	if errors.Is(err, errNoPager) {
		fmt.Println(infoCopying())
	} else if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
