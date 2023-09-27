FROM debian:stable-slim

ARG DEBIAN_FRONTEND=noninteractive

RUN apt-get update && \
    apt-get upgrade -y && \
    apt-get install -y --no-install-recommends --no-install-suggests \
        ca-certificates dumb-init nginx-full libnginx-mod-http-lua libnss3-tools curl && \
    rm -rf /var/lib/apt/lists/*

RUN curl -JLO "https://dl.filippo.io/mkcert/latest?for=linux/amd64" && \
    chmod +x mkcert-v*-linux-amd64 && \
    mv mkcert-v*-linux-amd64 /usr/local/bin/mkcert && \
    mkdir /etc/ssl/nginx && \
    /usr/local/bin/mkcert -cert-file /etc/ssl/nginx/_.homelab.dev.crt -key-file /etc/ssl/nginx/_.homelab.dev.key "*.homelab.dev"

ENTRYPOINT ["/usr/bin/dumb-init", "--"]
CMD ["/usr/sbin/nginx", "-g", "daemon off;"]
