Google Cloud DNS module for Caddy
===========================

This package contains a DNS provider module for [Caddy](https://github.com/caddyserver/caddy). It can be used to manage DNS records in Google Cloud DNS zones.

## Caddy module name

```
dns.providers.googleclouddns
```

## Authenticating

See [the associated README in the libdns package](https://github.com/libdns/googleclouddns) for important information about credentials.

## Building

To compile this Caddy module, follow the steps describe at the [Caddy Build from Source](https://github.com/caddyserver/caddy#build-from-source) instructions and import the `github.com/caddy-dns/googleclouddns` plugin

## Config examples

To use this module for the ACME DNS challenge, [configure the ACME issuer in your Caddy JSON](https://caddyserver.com/docs/json/apps/tls/automation/policies/issuer/acme/) like so:

```json
{
  "module": "acme",
  "challenges": {
    "dns": {
      "provider": {
        "name": "googleclouddns",
        "gcp_project": "{env.GCP_PROJECT}",
      }
    }
  }
}
```

or with the Caddyfile:

```
tls {
  dns googleclouddns {
    gcp_project {$GCP_PROJECT}
  }
}
```

You can replace `{$*}` or `{env.*}` with the actual values if you prefer to put it directly in your config instead of an environment variable.