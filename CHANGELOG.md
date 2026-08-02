# Changelog

Notable changes to this demo app. Use it as a reference for wiring BeforeProd preview deployments into a Go repo.

## [Unreleased]

### Added
- Minimal Go HTTP server used as a BeforeProd preview-deployment example
- PR workflows for deploy and cleanup via `beforeprod-com/preview-deployment-action`
- Linux `amd64` build producing binary `app` in `./build` for the GO platform

### Notes
- Deploy workflow requests `pull-requests: write` so the BeforeProd action can write the preview URL into the PR body
- Cleanup workflow requests `pull-requests: read` to read that URL when the PR closes
- Workflows use `checkout@v7` and `setup-go@v7`
- Server listens on `PORT` (default `4522`, BeforeProd's expected port)
