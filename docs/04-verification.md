# Verification evidence

These are the actual, sanitized results from the development machine after the final service and container implementation. They are evidence, not commands that a beginner must reproduce byte-for-byte.

## Go checks

```text
$ go test ./...
ok  github.com/tristanscholten/hackathon-demo/cmd/server  0.003s

$ go test -race ./...
ok  github.com/tristanscholten/hackathon-demo/cmd/server  1.012s

$ go vet ./...
(exit 0)

$ go build ./cmd/server
(exit 0)

$ go version
go version go1.26.0 linux/amd64
```

## Podman smoke test

The Containerfile built successfully with `golang:1.27.1-alpine` and `alpine:3.23`. Image inspection reported the final image user as `app`.

```text
GET /hello:    Hello from hackathon-demo!
GET /healthz:  ok
GET /not-here: 404
GET /metrics:  hackathon_demo_http_requests_total 5
container state: running
```

The counter reads `5` because the health readiness check, hello request, health request, missing-route request, and metrics request were all handled by the process. The container was removed after the smoke test.

## What was deliberately not done

No image was published, no release tag was created, and the repository remains private pending an independent review and an authorized public transition. See [the GHCR lesson](03-ghcr.md).
