# An nginx lab in Docker

This repository contains a virtual "lab" that can be used to experiment with Nginx; it sets up two reverse proxies
(which can simulate having Cloudflare in front of your server) and a simple backend application.

## Requirements

- Docker
- docker-compose

## Usage

Run `docker-compose up` to spin up the whole stack, then your main (outer) frontend can be accessed on `localhost:8080`
and your secondary (inner) frontend on `localhost:9090`; the backend application is bound to the `/app` location on both
proxies.
