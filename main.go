package main

import "person/person.go"

func main() {
	nico := person.Person{}
	nico.SetDetails("nico", 12)
	fmt.Println(nico)
	fmt.Println("Main's 'nico'", nico)
}