dev: build
	./caddy run -watch -config Caddyfile

run:
	./caddy run -watch -config Caddyfile

fmt:
	./caddy fmt -overwrite ./Caddyfile
	#./caddy fmt -overwrite ./Caddyfile.dev
	#./caddy fmt -overwrite ./Caddyfile.live

build:
	go build -o caddy ./cmd/main.go

xcaddy:
	# https://caddy.community/t/trying-to-build-caddy-from-source-to-test-2-5-0-beta-with-plugins/15393/2
	# xcaddy build --with github.com/airdb/caddywaf \
	# 	--with github.com/caddy-dns/dnspod
