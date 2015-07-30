package main

import (
	"github.com/go-martini/martini"
	"github.com/martini-contrib/gzip"
	"github.com/martini-contrib/render"

	"test2/controllers"
)

func main() {
	m := martini.Classic()

	//Make sure to include the Gzip middleware above other middleware that alter the response body (like the render middleware).
	m.Use(gzip.All())

	// render html templates from templates directory
	m.Use(render.Renderer(render.Options{
		Layout: "layout",
	}))

	controllers.RouterInit(m)

	m.Run()
}
