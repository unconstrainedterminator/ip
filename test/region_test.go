package test

import (
	"github.com/unconstrainedterminator/ip"
	"log"
	"testing"
)

func Test_Country(t *testing.T) {
	var addrList []string
	addrList = ip.RandomIp()

	for _, addr := range addrList {
		country := ip.New(addr).Country()
		log.Println("Country:", country)
	}
}

func Test_Province(t *testing.T) {
	var addrList []string
	addrList = ip.RandomIp()

	for _, addr := range addrList {
		province := ip.New(addr).Province()
		log.Println("Province:", province)
	}
}

func Test_City(t *testing.T) {
	var addrList []string
	addrList = ip.RandomIp()

	for _, addr := range addrList {
		city := ip.New(addr).City()
		log.Println("City:", city)
	}
}

func Test_Region(t *testing.T) {
	var addrList []string
	addrList = ip.RandomIp()

	for _, addr := range addrList {
		region := ip.New(addr).Region()
		log.Println("Region:", region)
	}
}
