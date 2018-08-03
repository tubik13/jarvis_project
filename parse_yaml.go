package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"gopkg.in/yaml.v2"
)

func main() {
	log.SetFlags(log.Lshortfile)

	data, err := ioutil.ReadFile("default.yaml")
	if err != nil {
		log.Fatalln(err)
	}

	var v interface{}
	err = yaml.Unmarshal(data, &v)
	if err != nil {
		log.Fatalln(err)
	}

	f, err := os.Create("default.yaml.go")
	if err != nil {
		log.Fatalln(err)
	}
	defer func() {
		err := f.Close()
		if err != nil {
			log.Fatalln(err)
		}
	}()

	fmt.Fprintf(f, "package main\n\n")
	fmt.Fprintf(f, "var regexes = %#v\n", v)
 }