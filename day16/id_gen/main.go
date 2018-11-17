package main

import (
	"github.com/sony/sonyflake"
	"net/http"
	"encoding/json"
	"fmt"
	"github.com/gostudy03/xlog"
)

const (
	ErrInternal = 1001
	ErrSucccess = 0
)

var (
	snowFlake *sonyflake.Sonyflake
)

func initSonyFlake() (err error) {
	settings := sonyflake.Settings{}
	settings.MachineID = func()(uint16, error) {
		return 1, nil
	}
	snowFlake = sonyflake.NewSonyflake(settings)
	return
}

type ResponseData struct {
	ErrNo  int `json:"errno"`
	Message string `json:"messsage"`
	Id uint64 `json:"id"`
}

func responseError(w http.ResponseWriter) {

	var responseData ResponseData
	responseData.ErrNo = ErrInternal
	responseData.Message = "生成id失败"

	data, _ := json.Marshal(&responseData)
	w.Write(data)
}

func responseSuccess(id uint64, w http.ResponseWriter) {
	
	var responseData ResponseData
	responseData.ErrNo = ErrSucccess
	responseData.Message = "success"
	responseData.Id = id

	data, _ := json.Marshal(&responseData)
	w.Write(data)
}

func idGen(w http.ResponseWriter, r* http.Request) {

	var err error
	var id uint64

	defer func() {
		xlog.LogInfo("idgen result:%v, id:%v, caller ip:%v, url:%v", err, id, r.RemoteAddr, r.RequestURI)
	}()

	id, err = snowFlake.NextID()
	if err != nil {
		responseError(w)
		return
	}
	responseSuccess(id, w)
}

func main() {
	err := initSonyFlake()
	if err != nil {
		fmt.Printf("init snoy flake failed, err:%v\n", err)
		return
	}

	http.HandleFunc("/id/gen", idGen)
	http.ListenAndServe(":9090", nil)
}
