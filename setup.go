package sgw

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/airdb/sgw/waf/checker"
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
}

// Middleware implements an HTTP handler that writes the
// visitor's IP address to a file or stream.
type Middleware struct {
	// The file or stream to write to. Can be "stdout"
	// or "stderr".
	Output         string   `json:"output,omitempty"`
	IPVendor       string   `json:"ipvendor,omitempty"`
	Orders         []string `json:"orders"`
	StrategyOrders []string `json:"strategy_orders"`
	w              io.Writer
}

var args Middleware

const ModuleName = "sgw"

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
	// init ip info.
	if args.IPVendor != "" {
		checker.NewIPIP(args.IPVendor)
	}

	if m.w == nil {
		return fmt.Errorf("no writer")
	}
	return nil
}

// ServeHTTP implements caddyhttp.MiddlewareHandler.
func (m Middleware) ServeHTTP(w http.ResponseWriter, r *http.Request, next caddyhttp.Handler) error {
	err := checker.RunSecCheck(w, r)
	if err != nil {
		fmt.Println("run security check middleware failed", err)
		return nil
	}

	return next.ServeHTTP(w, r)
}

// UnmarshalCaddyfile implements caddyfile.Unmarshaler, Syntax:
//
//	waf <name> [<option>]
func (m *Middleware) UnmarshalCaddyfile(d *caddyfile.Dispenser) error {
	for d.Next() {
		for nesting := d.Nesting(); d.NextBlock(nesting); {
			switch d.Val() {
			case "output":
				if d.NextArg() {
					m.Output = d.Val()
				}
			case "ipvendor":
				// TODO: check value
				if d.NextArg() {
					m.IPVendor = d.Val()
				}
			case "orders":
				m.Orders = d.RemainingArgs()
			case "strategyOrders":
				m.StrategyOrders = d.RemainingArgs()
			}
		}
		/*
			if !d.Args(&m.Output) {
				return d.ArgErr()
			}
		*/
	}

	return nil
}

// parseCaddyfile unmarshals tokens from h into a new Middleware.
func parseCaddyfile(h httpcaddyfile.Helper) (caddyhttp.MiddlewareHandler, error) {
	// var m Middleware
	err := args.UnmarshalCaddyfile(h.Dispenser)
	if err != nil {
		return nil, err
	}
	return args, nil
}

// Interface guards
var (
	_ caddy.Provisioner           = (*Middleware)(nil)
	_ caddy.Validator             = (*Middleware)(nil)
	_ caddyhttp.MiddlewareHandler = (*Middleware)(nil)
	_ caddyfile.Unmarshaler       = (*Middleware)(nil)
)
