package service

import (
	"MyChat/model"
	"errors"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"log"
)

var DbEngin *xorm.Engine

func init() {
	dirverName := "mysql"
	dirverSource := "root:root@(localhost:3306)/chat?charset=utf8"
	err := errors.New("")
	DbEngin, err = xorm.NewEngine(dirverName, dirverSource)
	if err != nil && err.Error() != "" {
		log.Fatal(err.Error())
	}
	DbEngin.ShowSQL(true)
	e := DbEngin.Sync(new(model.User))
	if e != nil {
		log.Println(e.Error())
	}
}
