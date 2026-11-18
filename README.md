# preview-demo-go

Minimal **Go** HTTP server that shows [Beforeprod](https://beforeprod.com) preview deployments on every pull request.

Open a PR, GitHub Actions builds the app, Beforeprod spins up an isolated preview, and the **HTTPS preview URL is written into the PR description**. Close the PR and the preview is cleaned up.

Fork this repo and open a pull request to try it.

Sibling demo (JavaScript / Next.js): [`preview-demo-js`](https://github.com/beforeprod-com/preview-demo-js)

Action used under the hood: [`preview-deployment-action`](https://github.com/beforeprod-com/preview-deployment-action)

## Run locally

Requires Go 1.24+.

```bash
go run .
```

Listens on port **4522** by default. Override with `PORT`:

```bash
PORT=8080 go run .
```

```bash
go build -o app .
./app
```

## What the workflows do

On every pull request (`opened`, `synchronize`, `reopened`), **Build Web Server**:

1. Builds a Linux `amd64` binary named `app` into `./build`
2. Deploys it to Beforeprod via `beforeprod-com/preview-deployment-action` (`platform: GO`)
3. Writes the preview link into the PR description

When the PR is closed, **Cleanup PR Deployments** stops the preview app.

### Required secrets

| Secret | Description |
|--------|-------------|
| `BP_USER` | Beforeprod username |
| `BP_PASSWORD` | Beforeprod password |

### Permissions

Deploy workflow needs `pull-requests: write`. Cleanup needs `pull-requests: read`.

## Want this on your repo?

Beforeprod is in early access. [Request access on beforeprod.com](https://beforeprod.com/?ref=demo-go)

## Learn more

- Product: [beforeprod.com](https://beforeprod.com)
- Blog: [What is Beforeprod](https://beforeprod.com/blog/what-is-beforeprod)