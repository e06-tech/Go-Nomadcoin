package main

import "fmt"

func main() {
	a := 2
	b := &a //memory adress
	a = 5000
	a = 12000
	//c := "Hello world!"
	//fmt.Println(b, &a) //&: memory adress
	fmt.Println(*b) //memory adress --> value
} 
