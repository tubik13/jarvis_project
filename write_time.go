package main

import (
	"os"
	"path/filepath"
	"time"
)

func main() {
	newpath := filepath.Join(".", "public/kura/rura")
	os.MkdirAll(newpath, os.ModePerm)
	a := filepath.Join("public/kura/rura", "hello.bin")
	file, _ := os.Create(a)

	t := time.Now()
	d1 := []byte(t.Format("20060102150405"))
	file.Write(d1)

}
