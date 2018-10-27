package collect_sys_info


import (
	"github.com/gostudy03/logagent/common"
	"time"
	"sync"
	"github.com/gostudy03/xlog"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	_"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/mem"
	"github.com/shirou/gopsutil/net"
	"github.com/shirou/gopsutil/load"
	"encoding/json"
	"github.com/gostudy03/logagent/kafka"
)

const (
	SystemTypeCpu = "cpu"
	SystemTypeMem = "mem"
	SystemTypeDisk = "disk"
	SystemTypeNet = "net"
)

var (
	collectSystemInfoTopic string
	lastCollectUnixStamp int64
	lastNetInfo *NetInfo
)

type PartitionStat struct {
	Device     string `json:"device"`
	Mountpoint string `json:"mountpoint"`
	Fstype     string `json:"fstype"`
	Opts       string `json:"opts"`

	Total uint64 `json:"total"`
	Free uint64 `json:"free"`
	Used uint64 `json:"used"`
	UsedRate float64 `json:"used_rate"`

	InodesTotal       uint64  `json:"inodesTotal"`
	InodesUsed        uint64  `json:"inodesUsed"`
	InodesFree        uint64  `json:"inodesFree"`
	InodesUsedPercent float64 `json:"inodesUsedPercent"`
}

type DiskIO struct {
	ReadCount        uint64 `json:"readCount"`
	WriteCount       uint64 `json:"writeCount"`
	ReadBytes        uint64 `json:"readBytes"`
	WriteBytes       uint64 `json:"writeBytes"`
}

type DiskInfo struct {
	Partitions []PartitionStat	`json:"partitions"`
	DiskIOMap map[string]DiskIO `json:"disk_io"`
}

type CpuInfo struct {
	Percent float64 `json:"percent"`
	Load1 float64	`json:"load1"`
	Load5 float64	`json:"load5`
	Load15 float64	`json:"load15"`
	CpuCoreInfo []*CoreInfo `json:"cpu_core_info"`
}

type MemInfo struct {
	Total uint64 `json:"total"`
	Free uint64 `json:"free"`
	Used uint64 `json:"used"`
	Cache uint64 `json:"cache"`
	Buffer uint64 `json:"buffer"`
	UseRate float64 `json:"use_rate"`
}

type SystemInfo struct {
	Type string `json:"type"`
	IP string 		`json:"ip"`
	Data string `json:"data"`
}

type CoreInfo struct {
	CPU        int32    `json:"cpu"`
	VendorID   string   `json:"vendorId"`
	Family     string   `json:"family"`
	Model      string   `json:"model"`
	ModelName  string   `json:"modelName"`
	Mhz        float64  `json:"mhz"`
	CacheSize  int32    `json:"cacheSize"`
	CoresNum    int32   `json:"cores_num"`
}

type IOCountersStat struct {
	Name        string `json:"name"`        // interface name
	BytesSentRate   float64 `json:"bytesSent_rate"`   // number of bytes sent
	BytesRecvRate   float64 `json:"bytesRecv_rate"`   // number of bytes received
	PacketsSentRate float64 `json:"packetsSent_rate"` // number of packets sent
	PacketsRecvRate float64 `json:"packetsRecv_rate"` // number of packets received

	BytesSent   uint64 `json:"bytesSent"`   // number of bytes sent
	BytesRecv   uint64 `json:"bytesRecv"`   // number of bytes received
	PacketsSent uint64 `json:"packetsSent"` // number of packets sent
	PacketsRecv uint64 `json:"packetsRecv"` // number of packets received
	Errin       uint64 `json:"errin"`       // total number of errors while receiving
	Errout      uint64 `json:"errout"`      // total number of errors while sending
	Dropin      uint64 `json:"dropin"`      // total number of incoming packets which were dropped
	Dropout     uint64 `json:"dropout"`     // total number of outgoing packets which were dropped (always 0 on OSX and BSD)
	Fifoin      uint64 `json:"fifoin"`      // total number of FIFO buffers errors while receiving
	Fifoout     uint64 `json:"fifoout"`     // total number of FIFO buffers errors while sending
}

type NetInfo struct {
	NetInterfaces map[string]*IOCountersStat `json:"net_interfaces"`
}

func Run(wg *sync.WaitGroup, interval time.Duration, topic string) {

	collectSystemInfoTopic = topic
	timer := time.NewTicker(interval)
	for {
		select {
		case <- timer.C:
			doCollect()
		}
	}

	wg.Done()
}

func doCollectCpu() {
	
	cpuInfos, err := cpu.Info()
	xlog.LogDebug("cpu info:%#v", cpuInfos)

	percent, err:= cpu.Percent(time.Second, false)
	if err != nil {
		xlog.LogError("collect cpu percent failed, err:%v", err)
		return
	}

	info, _ := load.Avg()
	
	var cpuInfo CpuInfo
	cpuInfo.Percent = percent[0]
	cpuInfo.Load1 = info.Load1
	cpuInfo.Load5 = info.Load5
	cpuInfo.Load15 = info.Load15

	for _, info := range cpuInfos {
		var coreInfo CoreInfo
		coreInfo.CacheSize = info.CacheSize
		coreInfo.CPU = info.CPU
		coreInfo.Family = info.Family
		coreInfo.Mhz = info.Mhz
		coreInfo.Model = info.Model
		coreInfo.ModelName = info.ModelName
		coreInfo.VendorID = info.VendorID
		coreInfo.CoresNum = info.Cores

		cpuInfo.CpuCoreInfo = append(cpuInfo.CpuCoreInfo, &coreInfo)
	}

	
	xlog.LogDebug("collect cpu succ, info:%#v", cpuInfo)
	sendToKafka(SystemTypeCpu, &cpuInfo)
}

func sendToKafka(systemType string, data interface{}) {
		
	byteData, err := json.Marshal(data)
	if err != nil {
		xlog.LogError("marshal cpu info failed, err:%v", err)
		return
	}

	localIP, _ := common.GetLocalIP()
	var systemInfo SystemInfo
	systemInfo.Type = systemType
	systemInfo.IP = localIP
	systemInfo.Data = string(byteData)

	jsonData, err := json.Marshal(systemInfo)
	if err != nil {
		return
	}

	msg := &kafka.Message{
		Data: string(jsonData),
		Topic: collectSystemInfoTopic,
	}

	err = kafka.SendLog(msg)
	if err != nil {
		xlog.LogError("send to kafka failed, err:%v", err)
		return
	}
}

func doCollectMem()  {
	mInfo, _ := mem.VirtualMemory()
	var memInfo MemInfo

	memInfo.Free = mInfo.Available
	memInfo.Total = mInfo.Total
	memInfo.Used = mInfo.Used

	memInfo.Buffer = mInfo.Buffers
	memInfo.Cache = mInfo.Cached
	memInfo.UseRate = mInfo.UsedPercent

	sendToKafka(SystemTypeMem, &memInfo)
}

func doCollectDisk() {
	
	var diskInfo DiskInfo

	diskInfo.DiskIOMap = make(map[string]DiskIO, 16)

	part, _ := disk.Partitions(true)
	for _, v := range part {
		var partition PartitionStat
		partition.Device = v.Device
		partition.Fstype = v.Fstype
		partition.Mountpoint = v.Mountpoint
		partition.Opts = v.Opts

		dInfo, _ := disk.Usage(v.Device)
		partition.Free = dInfo.Free
		partition.Total = dInfo.Total
		partition.Used = dInfo.Used
		partition.UsedRate = dInfo.UsedPercent

		partition.InodesFree = dInfo.InodesFree
		partition.InodesTotal = dInfo.InodesTotal
		partition.InodesUsed = dInfo.InodesUsed
		partition.InodesUsedPercent = dInfo.InodesUsedPercent

		diskInfo.Partitions = append(diskInfo.Partitions, partition)
	}

	ioStat, err := disk.IOCounters()
	if err != nil {
		xlog.LogError("get disk io failed, err:%v", err)
	}
	for k, v := range ioStat {
		var diskIO DiskIO
		diskIO.ReadBytes = v.ReadBytes
		diskIO.ReadCount = v.ReadCount
		diskIO.WriteBytes = v.WriteBytes
		diskIO.WriteCount = v.WriteCount

		diskInfo.DiskIOMap[k] = diskIO
	}
	
	sendToKafka(SystemTypeDisk, &diskInfo)
}

func doCollectNet() {

	var netInfo NetInfo
	netInfo.NetInterfaces = make(map[string]*IOCountersStat, 16)

	curTimeStamp := time.Now().Unix()
	info, _ := net.IOCounters(true)
	for _, value := range info {
		var ioStat IOCountersStat
		ioStat.BytesRecv = value.BytesRecv
		ioStat.BytesSent = value.BytesSent
		ioStat.Dropin = value.Dropin
		ioStat.Dropout = value.Dropout
		ioStat.Errin = value.Errin
		ioStat.Errout = value.Errout
		ioStat.Fifoin = value.Fifoin
		ioStat.Fifoout = value.Fifoout
		ioStat.Name = value.Name
		ioStat.PacketsRecv = value.PacketsRecv
		ioStat.PacketsSent = value.PacketsSent

		netInfo.NetInterfaces[ioStat.Name] = &ioStat

		if lastCollectUnixStamp == 0 || lastNetInfo == nil{
			continue
		}

		interval := (curTimeStamp - lastCollectUnixStamp)
		lastInfo, ok := lastNetInfo.NetInterfaces[ioStat.Name]
		if !ok {
			continue
		}

		ioStat.BytesRecvRate = float64(ioStat.BytesRecv - lastInfo.BytesRecv) / float64(interval)
		ioStat.BytesSentRate = float64(ioStat.BytesSent - lastInfo.BytesSent) / float64(interval)
		ioStat.PacketsRecvRate = float64(ioStat.PacketsRecv - lastInfo.PacketsRecv) / float64(interval)
		ioStat.PacketsSentRate  = float64(ioStat.PacketsSent - lastInfo.PacketsSent) / float64(interval)
	}

	lastNetInfo = &netInfo
	lastCollectUnixStamp = curTimeStamp
	sendToKafka(SystemTypeNet, &netInfo)
}

func doCollect() {
	xlog.LogDebug("start collect system info")
	doCollectCpu()
	doCollectMem()
	doCollectDisk()
	doCollectNet()
}