package main

import (
	"fmt"
	"github.com/astaxie/beego/session"
	_ "github.com/astaxie/beego/session/redis"
)

var globalSessions *session.Manager

func init() {
	config := &session.ManagerConfig{}
	config.CookieName = "sid"
	config.Gclifetime = 3600
	config.ProviderConfig = "127.0.0.1:6379,100,"
	config.EnableSetCookie = true
	config.Maxlifetime = 30 * 24*3600
	config.CookieLifeTime = 30 * 24*3600

	

	var err error
	globalSessions, err = session.NewManager("redis", config)
	if err != nil {
		panic(fmt.Sprintf("init session failed, err:%v", err))
		return
	}
	go globalSessions.GC()
}
