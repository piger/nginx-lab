# An nginx lab in Docker

This repository contains a virtual "lab" that can be used to experiment with Nginx; it sets up two reverse proxies
(which can simulate having Cloudflare in front of your server) and a simple backend application.

## Requirements

- Docker
- docker-compose

## Usage

Run `docker compose up` to spin up the whole stack, then your main (outer) frontend can be accessed
on `localhost:8080` (HTTP) or `localhost:8443` (HTTPS) and your secondary (inner) frontend on
`localhost:9090` (and 9443 for HTTPS); the backend application is bound to the `/app` location on
both proxies.

### TLS support

The two proxies are configured with self-signed TLS certificates created at build time for the wildcard
domain `*.homelab.dev`.

To send an HTTPS request to the outer proxy with `curl`:

```
$ curl --insecure --resolve '*:8443:127.0.0.1' https://example.homelab.dev:8443
```

### Logs

In an ideal world Docker Compose would distinguish between stdout and sterr when streaming logs, but
until then:

```
$ docker compose logs -f outer-proxy -n 0 --no-log-prefix | jq -R '. as $line | try (fromjson) catch $line'
```
