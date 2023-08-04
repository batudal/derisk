FROM caddy:2.4.6-alpine-builder AS builder
RUN xcaddy build \
  --with github.com/abiosoft/caddy-exec
FROM caddy:2.4.6-alpine
COPY --from=builder Caddyfile /etc/caddy/Caddyfile
