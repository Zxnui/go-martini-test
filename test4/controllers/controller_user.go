package controllers

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"

	"test4/middlewares"
	"test4/models"
)

func RouterUserInit(m *martini.ClassicMartini) {
	m.Group("/user", func(r martini.Router) {
		r.Get("/list", ListUser)
		r.Get("/upd/:id", UpdUser)
		r.Post("/upd/todo", UpdUserTodo)
		r.Get("/del/:id", DelUser)
	}, middlewares.Auth())
}

func ListUser(r render.Render) {
	users := models.GetUsers()

	data := make(map[string]interface{})
	data["Users"] = users

	r.HTML(200, "userlist", data)
}

func UpdUser(r render.Render, params martini.Params) {
	idStr := params["id"]
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		r.Redirect("/user/list")
		return
	}

	user := models.GetUserById(id)

	data := make(map[string]interface{})
	data["User"] = user

	r.HTML(200, "userupd", data)
}

func UpdUserTodo(r render.Render, rq *http.Request) {
	for i := 0; i < 1; i++ {

		idStr := strings.TrimSpace(rq.FormValue("id"))
		id, err := strconv.ParseInt(idStr, 10, 64)
		if err != nil {
			break
		}

		username := strings.TrimSpace(rq.FormValue("username"))
		if len(username) == 0 {
			break
		}

		password := strings.TrimSpace(rq.FormValue("password"))
		if len(password) == 0 {
			break
		}

		user := make(map[string]interface{}, 2)
		user["username"] = username
		user["password"] = password

		models.UpdUser(id, user)
	}

	r.Redirect("/user/list")
}

func DelUser(r render.Render, params martini.Params) {
	idStr := params["id"]
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err == nil {
		models.DelUserById(id)
	}

	r.Redirect("/user/list")
}
