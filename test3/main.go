package main

import (
	"log"

	"github.com/go-martini/martini"
	"github.com/martini-contrib/gzip"
	"github.com/martini-contrib/render"

	"test3/controllers"
	"test3/models"
)

func init() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
}

func main() {
	m := martini.Classic()

	//Make sure to include the Gzip middleware above other middleware that alter the response body (like the render middleware).
	m.Use(gzip.All())

	// render html templates from templates directory
	m.Use(render.Renderer(render.Options{
		Layout: "layout",
	}))

	controllers.RouterInit(m)

	models.ModelInit()

	m.Run()
}
