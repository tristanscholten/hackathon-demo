# Lesson 3: publish to GitHub Container Registry (GHCR)

GHCR stores OCI/Docker images. It is a registry, not a free service that keeps your HTTP program running.

Current GitHub documentation says public GitHub Packages use is free, and Container registry storage and bandwidth are currently free; GitHub says it will give at least one month's notice before changing the Container registry policy. Check the [billing documentation](https://docs.github.com/en/billing/concepts/product-billing/github-packages) before relying on it.

## Recommended: let GitHub Actions publish

This repository includes `.github/workflows/publish-ghcr.yml`. On a trusted push to `main` or a `v*` tag, it logs in using the short-lived `GITHUB_TOKEN` with only `contents: read` and `packages: write`, builds, and pushes `ghcr.io/tristanscholten/hackathon-demo`.

It does not publish on pull requests. That prevents unreviewed pull-request code from getting publishing credentials.

A maintainer can create a release tag after review, for example:

```sh
git tag -a v0.1.0 -m "v0.1.0"
git push origin v0.1.0
```

This project deliberately does not create that tag or publish an image during the initial review. Release and public-package steps need their own authorization.

The workflow publishes `main` pushes as `main` and `v0.1.0` as `v0.1.0`. Pull an image after it has been published and made public:

```sh
podman pull ghcr.io/tristanscholten/hackathon-demo:v0.1.0
podman run --rm -p 8080:8080 ghcr.io/tristanscholten/hackathon-demo:v0.1.0
```

## Package visibility is separate from repository visibility

A public repository does **not** automatically mean a Container registry package is public. GitHub says a newly published package is private by default. After the first publish, open the package page, select **Package settings**, then **Change visibility** and choose **Public**. Public Container registry packages can be pulled anonymously. See GitHub's [package visibility instructions](https://docs.github.com/en/packages/learn-github-packages/configuring-a-packages-access-control-and-visibility).

The Containerfile adds `org.opencontainers.image.source=https://github.com/tristanscholten/hackathon-demo`, which connects the image with its source repository before publication.

## Optional: manual push with a personal access token

Use this only when you cannot use the workflow. Create a classic PAT with the smallest required package scope: `write:packages` to push; use `read:packages` only if you only pull private images. Do not add broad `repo` scope just for this action. Never paste the token into a terminal history, source file, or Git commit.

Store it in an environment variable supplied safely by your shell, then pass it through standard input:

```sh
printf '%s' "$CR_PAT" | podman login ghcr.io -u YOUR_GITHUB_USERNAME --password-stdin
podman tag localhost/hackathon-demo:dev ghcr.io/tristanscholten/hackathon-demo:dev
podman push ghcr.io/tristanscholten/hackathon-demo:dev
unset CR_PAT
```

GitHub's [Container registry guide](https://docs.github.com/en/packages/working-with-a-github-packages-registry/working-with-the-container-registry) explains PAT scopes, `GITHUB_TOKEN`, package linkage, and visibility.

Next: [troubleshooting](05-troubleshooting.md).
