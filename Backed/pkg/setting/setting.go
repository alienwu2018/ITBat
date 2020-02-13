package setting

import (
	"gopkg.in/ini.v1"
	"log"
	"os"
	"time"
)

var (
	Cfg *ini.File

	RUNMODE string

	//server
	HTTPPORT     int
	READTIMEOUT  time.Duration
	WRITETIMEOUT time.Duration

	//app
	PAGE_SIZE  int
	JWT_SECRET string

	//database
	TYPE     string
	USER     string
	PASSWORD string
	HOST     string
	NAME     string

	//log
	SysLog_FILE_PATH   string
	InnerLog_FILE_PATH string
	SysLog_FILE_DIR    string
	InnerLog_FILE_DIR  string
)

func init() {
	var err error
	Cfg, err = ini.Load("conf/config.ini")
	if err != nil {
		log.Fatal("cant't load config.ini.", err)
	}
	LoadBase()
	LoadApp()
	LoadServer()
	LoadDatabase()
	LoadLog()
	//检测日志目录是否存在
	_, err = os.Stat(SysLog_FILE_DIR)
	if os.IsNotExist(err) {
		mkdir(SysLog_FILE_DIR)
	}
	_, err = os.Stat(InnerLog_FILE_DIR)
	if os.IsNotExist(err) {
		mkdir(InnerLog_FILE_DIR)
	}
}

func mkdir(path string) {
	dir, _ := os.Getwd()
	err := os.MkdirAll(dir+"/"+path, os.ModePerm)
	if err != nil {
		panic(err)
	}
}

func LoadBase() {
	RUNMODE = Cfg.Section("").Key("RUN_MODE").MustString("debug")
}

func LoadApp() {
	PAGE_SIZE = Cfg.Section("app").Key("PAGE_SIZE").MustInt(15)
	JWT_SECRET = Cfg.Section("app").Key("JWT_SECRET").MustString("love$vesan")
}

func LoadServer() {
	HTTPPORT = Cfg.Section("server").Key("HTTP_PORT").MustInt(8080)
	READTIMEOUT, _ = Cfg.Section("server").Key("READ_TIMEOUT").Duration()
	WRITETIMEOUT, _ = Cfg.Section("server").Key("WRITE_TIMEOUT").Duration()
}

func LoadDatabase() {
	TYPE = Cfg.Section("database").Key("TYPE").MustString("mysql")
	USER = Cfg.Section("database").Key("USER").MustString("root")
	PASSWORD = Cfg.Section("database").Key("PASSWORD").MustString("root")
	HOST = Cfg.Section("database").Key("HOST").MustString("127.0.0.1:3306")
	NAME = Cfg.Section("database").Key("NAME").MustString("itbat")

}

func LoadLog() {
	SysLog_FILE_DIR = Cfg.Section("log").Key("sysLog_FILE_DIR").MustString("runtime/logs/SystemLogs")
	InnerLog_FILE_DIR = Cfg.Section("log").Key("innerLog_FILE_DIR").MustString("runtime/logs/ErrorLogs")
	SysLog_FILE_PATH = Cfg.Section("log").Key("sysLog_FILE_PATH").MustString("runtime/logs/SystemLogs/syslog")
	InnerLog_FILE_PATH = Cfg.Section("log").Key("innerLog_FILE_PATH").MustString("runtime/logs/ErrorLogs/errlog")
}
