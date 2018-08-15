package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	mapD := map[string]int{"apples": 10, "banana": 7}
	mapB, _ := json.Marshal(mapD)
	fmt.Println(string(mapB))

	fruts := []string{"apple", "peach", "pear"}
	vegatables, _ := json.Marshal(fruts)
	fmt.Println(string(vegatables))

	byt := []byte(`{"num":6.13,"strs":["a","b"]}`)
	var dat map[string]interface{}
	if err := json.Unmarshal(byt, &dat); err != nil {
		panic(err)
	}
	fmt.Println(dat)

	num := dat["num"].(float64)
	fmt.Println(num)
	enc := json.NewEncoder(os.Stdout)
	d := map[string]int{"apple": 5, "banana": 7}
	enc.Encode(d)
}
