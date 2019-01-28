package main

import "fmt"

type hhh int

var x hhh
var y int

func main() {
	fmt.Println("x= ", x)
	fmt.Printf("%T\n", x)
	x = 42
	fmt.Println("x=", x)
	y = int(x)

	fmt.Println("y= ", y)
	fmt.Printf("%T\n", y)
}
