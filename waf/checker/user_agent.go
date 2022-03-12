package checker

import "strings"

var BlockUserAgentPrefixPatterns = []string{
	"curl/",
	"python",
}

var AllowUserAgentPrefixPatterns = []string{
	"Mozilla",
}

func CheckUserAgent(ua string) bool {
	for _, uaPrefix := range BlockUserAgentPrefixPatterns {
		if strings.HasPrefix(ua, uaPrefix) {
			return true
		}
	}

	return false
}
