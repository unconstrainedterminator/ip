package ip

import (
	"math/rand"
	"net"
)

func GenerateIp() string {
	addrList := make(net.IP, 4)
	for i := range addrList {
		addrList[i] = byte(rand.Intn(256))
	}
	return addrList.String()
}

func RandomIp(num ...int) []string {
	maxNum := 100
	if len(num) > 0 {
		maxNum = num[0]
	}
	ipList := make([]string, maxNum)

	for i := 0; i < maxNum; i++ {
		ipList[i] = GenerateIp()
	}
	return ipList
}
