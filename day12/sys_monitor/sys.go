package main

import (
	"fmt"
	"time"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/mem"
	"github.com/shirou/gopsutil/net"
	"github.com/shirou/gopsutil/load"
)

func cpuInfo()  {
	cpuInfos, err := cpu.Info()
	if err != nil {
		fmt.Printf("get cpu info failed, err:%v\n", err)
		return
	}
	for _, ci := range cpuInfos {
		fmt.Printf("cpu:%#v\n", ci)
	}

	for {
		percent, _ := cpu.Percent(time.Second, false)
		fmt.Printf("cpu percent:%v\n", percent)
	}
}

func memInfo() {
	mInfo, _ := mem.VirtualMemory()
	fmt.Printf("mem_info:%#v,uptime:%v\n", mInfo)
}

func hostInfo() {
	hInfo, _ := host.Info()
	fmt.Printf("hostInfo:%#v uptime:%v boottime:%v\n", hInfo, hInfo.Uptime, hInfo.BootTime)
}

func diskInfo() {
	dInfo, _ := disk.Usage("D:")
	fmt.Printf("disk:%#v total:%v free:%v used:%v\n", dInfo, dInfo.Total, dInfo.Free, dInfo.Used)

	part, _ := disk.Partitions(true)
	for _, v := range part {
		fmt.Printf("part: %#v\n", v)
	}

	ioStat, _ := disk.IOCounters()
	for k, v := range ioStat {
		fmt.Printf("%v: %v\n", k, v)
	}
}

func cpuLoad() {
	info, _ := load.Avg()
	fmt.Printf("load:%#v\n", info)
}

func netInfo() {
	info, _ := net.IOCounters(true)
	for index, value := range info {
		fmt.Printf("net%d, %#v, recv:%v send:%v\n", index, value, value.BytesRecv, value.BytesSent)
	}
}

func main() {
	//cpuInfo()
	//memInfo()
	//hostInfo()
	diskInfo()
	//cpuLoad()
	//netInfo()
	return
	v, _ := mem.VirtualMemory()
	c, _ := cpu.Info()
	cc, _ := cpu.Percent(time.Second, false)
	d, _ := disk.Usage("/")
	n, _ := host.Info()
	nv, _ := net.IOCounters(true)
	boottime, _ := host.BootTime()
	btime := time.Unix(int64(boottime), 0).Format("2006-01-02 15:04:05")

	fmt.Printf("        Mem       : %v MB  Free: %v MB Used:%v Usage:%f%%\n", v.Total/1024/1024, v.Available/1024/1024, v.Used/1024/1024, v.UsedPercent)
	if len(c) > 1 {
		for _, sub_cpu := range c {
			modelname := sub_cpu.ModelName
			cores := sub_cpu.Cores
			fmt.Printf("        CPU       : %v   %v cores \n", modelname, cores)
		}
	} else {
		sub_cpu := c[0]
		modelname := sub_cpu.ModelName
		cores := sub_cpu.Cores
		fmt.Printf("        CPU       : %v   %v cores \n", modelname, cores)

	}
	fmt.Printf("        Network: %v bytes / %v bytes\n", nv[0].BytesRecv, nv[0].BytesSent)
	fmt.Printf("        SystemBoot:%v\n", btime)
	fmt.Printf("        CPU Used    : used %f%% \n", cc[0])
	fmt.Printf("        HD        : %v GB  Free: %v GB Usage:%f%%\n", d.Total/1024/1024/1024, d.Free/1024/1024/1024, d.UsedPercent)
	fmt.Printf("        OS        : %v(%v)   %v  \n", n.Platform, n.PlatformFamily, n.PlatformVersion)
	fmt.Printf("        Hostname  : %v  \n", n.Hostname)
}
