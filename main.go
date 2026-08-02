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
  <title>Simple Web Server Go</title>
  <style>
    * { box-sizing: border-box; margin: 0; padding: 0; }
    body {
      font-family: ui-monospace, SFMono-Regular, Menlo, Monaco, Consolas, monospace;
      min-height: 100vh;
      display: flex;
      flex-direction: column;
      align-items: center;
      justify-content: space-between;
      padding: 6rem 1.5rem 4rem;
      background: #fafafa;
      color: #171717;
    }
    h1 { font-size: 2.25rem; font-weight: 700; text-align: center; margin-bottom: 2rem; }
    .hero { text-align: center; flex: 1; display: flex; align-items: center; }
    .hero h2 { font-size: 1.5rem; font-weight: 600; margin-bottom: 1rem; }
    .hero p { font-size: 1.125rem; color: #525252; }
    .cards {
      display: grid;
      gap: 1rem;
      width: 100%;
      max-width: 64rem;
      grid-template-columns: repeat(auto-fit, minmax(14rem, 1fr));
    }
    .card {
      border: 1px solid transparent;
      border-radius: 0.5rem;
      padding: 1.25rem;
      transition: border-color 0.15s, background 0.15s;
    }
    .card:hover { border-color: #d4d4d4; background: #f5f5f5; }
    .card h3 { font-size: 1.25rem; font-weight: 600; margin-bottom: 0.75rem; }
    .card p { font-size: 0.875rem; opacity: 0.6; max-width: 30ch; }
    .arrow { display: inline-block; transition: transform 0.15s; }
    .card:hover .arrow { transform: translateX(0.25rem); }
  </style>
</head>
<body>
  <header>
    <h1>Welcome to Simple Web Server Go</h1>
  </header>

  <section class="hero">
    <div>
      <h2>Your Go app is running!</h2>
      <p>This is a very simple Go HTTP web server.</p>
    </div>
  </section>

  <section class="cards">
    <div class="card">
      <h3>Go <span class="arrow">&rarr;</span></h3>
      <p>Built with the Go standard library</p>
    </div>
    <div class="card">
      <h3>BeforeProd <span class="arrow">&rarr;</span></h3>
      <p>Preview deployments on pull requests</p>
    </div>
    <div class="card">
      <h3>Minimal <span class="arrow">&rarr;</span></h3>
      <p>Single binary, no external dependencies</p>
    </div>
    <div class="card">
      <h3>Simple <span class="arrow">&rarr;</span></h3>
      <p>Clean and minimal setup</p>
    </div>
  </section>
</body>
</html>`
