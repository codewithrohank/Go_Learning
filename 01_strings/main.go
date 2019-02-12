package main

import (
	"fmt"
)

func main() {
	s := "hello Rohan"
	fmt.Println(s)

	bs := []byte(s)
	fmt.Println(bs)

	for i := 0; i < len(s); i++ {

		fmt.Printf("%#U \n", bs[i])
	}

	for i, v := range s {
		fmt.Println(i, v)
	}
}
