package main

import "fmt"

func main() {
	fmt.Println("pointers class")

	myNumber := 23

	var ptr = &myNumber
	fmt.Println(ptr)
	fmt.Println(*ptr)
	fmt.Println(*&myNumber)

	*ptr = *ptr * 2
	fmt.Println(myNumber)
}
