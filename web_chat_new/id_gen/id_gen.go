package id_gen

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type ResponseData struct {
	ErrNo   int    `json:"errno"`
	Message string `json:"messsage"`
	Id      uint64
}

func GetId(url string) (id uint64, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	var result ResponseData
	err = json.Unmarshal(data, &result)
	if err != nil {
		return
	}
	if result.ErrNo != 0 {
		err = fmt.Errorf("get id failed, err:%v, message:%v", result.ErrNo, result.Message)
		return
	}

	id = result.Id
	return
}
