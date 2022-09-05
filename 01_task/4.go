package main

import (
	"fmt"
	"math"
)

func main() {
	type UsaVel float64
	type EurVel float64
	ev := 120.4 * 3.6
	uv := 130 * 2.24
	var EV1 EurVel = EurVel(ev)
	var UV1 UsaVel = UsaVel(uv)
	fmt.Println(math.Round(float64(EV1)))
	fmt.Println(math.Round(float64(UV1)))
}
