# Lesson 2: put the service in a container

Use the repository root for every command below.

## 1. Build an image

```sh
podman build -t localhost/hackathon-demo:dev -f Containerfile .
```

`-t` gives the image a memorable local name. The final image is built from a pinned Alpine release, runs as a non-root user, and listens on port 8080. The first build downloads base images, so it can take longer.

## 2. Run a container

```sh
podman run --rm --name hackathon-demo -p 8080:8080 localhost/hackathon-demo:dev
```

Read `-p 8080:8080` left to right:

- The first `8080` is your computer's port.
- The second `8080` is the container's port.

In a second terminal, run the same probes as in [Lesson 1](01-local-go.md):

```sh
curl http://localhost:8080/hello
curl http://localhost:8080/healthz
curl http://localhost:8080/metrics
```

## 3. Stop it

If you started it in the foreground, press `Ctrl+C`. `--rm` removes the stopped container automatically. If it was started in the background, use:

```sh
podman stop hackathon-demo
```

Next: [Lesson 3: publish to GHCR](03-ghcr.md).
