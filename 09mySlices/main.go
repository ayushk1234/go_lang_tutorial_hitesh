package main

import (
	"fmt"
	"sort"
)

func main() {

	var fruits = []string{}
	fmt.Printf("%T", fruits)

	fruits = append(fruits, "Mango", "orage")

	fmt.Println(fruits)

	fruits = append(fruits[1:])
	fmt.Println(fruits)

	highScores := make([]int, 4)
	highScores[0] = 234
	highScores[1] = 135

	highScores[2] = 236
	highScores[3] = 37

	fmt.Println(highScores)

	sort.Ints(highScores)
	fmt.Println(highScores)

	highScores = append(highScores[:2], highScores[3:]...)
	fmt.Println(highScores)

}
