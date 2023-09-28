package main

import "fmt"

func main() {

	var fruits [4]string

	fruits[0] = "Apple"
	fruits[1] = "Tomato"

	fmt.Println(fruits[2])

	var veg = [3]string{"veg1", "veg2", "veg3"}
	fmt.Printf("%T", veg)

}
