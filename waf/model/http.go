package model

import (
	"net/http"
	"time"
)

type SecureInfo struct {
	Timestamp      float64        `json:"timestamp"`
	TimeCost       time.Duration  `json:"time_cost"`
	IPInfo         IPInfo         `json:"ip_info"`
	RequestHeader  RequestHeader  `json:"request_header"`
	RequestBody    RequestBody    `json:"request_body"`
	ResponseHeader ResponseHeader `json:"response_header"`
	ResponseBody   ResponseBody   `json:"response_body"`
	SecureAction   SecureAction   `json:"secure_action"`
}

type GeneralHeader struct {
	Scheme    string `json:"scheme"`
	Host      string `json:"host"`
	Endpoint  string `json:"endpoint"`
	QueryPair string `json:"query_pair"`

	Method        string `json:"method"`
	StatusCode    int    `json:"status_code"`
	ContentType   string `json:"content-type"`
	ContentLength int    `json:"content-length"`

	Protocol                string `json:"protocol"`
	RemoteAddr              string `json:"remote_addr"`
	UpgradeInsecureRequests string `json:"upgrade-insecure-requests"`
}

type RequestHeader struct {
	GeneralHeader       GeneralHeader       `json:"general_header"`
	ForwardHeader       ForwardHeader       `json:"forward_header,omitempty"`
	CommonHeader        CommonHeader        `json:"common_header"`
	GatewayCustomHeader GatewayCustomHeader `json:"gateway_custom_header"`
}

type RequestBody struct{}

type ResponseHeader struct {
	AddHeaders http.Header `json:"add_headers,omitempty"`
	DelHeaders []string    `json:"del_headers,omitempty"`
}

type ResponseBody struct{}

type IPInfo struct {
	IP        string `json:"ip"`
	ISP       string `json:"isp"`
	Domain    string `json:"domain"`
	UsageType string `json:"usage_type"`
	Region    string `json:"region"`
	City      string `json:"city"`
	Longitude string `json:"longitude"`
	Latitude  string `json:"latitude"`
}

type CommonHeader struct {
	UserAgent      string `json:"user-agent"`
	Connection     string `json:"connection"`
	Referer        string `json:"referer"`
	Accept         string `json:"accept"`
	AcceptEncoding string `json:"accept-encoding"`
	AcceptLanguage string `json:"accept-language"`

	// For secure header.
	SecFetchSite  string `json:"sec-fetch-site"`
	SecFetchMode  string `json:"sec-fetch-mode"`
	SecFetchUser  string `json:"sec-fetch-user"`
	SecFetchDest  string `json:"sec-fetch-dest"`
	SecChUa       string `json:"sec-ch-ua"`
	SecChUaMobile string `json:"sec-ch-ua-mobile"`

	// refer: https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Sec-CH-UA-Platform
	// One of the following strings: "Android", "Chrome OS", "Chromium OS", "iOS", "Linux", "macOS", "Windows", or "Unknown".
	SecChUaPlatform string `json:"sec-ch-ua-platform"`
}

// Forward header.
type ForwardHeader struct {
	XForwardedFor   string `json:"x-forwarded-for,omitempty"`
	XRealIP         string `json:"x-real-ip,omitempty"`
	XForwardedProto string `json:"x-forwarded-proto,omitempty"`
	XForwardedHost  string `json:"x-forwarded-host,omitempty"`
	XForwardedPort  string `json:"x-forwarded-port,omitempty"`
	XForwardedSsl   string `json:"x-forwarded-ssl,omitempty"`
	XUrlScheme      string `json:"x-url-scheme,omitempty"`
}

// Custom header.
type GatewayCustomHeader struct {
	XJA3Fingerprint   string `json:"X-Ja3-Fingerprint"`   // JA3 fingerprint
	XHTTP2Fingerprint string `json:"X-Http2-Fingerprint"` // HTTP2 fingerprint
	XDeviceID         string `json:"X-Device-Id"`         // Device ID

	XSgwServer string `json:"X-Sgw-Server"` // Server name
	XSgwRid    string `json:"X-Sgw-Rid"`    // Request ID
	XSgwUid    string `json:"X-Sgw-Uid"`    // User ID
	XSgwAction string `json:"X-Sgw-Action"` // Action
	XSgwDebug  string `json:"X-Sgw-Debug"`  // Debug
}
