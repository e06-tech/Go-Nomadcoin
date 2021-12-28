package main

import "fmt"

func main() {
	x := 440203829402
	/*fmt.Printf("%b\n", x)
	fmt.Printf("%o\n", x)
	fmt.Printf("%x\n", x)
	fmt.Printf("%U\n", x)*/
	xAsBinary := fmt.Sprintf("%b\n", x)
	fmt.Println(x, xAsBinary)
}
