package ipmod

// Ref: https://github.com/ip2location/ip2location-go

import (
	"fmt"

	"github.com/ip2location/ip2location-go/v9"
)

func NewIP2Location(file string) {
	fmt.Println("not implement")
	db, err := ip2location.OpenDB(file)
	if err != nil {
		fmt.Print(err)
		return
	}

	ip := "8.8.8.8"
	results, err := db.Get_all(ip)

	if err != nil {
		fmt.Print(err)
		return
	}

	fmt.Printf("country_short: %s\n", results.Country_short)
	fmt.Printf("country_long: %s\n", results.Country_long)
	fmt.Printf("region: %s\n", results.Region)
	fmt.Printf("city: %s\n", results.City)
	fmt.Printf("isp: %s\n", results.Isp)
	fmt.Printf("latitude: %f\n", results.Latitude)
	fmt.Printf("longitude: %f\n", results.Longitude)
	fmt.Printf("domain: %s\n", results.Domain)
	fmt.Printf("zipcode: %s\n", results.Zipcode)
	fmt.Printf("timezone: %s\n", results.Timezone)
	fmt.Printf("netspeed: %s\n", results.Netspeed)
	fmt.Printf("iddcode: %s\n", results.Iddcode)
	fmt.Printf("areacode: %s\n", results.Areacode)
	fmt.Printf("weatherstationcode: %s\n", results.Weatherstationcode)
	fmt.Printf("weatherstationname: %s\n", results.Weatherstationname)
	fmt.Printf("mcc: %s\n", results.Mcc)
	fmt.Printf("mnc: %s\n", results.Mnc)
	fmt.Printf("mobilebrand: %s\n", results.Mobilebrand)
	fmt.Printf("elevation: %f\n", results.Elevation)
	fmt.Printf("usagetype: %s\n", results.Usagetype)
	fmt.Printf("addresstype: %s\n", results.Addresstype)
	fmt.Printf("category: %s\n", results.Category)
	fmt.Printf("district: %s\n", results.District)
	fmt.Printf("asn: %s\n", results.Asn)
	fmt.Printf("as: %s\n", results.As)
	fmt.Printf("api version: %s\n", ip2location.Api_version())
}
