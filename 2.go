package main

import (
	"fmt"
)

func main() {
	arr01 := [...]bool{1: true, 5: true}
	fmt.Println(arr01)
	arr02 := arr01[2:]
	fmt.Println(arr02)
	arr02[0] = true
	fmt.Println(arr02)
	fmt.Println(arr01)
}
