package main

import "os"

// syncdel deletes a list of files synchronously.
func syncdel(files *[]string) {
	var path string
	var err error
	for _, path = range *files {
		currentTaskType = 4
		currentFile = path
		err = os.RemoveAll(path)
		if err != nil {
			errDeletingFile(path, err)
		}
	}
}
