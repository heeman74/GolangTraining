package main

import (
	"fmt"
)

// var g stirng // declaring the variables.
// g = "cowboy" // this is assignment.

// var e string = "initialization" // this is initialization

func main() {

	a := 10 // declaring the variables in short hand.
	b := "golang"
	c := 4.17
	d := true

	fmt.Printf("%v \n", a) // %v is use the default formatt. whic will use %d since a is integer
	fmt.Printf("%v \n", b)
	fmt.Printf("%v \n", c)
	fmt.Printf("%v \n", d)
}
