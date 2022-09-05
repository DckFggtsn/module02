package main

import (
	"fmt"
	"strconv"
)

func main() {
	str01 := "104"
	int01 := 35
	str01ToInt, _ := strconv.Atoi(str01)
	int01ToStr := strconv.Itoa(int01)
	fmt.Printf("%T %d %T %s", str01ToInt, str01ToInt, int01ToStr, int01ToStr)
}
