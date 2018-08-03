package main

import (
	"fmt"
	"time"
)

func main() {

	tm1 := time.Now()
	tm2 := time.Date(tm1.Year(), tm1.Month(), tm1.Day(), tm1.Hour()+1, 0, 0, 0, tm1.Location())

	fmt.Println(tm1)
	fmt.Println(tm2)

	time.Sleep(tm2.Sub(tm1))
}
