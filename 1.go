package main

import (
	"fmt"
	"time"
)

func main() {
	dt := time.Now()
	fmt.Println("Date & Time:", dt.Format("02.01.2006 15:04"))
}
