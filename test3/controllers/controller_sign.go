package controllers

import (
	"net/http"
	"strings"

	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"

	"test3/models"
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

func SignInTodo(r render.Render, rq *http.Request) {

	for i := 0; i < 1; i++ {

		username := strings.TrimSpace(rq.FormValue("username"))
		if len(username) == 0 {
			break
		}

		password := strings.TrimSpace(rq.FormValue("password"))
		if len(password) == 0 {
			break
		}

		user := models.GetUserByName(username)

		if user.Password == password {
			r.Redirect("/user/list")
			return
		}
	}

	r.Redirect("/sign/in")
}

func SignUp(r render.Render) {
	r.HTML(200, "signup", nil)
}

func SignUpTodo(r render.Render, rq *http.Request) {

	for i := 0; i < 1; i++ {

		username := strings.TrimSpace(rq.FormValue("username"))
		if len(username) == 0 {
			break
		}

		password := strings.TrimSpace(rq.FormValue("password"))
		if len(password) == 0 {
			break
		}

		user := new(models.User)
		user.Username = username
		user.Password = password

		models.SetUser(user)
	}

	r.Redirect("/sign/in")
}
