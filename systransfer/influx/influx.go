package influx


import (
   "encoding/json"
   "time"

   "github.com/influxdata/influxdb/client/v2"
   "github.com/gostudy03/xlog"
   "github.com/gostudy03/logagent/collect_sys_info"
)

const (
   MyDB          = "sys_info"
   username      = "admin"
   password      = ""
   CPUMeasurement = "cpu_usage"
)

var (
	influxClient client.Client
	msgChan chan string
)

func Init(addr string) (err error){
	influxClient, err = client.NewHTTPClient(client.HTTPConfig{
		Addr:     addr,
		Username: username,
		Password: password,
	})
	if err != nil {
		xlog.LogError("init influx db failed, err:%v", err)
		return
	}

	msgChan = make(chan string, 10000)
	for i := 0; i < 8; i++ {
		go insertData()
	}
	return 
}

func AppendMsg(data string) {
	msgChan <- data
}

func insertData() {
	for data := range msgChan {
		var sysInfo = &collect_sys_info.SystemInfo{}
		err := json.Unmarshal([]byte(data), sysInfo)
		if err != nil {
			xlog.LogError("unmarshal json failed, err:%v", err)
			continue
		}

		procSystemInfo(sysInfo)		
	}
}

func procSystemInfo(sysInfo *collect_sys_info.SystemInfo) {
	switch sysInfo.Type {
	case "cpu":
		procCpu(sysInfo)
	}
}

func procCpu(sysInfo *collect_sys_info.SystemInfo) {
	var cpuInfo = &collect_sys_info.CpuInfo{}
	err := json.Unmarshal([]byte(sysInfo.Data), cpuInfo)
	if err != nil {
		xlog.LogError("unmarshal failed, err:%v", err)
		return
	}

	xlog.LogDebug("influx get cpu info succ, cpu:%#v", cpuInfo)

	bp, err := client.NewBatchPoints(client.BatchPointsConfig{
		Database:  MyDB,
		Precision: "s",
	})
	if err != nil {
		xlog.LogError("new batch points failed, err:%v", err)
		return
	}
 
	tags := map[string]string{"host": sysInfo.IP}
	fields := map[string]interface{}{
		"usage":   cpuInfo.Percent,
	}
 
	pt, err := client.NewPoint(
		CPUMeasurement,
		tags,
		fields,
		time.Now(),
	)
	if err != nil {
		xlog.LogError("new point failed, err:%v", err)
		return
	}

	bp.AddPoint(pt)
	if err := influxClient.Write(bp); err != nil {
		xlog.LogError("insert data to influxdb failed, err:%v", err)
		return
	}

	xlog.LogDebug("insert cpu data to influx db succ")
}
	