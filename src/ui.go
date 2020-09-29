package main

import "fmt"

func printHelp() {
	fmt.Println("help");
}

func printVersion() {
	fmt.Println("version");
}

func unknownOption(option string) {
	fmt.Println("unrecognized option: " + option);
	printHelp();
}
