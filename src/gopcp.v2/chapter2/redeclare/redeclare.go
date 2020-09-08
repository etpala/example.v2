package main

import (
	"fmt"
)

var v = "1, 2, 3"

func init() {
	fmt.Printf("%v\n", v)
}

func main() {
	v := []int{1, 2, 3}
	fmt.Printf("%v\n", v)
	if v != nil {
		var v = 123
		fmt.Printf("%v\n", v)
	}
	fmt.Printf("%v\n", v)
}
