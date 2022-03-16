package caddywaf

import (
	"errors"
	"fmt"
	"io"
	"net"
	"net/http"
	"os"

	"github.com/airdb/caddywaf/waf/checker"
	"github.com/caddyserver/caddy/v2"
	"github.com/caddyserver/caddy/v2/caddyconfig/caddyfile"
	"github.com/caddyserver/caddy/v2/caddyconfig/httpcaddyfile"
	"github.com/caddyserver/caddy/v2/modules/caddyhttp"
	"go.uber.org/zap"
)

var log *zap.Logger

func init() {
	caddy.RegisterModule(Middleware{})
	httpcaddyfile.RegisterHandlerDirective("waf", parseCaddyfile)

	// init log.
	log = caddy.Log().Named(ModuleName)

	// init ip info.

	checker.NewIPIP()
}

// Middleware implements an HTTP handler that writes the
// visitor's IP address to a file or stream.
type Middleware struct {
	// The file or stream to write to. Can be "stdout"
	// or "stderr".
	Output string `json:"output,omitempty"`

	w io.Writer
}

const ModuleName = "caddywaf"

// CaddyModule returns the Caddy module information.
func (Middleware) CaddyModule() caddy.ModuleInfo {
	return caddy.ModuleInfo{
		ID:  "http.handlers.waf",
		New: func() caddy.Module { return new(Middleware) },
	}
}

// Provision implements caddy.Provisioner.
func (m *Middleware) Provision(ctx caddy.Context) error {
	switch m.Output {
	case "stdout":
		m.w = os.Stdout
	case "stderr":
		m.w = os.Stderr
	default:
		return fmt.Errorf("an output stream is required")
	}
	return nil
}

// Validate implements caddy.Validator.
func (m *Middleware) Validate() error {
	if m.w == nil {
		return fmt.Errorf("no writer")
	}
	return nil
}

// ServeHTTP implements caddyhttp.MiddlewareHandler.
func (m Middleware) ServeHTTP(w http.ResponseWriter, r *http.Request, next caddyhttp.Handler) error {
	cip, _, _ := net.SplitHostPort(r.RemoteAddr)

	check := checker.IPIP.CheckIP(cip)
	actionMsg := ""
	ua := r.Header.Get("user-agent")

	actionMsg = checker.ActionWatchMsg
	log.Info(actionMsg, zap.String("Sec-Ch-Ua-Platform", r.Header.Get("Sec-Ch-Ua-Platform")), zap.String("ua", ua))
	log.Info(actionMsg, zap.String("cip", cip), zap.Bool("is_idc", check), zap.String("ip", r.RequestURI))

	if check {
		if check {
			w.Write([]byte("server error 500\n"))
			actionMsg = checker.BlockByIPIDC
			log.Info(actionMsg, zap.String("ip", cip), zap.String("uri", r.RequestURI))
			return errors.New("500")
		}
	}

	check = checker.CheckUserAgent(ua)
	if check {
		w.Write([]byte("server error 500\n"))
		actionMsg = checker.BlockByUA
		log.Info(actionMsg, zap.String("ua", ua), zap.String("uri", r.RequestURI))
		return errors.New("500")
	}
	// w.Write([]byte("waf check pass\n"))

	return next.ServeHTTP(w, r)
}

// UnmarshalCaddyfile implements caddyfile.Unmarshaler.
func (m *Middleware) UnmarshalCaddyfile(d *caddyfile.Dispenser) error {
	for d.Next() {
		if !d.Args(&m.Output) {
			return d.ArgErr()
		}
	}
	return nil
}

// parseCaddyfile unmarshals tokens from h into a new Middleware.
func parseCaddyfile(h httpcaddyfile.Helper) (caddyhttp.MiddlewareHandler, error) {
	var m Middleware
	err := m.UnmarshalCaddyfile(h.Dispenser)
	return m, err
}

// Interface guards
var (
	_ caddy.Provisioner           = (*Middleware)(nil)
	_ caddy.Validator             = (*Middleware)(nil)
	_ caddyhttp.MiddlewareHandler = (*Middleware)(nil)
	_ caddyfile.Unmarshaler       = (*Middleware)(nil)
)
