package main

import "fmt"

// func main() {
// 	foods := [3]string{"potato", "chicken", "pasta"}
// 	/*
// 	for _, food := range foods {
// 		fmt.Println(food)
// 	}*/
// 	for i := 0; i < len(foods); i++ {
// 		fmt.Println(foods[i])
// 	}
// 	fmt.Println(foods[1]) //two chickens!
// } 

//slice (very simliar with string, but we can make as long as we want)
func main() {
	foods := []string{"potato", "chicken", "pasta"}
	fmt.Printf("%v\n", foods)
	fmt.Println(len(foods))
	foods = append(foods, "crisps")
	fmt.Printf("%v\n", foods)
	fmt.Println(len(foods))
} 
