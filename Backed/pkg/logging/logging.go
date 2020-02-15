package logging

import (
	"ITBat/pkg/setting"
	"github.com/sirupsen/logrus"
	"github.com/weekface/mgorus"
)

var Logger *logrus.Logger

func init() {
	Logger = logrus.New()
	hooker, err := mgorus.NewHookerWithAuthDb(setting.MongoCfg.HOST, "admin", setting.MongoCfg.DB, setting.MongoCfg.Conllection, setting.MongoCfg.User, setting.MongoCfg.Password)
	if err == nil {
		Logger.Hooks.Add(hooker)
	} else {
		Logger.Print(err)
	}
}

//系统出现的error日志记录
func DebugLog(v ...interface{}) {
	Logger.WithFields(logrus.Fields{
		"log_type": "model error",
		"error":    v,
	}).Error()
}
