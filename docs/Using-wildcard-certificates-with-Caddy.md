https://www.moerats.com/archives/900/

Provider DNSPod:
https://github.com/caddy-dns/dnspod


Caddyserver 配置
https://dbohdan.com/wiki/caddy

check dns provider: 
```bash
./caddy list-modules | grep dns
```

Here's an example Caddyfile config that lets you use wildcard certificates with Caddy:
```
*.{$DOMAIN} {

	tls YourEmail {
        dns cloudflare {env.CLOUDFLARE_API_TOKEN}
        dns dnspod {env.CLOUDFLARE_API_TOKEN}
        }

	@serviceA host serviceA.{$DOMAIN}
	reverse_proxy @serviceA IP:port

	@serviceB host serviceB.{$DOMAIN}
	reverse_proxy @serviceB IP:port

	@serviceC host serviceC.{$DOMAIN}
	reverse_proxy @serviceC IP:port

}
```