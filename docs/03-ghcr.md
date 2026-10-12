# Lesson 3: publish to GitHub Container Registry (GHCR)

An image is a reusable package of this program. GHCR is a registry that stores images so people can download them. It is not hosting: pushing an image does not keep a web service running.

A public GitHub repository and a public GHCR package are separate things. A newly published personal package is private by default.

Choose **one** publishing path:

- [Manual push](#path-a-manual-push-from-your-computer) teaches the registry commands.
- [GitHub Actions](#path-b-automatic-publish-with-github-actions) is the normal choice after a reviewed change.

Both paths need the one-time public-visibility step below if anonymous downloads are wanted.

## Before either path: make the package public once

Do this only after the first image exists. It is a manual GitHub web step; neither `podman push` nor this repository's workflow can make the package public for you.

1. On GitHub, open your personal profile → **Packages** → **hackathon-demo** → **Package settings**.
2. Under **Danger Zone**, choose **Change visibility** → **Public**.
3. Type the package name when GitHub asks, then click **I understand the consequences, change package visibility**.

Warning: GitHub says a package made public cannot be made private again. Package permissions/inheritance control who can manage the package; they do not make it public.

The expected public package page is <https://github.com/users/tristanscholten/packages/container/package/hackathon-demo>. If it is not public yet, anonymous pulls will fail.

## Path A: manual push from your computer

Use this path to learn or when you intentionally do not use CI. Do not paste a token into a command, a file, Git, or shell history.

1. Prerequisites: a GitHub account, Podman, and this repository checked out locally. Build the image first:

   ```sh
   podman build -t localhost/hackathon-demo:dev -f Containerfile .
   ```

2. Choose your own new version. Image owner and package names must be lowercase:

   ```sh
   OWNER=tristanscholten
   VERSION=v0.1.1
   IMAGE="ghcr.io/${OWNER}/hackathon-demo:${VERSION}"
   ```

   `v0.1.0` already exists for this project. Do not try to overwrite a released, immutable tag; choose a new version such as the example above.

3. Create a **classic** personal access token in GitHub's [token settings](https://github.com/settings/tokens/new). GitHub Packages supports classic PATs for this use. Select `write:packages` to push; select `read:packages` only when you need to pull a private image. GitHub's UI may automatically select broad `repo` when you select `write:packages`; clear it unless you separately need that broad repository access.

4. At the masked prompt below, paste the token. It is read from standard input, not placed in the command line:

   ```sh
   read -r -s -p 'GitHub classic PAT: ' CR_PAT; printf '\n'
   printf '%s' "$CR_PAT" | podman login ghcr.io -u "$OWNER" --password-stdin
   unset CR_PAT
   ```

5. Tag and push the local image:

   ```sh
   podman tag localhost/hackathon-demo:dev "$IMAGE"
   podman push "$IMAGE"
   ```

   Typical output names uploaded layers and then prints the image reference or digest. It is illustrative only: layer IDs and digests differ on every build.

6. Remove the saved registry login when finished, especially on a shared computer:

   ```sh
   podman logout ghcr.io
   ```

   Protect Podman's auth file as you would a password. Do not commit it or copy it into a tutorial.

## Path B: automatic publish with GitHub Actions

This repository already has `.github/workflows/publish-ghcr.yml`. Leave its workflow code alone for this lesson.

1. The workflow runs only on a push to trusted `main` or a pushed tag matching `v*`. It does not publish pull requests.
2. It has `contents: read` and `packages: write`, then logs in with the short-lived `GITHUB_TOKEN`. You do not create a PAT for this CI path. If your repository restricts the default Actions token, confirm **Settings** → **Actions** → **General** allows the workflow to receive the permissions declared in the file.
3. A push to `main` publishes the tag `main`. A tag push publishes exactly that tag. To publish the next release after review, use a new annotated SemVer tag; do not recreate `v0.1.0`:

   ```sh
   git switch main
   git pull --ff-only origin main
   git tag -a v0.1.1 -m 'v0.1.1'
   git push origin v0.1.1
   ```

4. Open the workflow run in the **Actions** tab and confirm the `Publish GHCR image` job succeeded. Then complete the one-time public-visibility step above if you want anonymous pulls.

## Verify a public image without saved credentials

Run this only after the package is public. It uses an empty temporary auth file and `--policy always` so Podman does not reuse a cached image or your normal credentials:

```sh
AUTHFILE="$(mktemp)"
trap 'rm -f "$AUTHFILE"' EXIT
podman pull --authfile "$AUTHFILE" --policy always ghcr.io/tristanscholten/hackathon-demo:v0.1.0
podman run --rm --name hackathon-demo-public -d -p 18080:8080 ghcr.io/tristanscholten/hackathon-demo:v0.1.0
curl http://localhost:18080/hello
curl http://localhost:18080/healthz
curl http://localhost:18080/metrics
podman stop hackathon-demo-public
```

The three routes are `/hello`, `/healthz`, and `/metrics`. Port `18080` avoids a local service already using `8080`. `--rm` removes the stopped container; the trap removes the temporary empty auth file.

## Cost and policy note

GitHub currently documents free storage and data transfer for public packages, including the Container registry, and says it will announce future Container registry billing changes at least one month ahead. That is a current policy, not a forever-free hosting promise. Read GitHub's [Packages billing documentation](https://docs.github.com/en/billing/concepts/product-billing/github-packages) before relying on it.

## Sources

- [GitHub: Container registry](https://docs.github.com/en/packages/working-with-a-github-packages-registry/working-with-the-container-registry)
- [GitHub: package visibility](https://docs.github.com/en/packages/learn-github-packages/configuring-a-packages-access-control-and-visibility)
- [GitHub: publishing with Actions](https://docs.github.com/en/packages/managing-github-packages-using-github-actions-workflows/publishing-and-installing-a-package-with-github-actions)

Next: [troubleshooting](05-troubleshooting.md).