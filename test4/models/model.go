package models

import (
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

var Engine *xorm.Engine

func ModelInit() {

	engine, err := xorm.NewEngine("mysql", "root:123456@/test?charset=utf8")
	if err != nil {
		log.Println(err)
	}

	if err = engine.Ping(); err != nil {
		log.Println(err)
		os.Exit(1)
	}

	engine.Sync2(new(User))

	Engine = engine
}
