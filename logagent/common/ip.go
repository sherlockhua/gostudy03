package common

import (
	"net"
	"github.com/gostudy03/xlog"
)


var (
	localIP string
)

func GetLocalIP() (ip string, err error) {
	if len(localIP) > 0 {
		ip = localIP
		return
	}

	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return
	}

	for _, addr := range addrs {
		ipAddr, ok := addr.(*net.IPNet)
		if !ok {
			continue
		}

		if ipAddr.IP.IsLoopback() {
			continue
		}

		if !ipAddr.IP.IsGlobalUnicast() {
			continue
		}

		xlog.LogDebug("get local ip:%#v\n", ipAddr.IP.String())
		localIP = ipAddr.IP.String()
		ip = localIP
		return
	}
	return
}