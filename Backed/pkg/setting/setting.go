package setting

import (
	"gopkg.in/ini.v1"
	"log"
	"time"
)

type Server struct {
	RUNMODE      string
	HTTPPORT     int
	READTIMEOUT  time.Duration
	WRITETIMEOUT time.Duration
}

type App struct {
	PAGE_SIZE  int
	JWT_SECRET string
}

type Mysql struct {
	TYPE     string
	USER     string
	PASSWORD string
	HOST     string
	NAME     string
}

type Mongo struct {
	HOST        string
	DB          string
	Conllection string
	User        string
	Password    string
}

var (
	Cfg         *ini.File
	AppCfg      App
	ServerCfg   Server
	DatabaseCfg Mysql
	MongoCfg    Mongo
)

func init() {
	var err error
	Cfg, err = ini.Load("conf/config.ini")
	if err != nil {
		log.Fatal("cant't load config.ini.", err)
	}
	AppCfg = loadApp()
	ServerCfg = loadServer()
	DatabaseCfg = loadDatabase()
	MongoCfg = loadMongo()
}

func loadApp() App {
	PAGE_SIZE := Cfg.Section("app").Key("PAGE_SIZE").MustInt(15)
	JWT_SECRET := Cfg.Section("app").Key("JWT_SECRET").MustString("love$vesan")
	return App{PAGE_SIZE, JWT_SECRET}
}

func loadServer() Server {
	RUNMODE := Cfg.Section("").Key("RUN_MODE").MustString("debug")
	HTTPPORT := Cfg.Section("server").Key("HTTP_PORT").MustInt(8080)
	READTIMEOUT, _ := Cfg.Section("server").Key("READ_TIMEOUT").Duration()
	WRITETIMEOUT, _ := Cfg.Section("server").Key("WRITE_TIMEOUT").Duration()
	return Server{RUNMODE, HTTPPORT, READTIMEOUT, WRITETIMEOUT}
}

func loadDatabase() Mysql {
	TYPE := Cfg.Section("database").Key("TYPE").MustString("mysql")
	USER := Cfg.Section("database").Key("USER").MustString("root")
	PASSWORD := Cfg.Section("database").Key("PASSWORD").MustString("root")
	HOST := Cfg.Section("database").Key("HOST").MustString("127.0.0.1:3306")
	NAME := Cfg.Section("database").Key("NAME").MustString("itbat")
	return Mysql{TYPE, USER, PASSWORD, HOST, NAME}
}

func loadMongo() Mongo {
	HOST := Cfg.Section("mongo").Key("Host").MustString("mongodb://localhost:27017")
	DB := Cfg.Section("mongo").Key("DB").MustString("itbatlogs")
	Conllection := Cfg.Section("mongo").Key("Conllection").MustString("logs")
	User := Cfg.Section("mongo").Key("User").MustString("root")
	PWD := Cfg.Section("mongo").Key("Password").MustString("root")
	return Mongo{HOST, DB, Conllection, User, PWD}
}
