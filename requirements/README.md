# Minimal requirements

Source and owner: Tristan Scholten. Purpose: a zero-experience beginner tutorial for a tiny Go HTTP service, container workflow, Prometheus-format metric, and GHCR publishing path.

## In scope

- `GET /hello`, `GET /healthz`, and `GET /metrics`.
- A concurrency-safe, Prometheus text counter for every HTTP request handled by the process, including `/metrics` and 404s.
- A non-root, pinned-base container image and verified Podman workflow.
- Linked beginner documentation for local Go, containers, GHCR, and troubleshooting.
- GitHub Actions checks and trusted-branch/tag GHCR publishing guidance.

## Out of scope

Kubernetes, Helm, databases, authentication, OAuth, GitOps, and a monitoring stack.

## Acceptance checks

| ID | Obligation | Verification |
| --- | --- | --- |
| R1 | The service returns a successful hello and health response. | `httptest` and container HTTP probes. |
| R2 | Metrics use stable Prometheus HELP, TYPE, and counter lines. | `httptest` asserts the exposition. |
| R3 | Concurrent requests do not race or lose counter increments. | `go test -race ./...` and concurrent unit test. |
| R4 | The documented container starts as non-root and serves port 8080. | Inspect Containerfile; build/run/probe with Podman. |
| R5 | Docs explain GHCR publishing without claiming deployment or automatic public visibility. | Documentation review and link check. |

Priority: all items are required for the initial tutorial because omitting one breaks the stated learning path. Acceptance authority: the repository reviewer.
