package model

type SecureInfo struct {
	GeneralHeader  GeneralHeader  `json:"general_header"`
	RequestHeader  RequestHeader  `json:"request_header"`
	RequestArgs    RequestArgs    `json:"request_args"`
	RequestBody    RequestBody    `json:"request_body"`
	ResponseHeader ResponseHeader `json:"response_header"`
	ResponseBody   ResponseBody   `json:"response_body"`
	IPInfo         IPInfo         `json:"ip_info"`
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

type IPInfo struct{}

type GatewayCustomHeader struct {
	XSgwServer string `json:"X-Sgw-Server"` // Server name
	XSgwRid    string `json:"X-Sgw-Rid"`    // Request ID
	XSgwUid    string `json:"X-Sgw-Uid"`    // User ID
	XSgwAction string `json:"X-Sgw-Action"` // Action
}
