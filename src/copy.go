package main

import "bytes"
import "io"
import "fmt"
import "strings"

var currentfile int = 0;

func copyFiles() {
	for i := 0; i < len(fileOrder); i++ {
		filesLock.RLock();
		defer filesLock.RUnlock();
		var source string = fileOrder[i];
		var dest string = targets[source];
		fmt.Println("src: ", source, "dest: ", dest);

		var reader io.Reader = strings.NewReader("test");
		reader = io.TeeReader(reader, LogProcessWriter{});

		var buf bytes.Buffer;
		io.Copy(&buf, reader);
	}
}

type LogProcessWriter struct {}

// Track progress,
// see https://groups.google.com/g/golang-nuts/c/8sdk5qkTRjM/m/nRYl8exeEQAJ
func (pw LogProcessWriter) Write(data []byte) (int, error) {
	fmt.Printf("wrote %d bytes\n", len(data));
	return len(data), nil;
}
