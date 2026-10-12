# Lesson 1: run Go on your computer

Work from the repository root (the directory that contains `go.mod`).

## 1. Check your Go installation

```sh
go version
go test ./...
go test -race ./...
```

`go test` runs the handler tests. `-race` also checks that concurrent access to the metric counter is safe.

## 2. Start the server

```sh
go run ./cmd/server
```

Keep this terminal open. It prints the local URL where the server is listening.

## 3. Ask the server for data

Open a second terminal in any directory:

```sh
curl http://localhost:8080/hello
curl http://localhost:8080/healthz
curl http://localhost:8080/metrics
curl -i http://localhost:8080/not-here
```

You should see `Hello from hackathon-demo!`, `ok`, a Prometheus-format metric, and a `404 Not Found` response. Visit `/hello` a few times, then request `/metrics` again: the number increases.

## 4. Stop it

Return to the terminal running `go run` and press `Ctrl+C`.

Next: [Lesson 2: build and run a container](02-container.md).
