package model

import "time"

type SecureInfo struct {
	Timestamp      int64          `json:"timestamp"`
	TimeCost       time.Duration  `json:"time_cost"`
	IPInfo         IPInfo         `json:"ip_info"`
	RequestHeader  RequestHeader  `json:"request_header"`
	RequestArgs    RequestArgs    `json:"request_args"`
	RequestBody    RequestBody    `json:"request_body"`
	ResponseHeader ResponseHeader `json:"response_header"`
	ResponseBody   ResponseBody   `json:"response_body"`
	SecureAction   SecureAction   `json:"secure_action"`
}

type GeneralHeader struct {
	Scheme                  string `json:"scheme"`
	Protocol                string `json:"protocol"`
	RemoteAddr              string `json:"remote_addr"`
	UpgradeInsecureRequests string `json:"upgrade-insecure-requests"`
	TLSFingerprint          string `json:"tls_fingerprint"`
}

type RequestHeader struct {
	GeneralHeader       GeneralHeader       `json:"general_header"`
	CommonHeader        CommonHeader        `json:"common_header"`
	GatewayCustomHeader GatewayCustomHeader `json:"gateway_custom_header"`
}

type RequestArgs struct {
	QueryArgs string `json:"query_args"`
}

type RequestBody struct{}

type ResponseHeader struct {
	GatewayCustomHeader GatewayCustomHeader `json:"gateway_custom_header"`
}

type ResponseBody struct{}

type SecureAction struct {
	Strategy string `json:"strategy"`
	Action   string `json:"action"`
}

type IPInfo struct {
	IP        string `json:"ip"`
	ISP       string `json:"isp"`
	UsageType string `json:"usage_type"`
	Region    string `json:"region"`
	City      string `json:"city"`
	Longitude string `json:"longitude"`
	Latitude  string `json:"latitude"`
}

type CommonHeader struct {
	Host          string `json:"host"`
	Method        string `json:"method"`
	StatusCode    int    `json:"status_code"`
	ContentType   string `json:"content-type"`
	ContentLength int    `json:"content-length"`

	UserAgent      string `json:"user-agent"`
	Connection     string `json:"connection"`
	Referer        string `json:"referer"`
	Accept         string `json:"accept"`
	AcceptEncoding string `json:"accept-encoding"`
	AcceptLanguage string `json:"accept-language"`

	// For secure header.
	SecFetchSite    string `json:"sec-fetch-site"`
	SecFetchMode    string `json:"sec-fetch-mode"`
	SecFetchUser    string `json:"sec-fetch-user"`
	SecFetchDest    string `json:"sec-fetch-dest"`
	SecChUa         string `json:"sec-ch-ua"`
	SecChUaMobile   string `json:"sec-ch-ua-mobile"`
	SecChUaPlatform string `json:"sec-ch-ua-platform"`
}

// Custom header.
type GatewayCustomHeader struct {
	XSgwServer string `json:"X-Sgw-Server"` // Server name
	XSgwRid    string `json:"X-Sgw-Rid"`    // Request ID
	XSgwUid    string `json:"X-Sgw-Uid"`    // User ID
	XSgwAction string `json:"X-Sgw-Action"` // Action
	XSgwDebug  string `json:"X-Sgw-Debug"`  // Debug
}
