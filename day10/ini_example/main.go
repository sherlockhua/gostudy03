package main

import (
	"github.com/gostudy03/oconfig"
	"fmt"
)

type Config struct {

	CartDb DbConf   `ini:"cartdb"`
	Server ServerConf `ini:"server"`
}

type ServerConf struct {
	Host string   `ini:"host"`
	Port int		`ini:"port"`
}

type DbConf struct {
	User string `ini:"user"`
	Password string `ini:"password"`
	Host string `ini:"host"`
	Port int  `ini:"port"`
	Database string `ini:"database"`
	Rate float32 `ini:"rate"`
}

func main() {
	var conf Config
	filename := "./example.ini"
	err := oconfig.UnMarshalFile(filename, &conf)
	if err != nil {
		fmt.Printf("unmarshal file failed, err:%v\n", err)
		return
	}

	fmt.Printf("conf:%#v\n", conf)

	oconfig.MarshalFile("c:/tmp/test.ini", conf)
}