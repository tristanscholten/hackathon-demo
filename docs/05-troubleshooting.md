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

## `denied` when pulling from GHCR

The package may still be private. A public repository does not automatically make its GHCR package public. Follow the one-time **Package settings** → **Danger Zone** → **Change visibility** → **Public** instructions in [Lesson 3](03-ghcr.md). Public GHCR packages allow anonymous pulls.

To test without your normal saved login or cached image, use the empty-auth-file command from Lesson 3. Do not put a PAT in `podman pull` command arguments.

## `denied` when pushing to GHCR

Check the image name is fully qualified and lowercase:

```text
ghcr.io/tristanscholten/hackathon-demo:TAG
```

For a manual push, authenticate with a **classic** PAT created with GitHub's [package-only `write:packages` link](https://github.com/settings/tokens/new?scopes=write:packages), then send it through `--password-stdin`. This documented link avoids the unnecessary broad `repo` scope; see [Lesson 3](03-ghcr.md).

For Actions, check that `.github/workflows/publish-ghcr.yml` still declares `packages: write`, runs from this repository, and was triggered by a push to `main` or a `v*` tag. The workflow does not publish pull requests.

## The image version already exists

Do not overwrite a released image tag. Choose a new SemVer version, such as `v0.1.1`, and for Actions create and push a new annotated Git tag. See [Lesson 3](03-ghcr.md).

## `go: command not found`

Install Go, then open a new terminal and retry:

```sh
go version
```

Do not set `GOPATH` just to run this module. Modern Go modules work outside GOPATH. If you install Go tools and their binary directory is missing from your `PATH`, follow the [official Go installation instructions](https://go.dev/doc/install).

## `/metrics` looks unfamiliar

Prometheus reads plain text. `# HELP` describes a metric, `# TYPE` says how Prometheus should interpret it, and the final line is the current number. The number increases for every request the application handles, including the request that reads `/metrics`.