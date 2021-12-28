package main

import "fmt"

/*func plus(a int, b int, name string) (int, string) {
	return a+b, name
}

func main() {
	result, name := plus(2,2, "Nico")
	fmt.Println(result, name)
}*/


// func plus(a ...int) (int) {
// 	/*range //iteration function */
// 	total := 0
// 	for _, item := range a {
// 		total += item
// 	} //the most important function in this section
// 	return total
// }

// func main() {
// 	result := plus(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
// 	fmt.Println(result)
// }

func main() {
	name := "Nicolas is my name."
	for _, letter := range name {
		fmt.Println(string(letter))
	}
}
