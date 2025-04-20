package main

import (
	"fmt"
)

func main() {
	var name = "Yummy"
	var age int = 24
	var isStudent = true
	var weight float32 = 60.5

	fmt.Println(name)
	fmt.Println(age)
	fmt.Println(isStudent)
	fmt.Println(weight)

	fmt.Printf("%T\n", name)      // string
	fmt.Printf("%T\n", age)       // int
	fmt.Printf("%T\n", isStudent) // bool
	fmt.Printf("%T\n", weight)    // float32
}
