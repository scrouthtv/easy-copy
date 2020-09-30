package main

import "bytes"
import "io"
import "fmt"
import "strings"

func copyFiles() {
	for !done {
		var source string = nextMapPair(files);
		var dest string = files[source];
		fmt.Println("src: ", source, "dest: ", dest);
		delete(files, source);

		var reader io.Reader = strings.NewReader("test");
		reader = io.TeeReader(reader, LogProcessWriter{});

		var buf bytes.Buffer;
		io.Copy(&buf, reader);


		// check if done:
		if len(unsearchedPaths) == 0 && len(files) == 0 {
			done = true;
		}
	}
}

type LogProcessWriter struct {}

// Track progress,
// see https://groups.google.com/g/golang-nuts/c/8sdk5qkTRjM/m/nRYl8exeEQAJ
func (pw LogProcessWriter) Write(data []byte) (int, error) {
	fmt.Printf("wrote %d bytes\n", len(data));
	return len(data), nil;
}
