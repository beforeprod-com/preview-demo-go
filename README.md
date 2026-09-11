# preview-demo-go

Minimal **Go** HTTP server that shows [Beforeprod](https://beforeprod.com) preview deployments on every pull request.

Open a PR → GitHub Actions builds the app → Beforeprod spins up an isolated preview → the **HTTPS preview URL is written into the PR description**. Close the PR → the preview is cleaned up.

Sibling demo (JavaScript / Next.js): [`preview-demo-js`](https://github.com/beforeprod-com/preview-demo-js)  
Action used under the hood: [`preview-deployment-action`](https://github.com/beforeprod-com/preview-deployment-action)

## Run locally

Requires Go 1.24+.

```bash
go run .
