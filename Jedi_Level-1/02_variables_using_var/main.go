package main

import (
	"fmt"
	_ "fmt"
)

var x int = 42
var y string = "Rohan"
var z bool = true
var s string

func main() {
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)

	s = fmt.Sprintf("%v %v %v", x, y, z)

	fmt.Println(s)
}
