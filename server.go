package main

import (
	"github.com/go-martini/martini"
	"time"
)

func main() {
	m := martini.Classic()
	m.Get("/", top)
	m.Run()
}

func top(params martini.Params) (int, string) {
	t := time.Now()
	return 200, "Hello World!\n" + t.Format("2006/01/02 15:04:05")
}
