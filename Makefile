dev:
	 xcaddy run -watch -config Caddyfile.test

run:
	 xcaddy run -watch

build:
	 xcaddy build --with github.com/airdb/caddywaf
