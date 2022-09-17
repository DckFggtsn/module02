package main

import (
	"fmt"
)

func main() {
	//s := make([]int, 4, 6)
	d := []string{"s", "o", "n", "n", "y"}
	d1 := d[:3]
	d2 := d[2:4]
	fmt.Println(d1, d2)
	d2[0] = "N"
	fmt.Println(d1, d2, d)
	d1 = append(d1, "N", "Y", "b", "o", "y")
	fmt.Println(d1, cap(d1))
	sc := []int{3, 4, 5, 6, 7, 8}
	dc := make([]int, 3, 10)
	fmt.Println(dc)
	qc := copy(dc, sc)
	fmt.Println(sc, dc, qc)
	days := []string{"mon", "tue", "wed", "thu", "fri", "sat", "sun"}
	work := make([]string, 5, 5)
	_ = copy(work, days)
	days = days[5:]
	fmt.Println(days, work)
	days = append(work, days...)
	fmt.Println(days, len(days), cap(days))
}
