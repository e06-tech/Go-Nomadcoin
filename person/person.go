package main

type Person struct {
	name string
	age int
}

func (p Person) SetDetails(name string, age int){
	p.name = name
	p.age = age
	fmt.Println("SeeDetails' nico:", p)
}