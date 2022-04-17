dev: build
	 ./caddy run -watch -config Caddyfile.dev

run:
	 ./caddy run -watch


build:
	#https://caddy.community/t/trying-to-build-caddy-from-source-to-test-2-5-0-beta-with-plugins/15393/2
	 xcaddy build --with github.com/airdb/caddywaf \
			 --with github.com/caddy-dns/dnspod \
			 --with github.com/lucas-clemente/quic-go@v0.25.0
