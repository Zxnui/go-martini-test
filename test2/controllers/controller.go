package controllers

import (
	"github.com/go-martini/martini"
)

func RouterInit(m *martini.ClassicMartini) {
	RouterSiginInit(m)
}
