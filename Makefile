dev: build
	./caddy run -watch -config Caddyfile.dev

run:
	./caddy run -watch -config Caddyfile.prod

fmt:
	./caddy fmt -overwrite

build:
	# https://caddy.community/t/trying-to-build-caddy-from-source-to-test-2-5-0-beta-with-plugins/15393/2
	# xcaddy build --with github.com/airdb/caddywaf \
	# 	--with github.com/caddy-dns/dnspod
	go build -o caddy ./cmd