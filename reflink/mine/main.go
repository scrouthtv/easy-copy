package main;

import "fmt";

func main() {
	var err error = reflink("/mnt2/notfile", "/mnt2/myreflink");
	if err == nil {
		fmt.Println("Done.");
	} else {
		fmt.Println(err);
	}
}
