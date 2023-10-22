package ipmod

// Ref: github.com/ip2location/ip2proxy-go/
import (
	"fmt"

	"github.com/ip2location/ip2proxy-go/v4"
)

func NewIP2Proxy(file string) {
	db, err := ip2proxy.OpenDB("./IP2PROXY-IP-PROXYTYPE-COUNTRY-REGION-CITY-ISP-DOMAIN-USAGETYPE-ASN-LASTSEEN-THREAT-RESIDENTIAL-PROVIDER.BIN")

	if err != nil {
		return
	}
	ip := "199.83.103.79"
	all, err := db.GetAll(ip)
	fmt.Println(all)
}
