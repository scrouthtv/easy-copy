package main

import "os"
import "fmt"

// Delete a list of files synchronously
func syncdel(files *[]string) {
	var path string
	var err error
	for _, path = range *files {
		fmt.Println()
		fmt.Println()
		fmt.Println(path)
		fmt.Println()
		fmt.Println()
		currentTaskType = 4
		currentFile = path
		err = os.RemoveAll(path)
		if err != nil {
			errDeletingFile(path, err)
		}
	}
}
