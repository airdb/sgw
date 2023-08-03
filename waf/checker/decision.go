package checker

import (
	"encoding/json"
	"errors"
	"fmt"
	"net"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/airdb/caddywaf/waf/model"
	"github.com/google/uuid"
)

// Scroe is range from 0 to 100.
var CreditScore uint

const (
	BlockByIP      = "block by IP"
	BlockByIPIDC   = "block by IP IDC"
	BlockByUA      = "block by UA"
	BlockByURI     = "block by URI"
	ActionWatchMsg = "watch"
)

func RunSecCheck(w http.ResponseWriter, r *http.Request) error {
	var err error

	var info model.SecureInfo
	start := time.Now()

	// info.Timestamp = // start.UnixNano()
	info.Timestamp = float64(start.UnixMicro()) / 1000000

	// fmt.Println("waf middleware", m.Orders)
	w.Header().Set("Server", "secure-gateway")

	requestID := uuid.New().String()
	w.Header().Add("X-Sgw-Rid", requestID)
	w.Header().Add("X-Sgw-Uid", "uid-test")
	w.Header().Add("X-Sgw-Action", "pass")

	cip, _, _ := net.SplitHostPort(r.RemoteAddr)

	info.IPInfo.IP = cip

	dbmap, _ := IPIP.DB.FindMap(cip, IPIPEN)
	fmt.Println("dbmap", dbmap["idc"])

	info.IPInfo.UsageType = ""
	if dbmap["idc"] == "IDC" || dbmap["idc"] == "VPN" {
		info.IPInfo.UsageType = "DCH"
	}

	info.IPInfo.Region = dbmap["region_name"]
	info.IPInfo.ISP = dbmap["isp_domain"]
	info.IPInfo.Domain = dbmap["owner_domain"]
	info.IPInfo.City = dbmap["city_name"]
	info.IPInfo.Latitude = dbmap["latitude"]
	info.IPInfo.Longitude = dbmap["longitude"]

	// Handle commmon request header.
	for k, v := range r.Header {
		switch k {
		case "Sec-Fetch-Site":
			info.RequestHeader.CommonHeader.SecFetchSite = v[0]
		case "Sec-Fetch-Mode":
			info.RequestHeader.CommonHeader.SecFetchMode = v[0]
		case "Sec-Fetch-User":
			info.RequestHeader.CommonHeader.SecFetchUser = v[0]
		case "Sec-Fetch-Dest":
			info.RequestHeader.CommonHeader.SecFetchDest = v[0]
		case "Sec-Ch-Ua":
			info.RequestHeader.CommonHeader.SecChUa = v[0]
		case "Sec-Ch-Ua-Mobile":
			info.RequestHeader.CommonHeader.SecChUaMobile = v[0]
		case "Sec-Ch-Ua-Platform":
			info.RequestHeader.CommonHeader.SecChUaPlatform = strings.Trim(v[0], "\"")
		case "User-Agent":
			info.RequestHeader.CommonHeader.UserAgent = v[0]
		case "Accept":
			info.RequestHeader.CommonHeader.Accept = v[0]
		case "Accept-Encoding":
			info.RequestHeader.CommonHeader.AcceptEncoding = v[0]
		case "Accept-Language":
			info.RequestHeader.CommonHeader.AcceptLanguage = v[0]
		case "Referer":
			info.RequestHeader.CommonHeader.Referer = v[0]
		case "Connection":
			info.RequestHeader.CommonHeader.Connection = v[0]
		default:
			fmt.Println("Not-support-common-head", k, v)
		}
	}

	info.RequestHeader.GeneralHeader.ContentType = r.Header.Get("Content-Type")

	info.RequestHeader.GeneralHeader.Protocol = r.Proto
	info.RequestHeader.GeneralHeader.RemoteAddr = r.RemoteAddr

	contentLength, err := strconv.Atoi(r.Header.Get("Content-Length"))
	if err == nil {
		info.RequestHeader.GeneralHeader.ContentLength = contentLength
	}

	if r.TLS == nil {
		info.RequestHeader.GeneralHeader.Scheme = "HTTP"
	} else {
		info.RequestHeader.GeneralHeader.Scheme = "HTTPS"

		// refer: https://github.com/caddyserver/caddy/issues/4504
		info.RequestHeader.GatewayCustomHeader.XJA3Fingerprint = "771,4567-458-459,0-1-2,0,0"
	}

	info.RequestHeader.GeneralHeader.Host = r.Host
	info.RequestHeader.GeneralHeader.Endpoint = r.URL.Path
	info.RequestHeader.GeneralHeader.QueryPair = r.URL.RawQuery

	info.RequestHeader.GeneralHeader.Method = r.Method

	/* Construct response header. */
	respHeader := make(map[string]string)
	respHeader["X-Sgw-Server"] = "secure-gateway"

	/*

				check := checker.IPIP.CheckIP(cip)

				actionMsg = checker.ActionWatchMsg
				log.Info(actionMsg, zap.String("Sec-Ch-Ua-Platform", r.Header.Get("Sec-Ch-Ua-Platform")), zap.String("ua", ua))
				log.Info(actionMsg, zap.String("cip", cip), zap.Bool("is_idc", check), zap.String("ip", r.RequestURI))

				if check {
					if check {
						w.Write([]byte("server error 500\n"))
						actionMsg = checker.BlockByIPIDC
						w.Header().Set("X-Airdb-Action", actionMsg)
						log.Info(actionMsg, zap.String("ip", cip), zap.String("uri", r.RequestURI))
						return errors.New("500")
					}
				}

			check := checker.CheckUserAgent(ua)
			if check {
				actionMsg = checker.BlockByUA
				w.Header().Set("X-Airdb-Action", actionMsg)
				w.Write([]byte("server error 500\n"))
				log.Info(actionMsg, zap.String("ua", ua), zap.String("uri", r.RequestURI))
				return errors.New("500")
			}

		w.Header().Add("X-Airdb-Action", actionMsg)
		// w.Write([]byte("waf check pass\n"))
	*/

	// log.Info("waf check pass", zap.String("request_id", requestID), zap.String("uri", r.RequestURI))

	info.TimeCost = time.Since(start)
	data, err := json.Marshal(info)
	if err != nil {
		return err
	}

	fmt.Printf("info: %s\n", data)

	w.Write(data)

	err = errors.New("debug")

	return err
}
