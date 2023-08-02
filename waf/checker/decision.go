package checker

import (
	"fmt"
	"net/http"

	"github.com/airdb/caddywaf/waf/model"
	"github.com/google/uuid"
)

const (
	ActionAllow   = "allow"
	ActionWatch   = "watch"
	ActionBlock   = "block"
	ActionDefault = ActionAllow
)

// Scroe is range from 0 to 100.
var CreditScore uint

// Block Message.
const (
	BlockByIP      = "block by IP"
	BlockByIPIDC   = "block by IP IDC"
	BlockByUA      = "block by UA"
	BlockByURI     = "block by URI"
	ActionWatchMsg = "watch"
)

func RunSecCheck(w http.ResponseWriter, r *http.Request) error {
	var err error
	var databus model.SecureInfo
	fmt.Println("waf middleware", err, databus)

	// fmt.Println("waf middleware", m.Orders)
	w.Header().Set("X-Sgw-Server", "airdb-secure-gateway")
	w.Header().Set("Server", "airdb-secure-gateway")

	requestID := uuid.New().String()
	w.Header().Add("X-Sgw-Rid", requestID)
	w.Header().Add("X-Sgw-Uid", "uid-test")
	w.Header().Add("X-Sgw-Action", "pass")

	// cip, _, _ := net.SplitHostPort(r.RemoteAddr)
	// ua := r.Header.Get("user-agent")

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

	return err
}
