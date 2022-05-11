package main

import (
	caddycmd "github.com/caddyserver/caddy/v2/cmd"

	// plug in Caddy modules here
	_ "github.com/caddyserver/caddy/v2/modules/standard"
	_ "github.com/airdb/caddywaf"
	_ "github.com/caddy-dns/dnspod"
)

func main() {
	caddycmd.Main()
}
