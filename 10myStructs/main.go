package main

import "fmt"

func main() {

	ayush := User{"usernamae", "go.dev", true, 4}
	fmt.Println(ayush)
	fmt.Printf("%+v", ayush)
	fmt.Printf("%v", ayush.Email)

}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
