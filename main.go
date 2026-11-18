package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

const defaultPort = "4522"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	http.HandleFunc("/", handleHome)
	http.HandleFunc("/health", handleHealth)

	addr := fmt.Sprintf(":%s", port)
	log.Printf("Starting server on %s", addr)
	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatal(err)
	}
}

func handleHealth(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "ok")
}

func handleHome(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, homePage)
}

const homePage = `<!DOCTYPE html>
<html lang="en">
<head>
  <meta charset="utf-8">
  <meta name="viewport" content="width=device-width, initial-scale=1">
  <title>This is a Beforeprod preview</title>
  <style>
    * { box-sizing: border-box; margin: 0; padding: 0; }
    body {
      font-family: ui-sans-serif, system-ui, -apple-system, "Segoe UI", Roboto, sans-serif;
      min-height: 100vh;
      display: flex;
      align-items: center;
      justify-content: center;
      padding: 6rem 1.5rem;
      background: #fafafa;
      color: #171717;
    }
    main { max-width: 36rem; width: 100%; text-align: center; display: flex; flex-direction: column; align-items: center; gap: 1.5rem; }
    h1 { font-size: 2.25rem; font-weight: 700; letter-spacing: -0.025em; }
    .badge { display: inline-block; border: 1px solid #cbd5e1; background: #e2e8f0; color: #0f172a; border-radius: 9999px; padding: 0.25rem 0.75rem; font-size: 0.875rem; font-weight: 600; }
    h2 { font-size: 1.5rem; font-weight: 600; }
    .lead { font-size: 1.125rem; color: #4b5563; }
    a { color: inherit; text-decoration: underline; text-underline-offset: 4px; }
    a:hover { color: #0f172a; }
  </style>
</head>
<body>
  <main>
    <h1>This is a Beforeprod preview</h1>
    <span class="badge">Go</span>
    <h2>A Go app, deployed from GitHub Actions for this PR.</h2>
    <p class="lead">Real HTTPS. Torn down when the PR closes.</p>
    <p><a href="https://beforeprod.com">beforeprod.com</a> &middot; Early access</p>
  </main>
</body>
</html>`