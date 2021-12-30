package main

import "fmt"

type person struct {
	name string
	age int
	realAge int
	chZodiac string
}

func (instant person) sayHello() {
	fmt.Printf("Hello! My name is %s, Korean age is %d and my real age is %d.\n", instant.name, instant.age, instant.realAge)
	fmt.Printf("Also, my Chinese zodiac is %s!\n", instant.chZodiac)
}

func main() {
	Epark := person{"Epark", 23, 22, "Rabbit"}
	Epark.sayHello()
	Skim := person{"Skim", 52, 51, "Dog"}
	Skim.sayHello()
}