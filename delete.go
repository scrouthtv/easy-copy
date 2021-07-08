package main

import "fmt"
import "os"

// syncdel deletes a list of files synchronously.
func syncdel(files *[]string) {
	var err error

	for _, path := range *files {
		currentTaskType = 4
		currentFile = path

		if !isnodelete(path) && !dryrun {
			err = os.RemoveAll(path)

			if err != nil {
				errDeletingFile(path, err)
			}
		}
	}
}

func isnodelete(path string) bool {
	fmt.Println("\n\n\n\ninod:" + path + "\n\n\n\n\n")
	for _, p := range nodelete {
		if p == path {
			return true
		}
	}

	return false
}
