package model

type SecureAction struct {
	AllowStrategy Strategy `json:"strategy"`      // AllowStrategy first
	DenyStrategy  Strategy `json:"deny_strategy"` // Deny Strategy last
	HitStrategy   []int64  `json:"hit_strategy"`
	Action        string   `json:"action"`
}

type Strategy struct {
	Rules []Rule `json:"rules"`
}

type Rule struct {
	ID        int64  `json:"id"`
	MatchType string `json:"match_type"`
	LeftVal   string `json:"left_val"`
	RightVal  string `json:"right_val"`
	Action    string `json:"action"`
}
