package main;

//#include <stdio.h>
//#include <errno.h>
//int fortytwo() {
//  return 42;
//}
import "C";

import "fmt";

func main() {
	fmt.Println(C.fortytwo());
}
