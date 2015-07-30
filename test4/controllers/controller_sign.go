package controllers

import (
	"net/http"
	"strings"

	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	"github.com/martini-contrib/sessions"

	"test4/models"
	"time"
)

func RouterSiginInit(m *martini.ClassicMartini) {
	m.Group("/sign", func(r martini.Router) {
		r.Get("/in", SignIn)
		r.Post("/in/todo", SignInTodo)
		r.Get("/up", SignUp)
		r.Post("/up/todo", SignUpTodo)
	})
}

func SignIn(r render.Render, rq *http.Request) {

	data := make(map[string]interface{})

	cname, err := rq.Cookie("username")
	if err == nil {
		data["Username"] = cname.Value
	}

	cpass, err := rq.Cookie("passwrod")
	if err == nil {
		data["Passwrod"] = cpass.Value
	}

	r.HTML(200, "signin", data)
}

func SignInTodo(r render.Render, session sessions.Session, rq *http.Request, rw http.ResponseWriter) {

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

		if user.Password != password {
			break
		}

		session.Set("user", user)

		remember := strings.TrimSpace(rq.FormValue("remember"))

		if remember == "on" {
			expiration := time.Now().Add(7 * 24 * time.Hour)
			cookie1 := http.Cookie{Name: "username", Value: username, Expires: expiration, HttpOnly: true}
			cookie2 := http.Cookie{Name: "password", Value: password, Expires: expiration, HttpOnly: true}
			http.SetCookie(rw, &cookie1)
			http.SetCookie(rw, &cookie2)
		}

		r.Redirect("/user/list")
		return
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
