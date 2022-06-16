package dnspod

import (
	"os"

	"github.com/caddyserver/caddy/v2"
	"github.com/caddyserver/caddy/v2/caddyconfig/caddyfile"
	"go.uber.org/zap"
	// dnspod "github.com/libdns/dnspod"
)

// Provider wraps the provider implementation as a Caddy module.
type Provider struct{ *DNSProvider }

const ModuleName = "dnspod"

var log *zap.Logger

func init() {
	caddy.RegisterModule(Provider{})
	log = caddy.Log().Named(ModuleName)

}

// CaddyModule returns the Caddy module information.
func (Provider) CaddyModule() caddy.ModuleInfo {
	return caddy.ModuleInfo{
		ID:  "dns.providers.dnspod",
		New: func() caddy.Module { return &Provider{new(DNSProvider)} },
	}
}

// Before using the provider config, resolve placeholders in the API token.
// Implements caddy.Provisioner.
func (p *Provider) Provision(ctx caddy.Context) error {
	repl := caddy.NewReplacer()
	p.DNSProvider.APIToken = repl.ReplaceAll(p.DNSProvider.APIToken, "")
	return nil
}

// UnmarshalCaddyfile sets up the DNS provider from Caddyfile tokens. Syntax:
//
// dnspod [<api_token>] {
//     api_token <api_token>
// }
//
func (p *Provider) UnmarshalCaddyfile(d *caddyfile.Dispenser) error {
	token, exist := os.LookupEnv("DNSPOD_TOKEN")
	if exist {
		p.DNSProvider.APIToken = token
	}

	log.Info("get token", zap.String("token", p.DNSProvider.APIToken))
	for d.Next() {
		if d.NextArg() {
			p.DNSProvider.APIToken = d.Val()
		}
		if d.NextArg() {
			return d.ArgErr()
		}

		for nesting := d.Nesting(); d.NextBlock(nesting); {
			switch d.Val() {
			case "api_token":
				if p.DNSProvider.APIToken != "" {
					return d.Err("API token already set")
				}
				p.DNSProvider.APIToken = d.Val()
				if d.NextArg() {
					return d.ArgErr()
				}
			default:
				return d.Errf("unrecognized subdirective '%s'", d.Val())
			}
		}
	}
	if p.DNSProvider.APIToken == "" {
		return d.Err("missing API token")
	}
	return nil
}

// Interface guards
var (
	_ caddyfile.Unmarshaler = (*Provider)(nil)
	_ caddy.Provisioner     = (*Provider)(nil)
)
