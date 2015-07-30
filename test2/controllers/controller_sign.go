package controllers

import (
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
)

func RouterSiginInit(m *martini.ClassicMartini) {
	m.Group("/sign", func(r martini.Router) {
		r.Get("/in", SignIn)
		r.Post("/in/todo", SignInTodo)
		r.Get("/up", SignUp)
		r.Post("/up/todo", SignUpTodo)
	})
}

func SignIn(r render.Render) {
	r.HTML(200, "signin", nil)
}

func SignInTodo(r render.Render) {
	r.Redirect("/user/list")
}

func SignUp(r render.Render) {
	r.HTML(200, "signup", nil)
}

func SignUpTodo(r render.Render) {
	r.Redirect("/sign/in")
}
