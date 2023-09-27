package checker

import (
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

// https://github.com/porech/caddy-maxmind-geolocation
// https://github.com/mholt/caddy-ratelimit

type Ipip struct {
	DB *ipdb.City
}

var IPIP Ipip

func NewIPIP(ipVendor string) {
	var err error

	IPIP.DB, err = ipdb.NewCity(ipVendor)
	if err != nil {
		panic(err)
	}
}

func IsIPv6(str string) bool {
	ip := net.ParseIP(str)
	return ip != nil && strings.Contains(str, ":")
}

// https://github.com/ipipdotnet/ipdb-go
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
