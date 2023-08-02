#FROM airdb/builder:xcaddy as builder
FROM golang:1.20-alpine as builder

WORKDIR /build

ADD ./ /build

RUN cd /build/ && \
	CGO_ENABLED=0 GOOS=linux go build -ldflags '-w -s' -o caddy ./cmd/main.go


FROM airdb/base:latest

WORKDIR /app/

COPY --from=builder /build/caddy /app/
COPY --from=builder /build/conf /app/conf

EXPOSE 8080

RUN chmod +x /app/caddy

#ENTRYPOINT ["sleep", "3600"]
#ENTRYPOINT ["/app/caddy", "run"]
CMD ["/app/caddy", "run", "--config", "conf/Caddyfile"]
