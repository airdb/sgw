package checker

import (
	"log"
	"net"
	"strings"

	"github.com/ipipdotnet/ipdb-go"
)

const IPIPEN = "EN"
const IPIPCN = "CN"

// idc : "" | IDC |  VPN
const (
	IPIDC  = "IDC"
	IPVPN  = "VPN"
	IPNULL = ""
)

type Ipip struct {
	DB *ipdb.City
}

var IPIP Ipip

func NewIPIP() {
	var err error

	IPIP.DB, err = ipdb.NewCity("ipv4_en.ipdb")
	if err != nil {
		panic(err)
	}
}

func IsIPv6(str string) bool {
	ip := net.ParseIP(str)
	return ip != nil && strings.Contains(str, ":")
}

// https://github.com/ipipdotnet/ipdb-go
func CheckIP(ip string) bool {
	if IsIPv6(ip) {
		return false
	}

	db, err := ipdb.NewCity("ipv4_en.ipdb")
	if err != nil {
		log.Fatal(err)
	}

	dbmap, _ := db.FindMap(ip, IPIPEN)
	idc := dbmap["idc"]

	if idc == IPIDC || idc == IPVPN {
		return true
	}

	return false
}

func (ipip Ipip) CheckIP(ip string) bool {
	if IsIPv6(ip) {
		return false
	}

	dbmap, _ := ipip.DB.FindMap(ip, IPIPEN)
	idc := dbmap["idc"]

	if idc == IPIDC || idc == IPVPN {
		return true
	}

	return false
}

func GetIPInfo(ip string) *ipdb.CityInfo {
	if IsIPv6(ip) {
		return nil
	}

	db, err := ipdb.NewCity("ipv4_en.ipdb")
	if err != nil {
		log.Fatal(err)
	}

	info, err := db.FindInfo(ip, IPIPEN)
	if err != nil {
		log.Fatal(ip, err)
	}

	return info
}
