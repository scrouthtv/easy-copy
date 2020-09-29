package main

import "fmt"
import "os"
import "strings"

var unsearchedPaths []string;
var files map[string]string = make(map[string]string);
var pwd string;
var target string;

func iteratePaths() {
	for len(unsearchedPaths) >= 0 {
		var source string = nextMapPair(files);
		var dest string = files[source];
		fmt.Println("src: ", source, "dest: ", dest);
		delete(files, source);

		fmt.Println("end of loop");
	}
}

func main() {
	args := os.Args[1:];
	var isFiles bool = false;
	for _, arg := range args {
		if (arg == "--") {
			isFiles = true;
		} else if (isFiles) {
			unsearchedPaths = append(unsearchedPaths, arg);
		} else if (strings.HasPrefix(arg, "--")) {
			parseFlag("--", arg[2:len(arg)]);
		} else if (strings.HasPrefix(arg, "-")) {
			for i := 1; i < len(arg); i++ {
				parseFlag("-", arg[i:i+1]);
			}
		} else {
			unsearchedPaths = append(unsearchedPaths, arg);
		}
	}
}
