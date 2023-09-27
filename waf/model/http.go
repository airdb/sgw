package model

type SecureInfo struct {
	Timestamp      float64        `json:"timestamp"`
	TimeCost       int64          `json:"time_cost"`
	IPInfo         IPInfo         `json:"ip_info"`
	RequestHeader  RequestHeader  `json:"request_header"`
	RequestBody    RequestBody    `json:"request_body"`
	ResponseHeader ResponseHeader `json:"response_header"`
	ResponseBody   ResponseBody   `json:"response_body"`
	SecureAction   SecureAction   `json:"secure_action"`
}

type RequestBody struct{}

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
