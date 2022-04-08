FROM debian:stable-slim

ARG DEBIAN_FRONTEND=noninteractive

RUN apt update && \
    apt upgrade -y && \
    apt install -y --no-install-recommends --no-install-suggests ca-certificates dumb-init nginx-full libnginx-mod-http-lua && \
    rm -rf /var/lib/apt/lists/*

ENTRYPOINT ["/usr/bin/dumb-init", "--"]
CMD ["/usr/sbin/nginx", "-g", "daemon off;"]
