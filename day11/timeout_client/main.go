package main

import (
	"sync"
	"io/ioutil"
	"time"
	"context"
	"net/http"
	"fmt"
)

type RespData struct {
	err error
	resp *http.Response
}

func doCall(ctx context.Context) {
	transport := http.Transport{}
	client := http.Client{
		Transport:&transport,
	}

	respChan := make(chan *RespData, 1)
	req, err := http.NewRequest("GET", "http://localhost:10000/", nil)
	if err != nil {
		fmt.Printf("new request failed, err:%v\n", err)
		return
	}

	var wg sync.WaitGroup
	wg.Add(1)
	defer wg.Wait()
	go func() {
		
		resp, err := client.Do(req)
		fmt.Printf("client.do resp:%v, err:%v\n", resp, err)
		respData := &RespData{
			resp:resp,
			err:err,
		}
		respChan <- respData
		wg.Done()
	}()

	select {
	case <- ctx.Done():
		transport.CancelRequest(req)
		fmt.Printf("call api timeout\n")
	case result := <- respChan:
		fmt.Printf("call server api succ\n")
		if (result.err != nil) {
			fmt.Printf("call api failed, err:%v\n", err)
			return
		}
		
		data, _ := ioutil.ReadAll(result.resp.Body)
		fmt.Printf("resp:%v\n", string(data))
	}

}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*100)
	defer cancel()
	doCall(ctx)
}