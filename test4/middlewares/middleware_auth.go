package middlewares

import (
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	"github.com/martini-contrib/sessions"

	"net/http"
)

func Auth() martini.Handler {
	return func(context martini.Context, req *http.Request, r render.Render, session sessions.Session) {

		user := session.Get("user")
		if user == nil {
			r.Redirect("/sign/in")
			return
		}

		context.Next()
	}
}
