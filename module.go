package googleclouddns

import (
	"github.com/caddyserver/caddy/v2"
	"github.com/caddyserver/caddy/v2/caddyconfig/caddyfile"
	libgoogleclouddns "github.com/libdns/googleclouddns"
)

// Provider lets Caddy read and manipulate DNS records hosted by this DNS provider.
type Provider struct{ *libgoogleclouddns.Provider }

func init() {
	caddy.RegisterModule(Provider{})
}

// CaddyModule returns the Caddy module information.
func (Provider) CaddyModule() caddy.ModuleInfo {
	return caddy.ModuleInfo{
		ID:  "dns.providers.googleclouddns",
		New: func() caddy.Module { return &Provider{new(libgoogleclouddns.Provider)} },
	}
}

// Provision sets up the module. Implements caddy.Provisioner.
func (p *Provider) Provision(ctx caddy.Context) error {
	repl := caddy.NewReplacer()
	p.Provider.Project = repl.ReplaceAll(p.Provider.Project, "")
	p.Provider.ServiceAccountJSON = repl.ReplaceAll(p.Provider.ServiceAccountJSON, "")
	return nil
}

// UnmarshalCaddyfile sets up the DNS provider from Caddyfile tokens. Syntax:
//
// googleclouddns {
//     gcp_project <project ID>
//     gcp_application_default <path to service account JSON (optional)>
// }
//
func (p *Provider) UnmarshalCaddyfile(d *caddyfile.Dispenser) error {
	for d.Next() {
		if d.NextArg() {
			return d.ArgErr()
		}
		for nesting := d.Nesting(); d.NextBlock(nesting); {
			switch d.Val() {
			case "gcp_project":
				if d.NextArg() {
					p.Provider.Project = d.Val()
				}
				if d.NextArg() {
					return d.ArgErr()
				}
			case "gcp_application_default":
				if d.NextArg() {
					p.Provider.ServiceAccountJSON = d.Val()
				}
				if d.NextArg() {
					return d.ArgErr()
				}
			default:
				return d.Errf("unrecognized subdirective '%s'", d.Val())
			}
		}
	}
	if p.Provider.Project == "" {
		return d.Err("missing Google Cloud project ID")
	}
	return nil
}

// Interface guards
var (
	_ caddyfile.Unmarshaler = (*Provider)(nil)
	_ caddy.Provisioner     = (*Provider)(nil)
)
