# Troubleshooting

## `bind: address already in use`

Something already uses the local port. Stop the other program, or use a different host port:

```sh
podman run --rm -p 9090:8080 localhost/hackathon-demo:dev
curl http://localhost:9090/hello
```

For a direct Go run, choose a different application port:

```sh
PORT=9090 go run ./cmd/server
```

## Podman cannot connect or cannot start a container

Check that Podman works before blaming this image:

```sh
podman info
podman ps
```

On rootless Linux setups, Podman normally runs as your normal user. Do not retry with `sudo` just to make this demo work; inspect the reported rootless/networking error and follow your operating system's Podman instructions.

## `denied` when pushing to GHCR

Check the image name is fully qualified:

```text
ghcr.io/tristanscholten/hackathon-demo:TAG
```

For Actions, check the workflow has `packages: write` and is running from the repository that owns the package. For manual pushes, authenticate with a classic PAT that has `write:packages`, sent through `--password-stdin`. A package previously created without a repository link can need manual access configuration; see the [GHCR guide](https://docs.github.com/en/packages/working-with-a-github-packages-registry/working-with-the-container-registry).

## `go: command not found`

Install Go, then open a new terminal and retry:

```sh
go version
```

Do not set `GOPATH` just to run this module. Modern Go modules work outside GOPATH. If you install Go tools and their binary directory is missing from your `PATH`, follow the [official Go installation instructions](https://go.dev/doc/install).

## `/metrics` looks unfamiliar

Prometheus reads plain text. `# HELP` describes a metric, `# TYPE` says how Prometheus should interpret it, and the final line is the current number. The number increases for every request the application handles, including the request that reads `/metrics`.
