package main

import (
	"net/http"

	"github.com/go-martini/martini"
)

func main() {
	m := martini.Classic()

	// http://127.0.0.1:3000/
	m.Get("/", func() string {
		return "Hello world!"
	})

	// http://127.0.0.1:3000/hello/world
	m.Get("/hello/:name", func(params martini.Params) string {
		return "Hello " + params["name"]
	})

	// http://127.0.0.1:3000/hello?key=world
	m.Get("/hello", func(req *http.Request) string {
		return "Hello " + req.FormValue("key")
	})

	m.Run()
}
