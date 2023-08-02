package model

type SecureInfo struct {
	GeneralHeader  GeneralHeader  `json:"general_header"`
	RequestHeader  RequestHeader  `json:"request_header"`
	RequestArgs    RequestArgs    `json:"request_args"`
	RequestBody    RequestBody    `json:"request_body"`
	ResponseHeader ResponseHeader `json:"response_header"`
	ResponseBody   ResponseBody   `json:"response_body"`
	IPInfo         IPInfo         `json:"ip_info"`
	SecureAction   SecureAction   `json:"secure_action"`
}

type GeneralHeader struct{}

type RequestHeader struct {
	GatewayCustomHeader GatewayCustomHeader `json:"gateway_custom_header"`
}

type RequestArgs struct{}

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

// Custom header.
type GatewayCustomHeader struct {
	XSgwServer string `json:"X-Sgw-Server"` // Server name
	XSgwRid    string `json:"X-Sgw-Rid"`    // Request ID
	XSgwUid    string `json:"X-Sgw-Uid"`    // User ID
	XSgwAction string `json:"X-Sgw-Action"` // Action
}
