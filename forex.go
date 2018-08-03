package main

import (
	"fmt"
	"time"
)

func main() {
	layout := "2006.01.02"
	t1, err := time.Parse(layout, "2017.03.01")
	if err != nil {
		fmt.Println(err)
	}
	t2, err := time.Parse(layout, "2018.08.31")
	if err != nil {
		fmt.Println(err)
	}

	moscow, _ := time.LoadLocation("Europe/Moscow")
	tm := time.Now().In(moscow)

	if tm.Unix() > t1.Unix() && tm.Unix() <= t2.Unix() {

	}

}
