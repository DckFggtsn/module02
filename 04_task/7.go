package main

import (
	"fmt"
	"sort"
)

func main() {
	m := map[string]int{"first": 1, "second": 2}
	/*_, ok := m["third"]
	if !ok {
		fmt.Println("Key not exist")
		return
	}*/
	//delete(m, "second")
	m["third"] = 3
	//fmt.Println(m)
	editMap(m)
	fmt.Println(m)
	for k, v := range m {
		fmt.Println(k, v)
	}

	m2 := map[int]map[string]int{
		2020: {
			"books": 10,
			"per":   8,
		},
		2021: {
			"books": 12,
			"per":   10,
		},
	}

	m2[2020] = map[string]int{"books": 5, "per": 3}
	a, ok := m2[2022]
	fmt.Println(m2, a, ok)
	book := make(map[string]int, 4)
	book = map[string]int{
		"bk2": 34,
		"bk3": 23,
		"bk1": 11,
		"bk4": 555,
	}

	keys := make([]string, 0, len(book))
	for k, v := range book {
		fmt.Println(k, v)
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for k, v := range keys {
		fmt.Println(k, v, book[v])
	}
	fmt.Println(len(keys), cap(keys))

	fmt.Println()
	bks := make(map[string]map[string][]string, 5)
	bks = map[string]map[string][]string{
		"client1": {
			"books": {"bk34", "bk67"},
			"pers":  {},
		},
		"client3": {
			"books": {"bk45"},
			"pers":  {"per321"},
		},
		"client2": {
			"books": {},
			"pers":  {"per321", "per5678", "per450"},
		},
		"client4": {
			"books": {},
			"pers":  {},
		},
	}

	for cl, v := range bks {
		f := false
		cnt := 0
		for _, v2 := range v {
			if len(v2) != 0 {
				f = true
			}
			cnt += len(v2)
		}
		if f == true {
			fmt.Println(cl, cnt)
		}
	}
}

func editMap(m map[string]int) {
	m["fourth"] = 4
}
