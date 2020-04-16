package service

import (
	"errors"
	"fmt"
	"gitapp.com/v1/im/model"
	"github.com/go-xorm/xorm"
	"log"
)

var DbEngin *xorm.Engine
func init()  {
	drivename :="mysql"
	DsName := "root:toor@(127.0.0.1:3306)/chat?charset=utf8"
	err := errors.New("")
	DbEngin, err = xorm.NewEngine(drivename, DsName)
	if nil!=err && ""!=err.Error() {
		log.Fatal(err.Error())
	}
	//show sql
	DbEngin.ShowSQL(false)

	DbEngin.SetMaxIdleConns(2)

	//DbEngin.NewSession().InsertOne(new())
	DbEngin.Sync2(new(model.User),
	new(model.Community),
	new(model.Contact),
	)
	//if err!=nil {
	//	log.Fatal(err.Error())
	//}
	fmt.Println("init data bases ok")

}
