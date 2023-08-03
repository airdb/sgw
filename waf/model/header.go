package model

import "net/http"

type ResponseHeader struct {
	AddHeaders http.Header `json:"add_headers,omitempty"`
	DelHeaders []string    `json:"del_headers,omitempty"`
}
type RequestHeader struct {
	GeneralHeader GeneralHeader `json:"general_header"`
	ForwardHeader ForwardHeader `json:"forward_header,omitempty"`
	CommonHeader  CommonHeader  `json:"common_header"`
	CustomHeader  CustomHeader  `json:"gateway_custom_header"`
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
type CustomHeader struct {
	XJA3Fingerprint   string `json:"X-Ja3-Fingerprint"`   // JA3 fingerprint
	XHTTP2Fingerprint string `json:"X-Http2-Fingerprint"` // HTTP2 fingerprint
	XDeviceID         string `json:"X-Device-Id"`         // Device ID

	XSgwServer string `json:"X-Sgw-Server"` // Server name
	XSgwRid    string `json:"X-Sgw-Rid"`    // Request ID
	XSgwUid    string `json:"X-Sgw-Uid"`    // User ID
	XSgwAction string `json:"X-Sgw-Action"` // Action
	XSgwDebug  string `json:"X-Sgw-Debug"`  // Debug
}
