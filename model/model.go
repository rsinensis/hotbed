package model

import (
	"fmt"
	"hotbed/module/engine"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/core"
	"github.com/go-xorm/xorm"
	_ "github.com/lib/pq"
	macaron "gopkg.in/macaron.v1"
)

func ModelInit() {
	createEngine("default")
}

func createEngine(key string) {

	var orm *xorm.Engine

	user := macaron.Config().Section("database." + key).Key("user").String()
	pass := macaron.Config().Section("database." + key).Key("pass").String()
	host := macaron.Config().Section("database." + key).Key("host").String()
	port := macaron.Config().Section("database." + key).Key("port").String()
	name := macaron.Config().Section("database." + key).Key("name").String()

	switch macaron.Config().Section("database." + key).Key("type").String() {
	case "postgres":
		link := fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=disable", user, pass, host, port, name)
		x, err := xorm.NewEngine("postgres", link)
		if err != nil {
			log.Println(fmt.Sprintf("postgres:%v connection create failed:%v", link, err))
			os.Exit(1)
		} else if err = x.Ping(); err != nil {
			log.Println(fmt.Sprintf("postgres:%v connection ping failed:%v", link, err))
			os.Exit(1)
		}
		orm = x
	case "mysql":
		link := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8", user, pass, host, port, name)
		x, err := xorm.NewEngine("mysql", link)
		if err != nil {
			log.Println(fmt.Sprintf("mysql:%v connection create failed:%v", link, err))
			os.Exit(1)
		} else if err = x.Ping(); err != nil {
			log.Println(fmt.Sprintf("mysql:%v connection ping failed:%v", link, err))
			os.Exit(1)
		}
		orm = x
	default:
		log.Println("Unknown database type")
		os.Exit(1)
	}

	//DB.SetLogger(xorm.NewSimpleLogger(logger.GetOrmLogHandle()))

	prefix := macaron.Config().Section("database." + key).Key("prefix").String()
	tbMapper := core.NewPrefixMapper(core.SnakeMapper{}, prefix)

	orm.SetTableMapper(tbMapper)

	switch macaron.Env {
	case macaron.TEST:
	case macaron.DEV:
		orm.ShowSQL(true)
		orm.Logger().SetLevel(core.LOG_DEBUG)
	case macaron.PROD:
		//orm.Logger().SetLevel(core.LOG_ERR)
		orm.ShowSQL(false)
	}

	maxIdleConns := macaron.Config().Section("database." + key).Key("MaxIdleConns").MustInt(10)
	maxOpenConns := macaron.Config().Section("database." + key).Key("MaxOpenConns").MustInt(100)

	orm.SetMaxIdleConns(maxIdleConns)
	orm.SetMaxOpenConns(maxOpenConns)

	engine.SetEngine(key, orm)
}
