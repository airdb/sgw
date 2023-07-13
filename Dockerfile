FROM airdb/builder:xcaddy as builder

WORKDIR /build

ADD ./ /build

RUN cd /build/ && \
	CGO_ENABLED=0 GOOS=linux go build -ldflags '-w -s' -o caddy ./cmd/main.go


FROM airdb/base:latest

WORKDIR /app/

COPY --from=builder /build/caddy /app/
COPY --from=builder /build/Caddyfile /app/

EXPOSE 8080

RUN chmod +x /app/caddy

#ENTRYPOINT ["sleep", "3600"]
ENTRYPOINT ["/app/caddy", "run"]
