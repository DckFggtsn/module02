package main

import (
	"fmt"
	"math"
)

func main() {
	R := new(float64)
	L := 35
	PI := 3.14
	*R = float64(L) / (2 * PI)
	var S float64
	S = PI * math.Pow(*R, 2)

	fmt.Printf("R = %.2f  S = %.2f", math.Round(*R*100)/100, math.Round(S*100)/100)
}
