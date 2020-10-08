package main;

import "fmt";

func main() {
	var err error = Always("/media/test.a", "/media/refcopy");
	if err != nil { fmt.Println(err); }
}
