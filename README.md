# simple-webserver-go

Minimal Go HTTP server used to test BeforeProd preview deployments with the GO platform.

## Run locally

Requires Go 1.24+.

```bash
go run .
```

The server listens on port **4522** by default (BeforeProd's expected port). Override with `PORT`:

```bash
PORT=8080 go run .
```

Build a local binary:

```bash
go build -o app .
./app
```

## Preview deployments

On every pull request (`opened`, `synchronize`, `reopened`), the **Build Web Server** workflow:

1. Builds a Linux `amd64` binary named `app` into `./build`
2. Deploys it to [BeforeProd](https://beforeprod.com) via `beforeprod-com/preview-deployment-action`
3. Updates the PR description with the preview URL

When the PR is closed, **Cleanup PR Deployments** stops the preview app.

### Required secrets

Add these repository secrets (Settings → Secrets and variables → Actions):

| Secret | Description |
|--------|-------------|
| `BP_USER` | BeforeProd username |
| `BP_PASSWORD` | BeforeProd password |

### Workflow permissions

The deploy workflow needs `pull-requests: write` so the BeforeProd action can append the preview URL to the PR body. The cleanup workflow needs `pull-requests: read` to find that URL when the PR closes.

## Related

- JS/Next.js sibling: [simple-webserver-js](https://github.com/wunderkind2k1/simple-webserver-js)
- Preview action: [beforeprod-com/preview-deployment-action](https://github.com/beforeprod-com/preview-deployment-action)
