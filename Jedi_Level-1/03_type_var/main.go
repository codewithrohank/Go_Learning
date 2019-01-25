package main

import "fmt"

type xxx int

var x xxx

func main() {

	fmt.Println(x)
	fmt.Printf("%T\n", x)

	x = 42

	fmt.Println(x)

}
