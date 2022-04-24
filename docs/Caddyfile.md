docker-compose.yaml

```version: '3'

services: 
    # caddy
    caddy:
        image: caddy:latest
        ports:
            - "8848:80"
        volumes:
            - ${CONFIG_PATH}:/etc/caddy/Caddyfile
        environment:
            - DOMAIN
            - BACKEND
        logging:
            driver: "json-file"
            options:
                max-size: "1m"
                max-file: "10"
```


Caddyfile
```
http://{$DOMAIN} {
    @options {
        method OPTIONS
    }
    header {
        Access-Control-Allow-Origin *
        Access-Control-Allow-Credentials true
        Access-Control-Allow-Methods *
        Access-Control-Allow-Headers *
        defer
    }
    reverse_proxy shuwashuwa:8848 {
        header_down -Access-Control-Allow-Origin
    }
    respond @options 204
}
```


```
*.airdb.net {
	tls {
		#dns dnspod {env.DNSPOD_TOKEN}
		dns dnspod 307575,dec11869aee1bada71d4bac33f9b750c
	}

	encode {
		#zstd
		gzip
	}

	header /* {
		-Server
		Access-Control-Allow-Origin *
		x-airdb-server: airdb-gateway
		x-airdb-action: watch
    }
	@not_static {
		not {
			path /BingSiteAuth.xml /robots.txt
			path /favicon.ico
		}
	}

	handle /robots.txt {
		root * /srv/www
		file_server
	}

	handle /favicon.ico {
		root * /srv/www
		file_server
	}
	#root /robots.txt /srv/www
	#file_server

	# https://caddyserver.com/docs/caddyfile/directives/reverse_proxy
	@service host dev.airdb.net
	reverse_proxy @service http://127.0.0.1:7000

	@serviceA host a.airdb.net
	reverse_proxy @serviceA http://127.0.0.1:8000

	@serviceB host b.airdb.net
	reverse_proxy @serviceB http://115.159.93.223:8888

	root * /srv/www
	file_server
}
```