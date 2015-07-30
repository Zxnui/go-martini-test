package main

import (
	"log"

	"github.com/go-martini/martini"
	"github.com/martini-contrib/gzip"
	"github.com/martini-contrib/render"
	"github.com/martini-contrib/sessions"

	"test4/controllers"
	"test4/models"
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

	store := sessions.NewCookieStore([]byte("secret123"))
	m.Use(sessions.Sessions("mysession", store))

	controllers.RouterInit(m)

	models.ModelInit()

	m.Run()
}
