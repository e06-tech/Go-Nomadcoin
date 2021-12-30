package main

import "fmt"

type person struct {
	name string
	age int
}

func (p person) sayHello() {
	fmt.Printf("Hello! My name is %s and I'm %d", p.name, p.age)
} //p: person type. person that call 'sayHello method'
//we can use 'instance' instead of 'p'
//'p' is the initial letter of person, and we can use it instead of 'instance' 

func main() {
	nico := person{"Epark", 23} //Doesn't matter wheter using label or not. The point is, the orer of variables must be aligned.
	nico.sayHello()
} 
