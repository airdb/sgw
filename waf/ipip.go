package waf

import (
	"log"

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

// https://github.com/ipipdotnet/ipdb-go
func CheckIP(ip string) bool {
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
