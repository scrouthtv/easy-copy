package main;

import "fmt";

func main() {
	var err error = reflink("/mnt2/src", "/mnt2/reflink");
	if err == nil {
		fmt.Println("Done.");
	} else {
		fmt.Println(err);
	}
}
