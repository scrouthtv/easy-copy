package main;

import "fmt";

func main() {
	var char rune = getch();
	fmt.Println("got something");
	fmt.Println(string(char));
}
