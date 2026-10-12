# hackathon-demo

A tiny, beginner-friendly Go service. It teaches four things in one small project:

1. Serve `GET /hello` and `GET /healthz`.
2. Expose a Prometheus text endpoint at `GET /metrics`.
3. Put the program in a small container image.
4. Publish that image to GitHub Container Registry (GHCR).

Start here: [your first local run](docs/01-local-go.md). Then use [the container lesson](docs/02-container.md), the [GHCR lesson](docs/03-ghcr.md) (choose manual push **or** GitHub Actions publishing), and [troubleshooting](docs/05-troubleshooting.md). The requirements are in [requirements/README.md](requirements/README.md).

## What you will build

`/hello` says hello. `/healthz` says `ok`. `/metrics` exposes one Prometheus counter:

```text
hackathon_demo_http_requests_total
```

It counts every HTTP request handled by this server, including requests to `/metrics` and requests that receive a `404` response. The counter uses Go's `sync/atomic`, so concurrent requests do not lose increments.

## Fast start

Run these commands from the repository root:

```sh
go test ./...
go run ./cmd/server
```

In a second terminal:

```sh
curl http://localhost:8080/hello
curl http://localhost:8080/healthz
curl http://localhost:8080/metrics
```

Expected response snippets:

```text
Hello from hackathon-demo!
ok
# TYPE hackathon_demo_http_requests_total counter
hackathon_demo_http_requests_total 3
```

The precise counter value may be higher if you refresh a route.

## Prerequisites

- A current Go toolchain. This repository's module baseline is Go 1.26.
- [Podman](https://podman.io/) for the container lesson. Docker commands are similar, but the tested commands in this project use Podman.
- A GitHub account only for the GHCR lesson.

On Linux, install Go and Podman with your distribution's package manager or their official installers. On macOS or Windows, install Go normally and use Podman Desktop or another supported Podman setup. The service itself has no non-standard Go dependencies.

## Concepts in plain English

- `localhost` means “this computer.” `localhost:8080` is the service on your computer, at port 8080.
- A **port** is a numbered doorway used by network programs. This demo listens on port 8080 by default.
- An **image** is the reusable package containing an application and its runtime files.
- A **container** is one running copy of an image.
- GHCR is an image **registry**: it stores images for download. It does not run or host the service for you.

## Configuration

The program listens on all interfaces (`0.0.0.0`) at port 8080 by default. Set `PORT` to use a different port:

```sh
PORT=9090 go run ./cmd/server
```

For a complete beginner path, read the linked lessons in order. For actual verification recorded while preparing this change, see [docs/04-verification.md](docs/04-verification.md).

## License

[MIT](LICENSE)
