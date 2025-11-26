package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

const (
	appName    = "Blue-Green and Canary Demo App"
	appVersion = "v1.0.0"
	appColor   = "BLUE"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handleRoot)
	mux.HandleFunc("/healthz", handleHealth)
	mux.HandleFunc("/version", handleVersion)

	addr := ":8080"
	if port := os.Getenv("PORT"); port != "" {
		addr = ":" + port
	}

	log.Printf("Starting %s %s (%s) on %s\n", appName, appVersion, appColor, addr)
	if err := http.ListenAndServe(addr, loggingMiddleware(mux)); err != nil {
		log.Fatalf("server failed: %v", err)
	}
}

func handleRoot(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	html := fmt.Sprintf(`
<!doctype html>
<html>
  <head>
    <title>%s %s - %s</title>
    <style>
      body {
        font-family: system-ui, -apple-system, BlinkMacSystemFont, "Segoe UI", sans-serif;
        background: #0f172a;
        color: #e5e7eb;
        margin: 0;
        padding: 0;
        display: flex;
        flex-direction: column;
        height: 100vh;
      }
      .header {
        background: #1e293b;
        padding: 16px;
        text-align: center;
        font-size: 20px;
        font-weight: 600;
        letter-spacing: 0.16em;
        text-transform: uppercase;
        border-bottom: 1px solid #334155;
      }
      .footer {
        background: #1e293b;
        padding: 12px;
        text-align: center;
        font-size: 14px;
        color: #cbd5e1;
        border-top: 1px solid #334155;
      }
.content {
    flex: 1;
    display: flex;
    flex-direction: column;
    justify-content: center;   /* centers card vertically */
    align-items: center;
    gap: 0px;
    position: relative;
}

.title-box {
    padding: 10px 24px;
    border-radius: 999px;
    background: linear-gradient(90deg, #1d4ed8, #38bdf8);
    box-shadow: 0 14px 30px rgba(37, 99, 235, 0.45);
    border: 1px solid rgba(191, 219, 254, 0.4);
    margin-bottom: 12px;      /* small upward lift */
    position: relative;
    top: -40px;               /* moves title upward */
}

      .title {
        font-size: 18px;
        font-weight: 500;
        text-align: center;
        color: #e5e7eb;
        white-space: nowrap;
      }
      .card {
        padding: 40px 60px;
        border-radius: 22px;
        background: radial-gradient(circle at top, #1d4ed8, #020617);
        box-shadow: 0 25px 55px rgba(15, 23, 42, 0.75);
        text-align: center;
        max-width: 600px;
        margin-top: 12px;
      }
      .label {
        font-size: 14px;
        letter-spacing: 0.18em;
        text-transform: uppercase;
        color: #a5b4fc;
      }
      .version {
        font-size: 54px;
        font-weight: 900;
        margin: 16px 0 10px;
      }
      .pill {
        display: inline-flex;
        align-items: center;
        gap: 8px;
        padding: 8px 18px;
        border-radius: 999px;
        font-size: 14px;
        margin-top: 18px;
        background: rgba(15, 23, 42, 0.7);
        border: 1px solid rgba(129, 140, 248, 0.5);
      }
      .dot {
        width: 10px;
        height: 10px;
        border-radius: 999px;
        background: #22c55e;
      }
      .meta {
        margin-top: 22px;
        font-size: 15px;
        color: #cbd5f5;
      }
      a {
        color: #38bdf8;
        text-decoration: none;
      }
      a:hover {
        text-decoration: underline;
      }
    </style>
  </head>
  <body>

    <div class="header">
      KUBELANCER LABS
    </div>

<div class="content">
  <div class="title-box">
    <div class="title">
      Blue-Green &amp; Canary Demo Application
    </div>
  </div>

  <div class="card">
    <div class="label">BLUE / STABLE / CURRENT PROD VERSION</div>
    <div class="version">%s</div>

    <div class="pill">
      <span class="dot"></span>
      <span>Healthy • Listening on 8080</span>
    </div>

    <div class="meta">
      Try:
      <code><a href="/healthz">/healthz</a></code> •
      <code><a href="/version">/version</a></code>
    </div>
  </div>
</div>


<div class="footer">
  <strong><a href="https://labs.kubelancer.com" target="_blank" style="color:#38bdf8; text-decoration:none;">labs.kubelancer.com</a></strong>
  for Cloud, DevOps, GitOps, DevSecOps, Service Mesh & Monitoring Labs in Simple Demos</div>

  </body>
</html>
`, appName, appVersion, appColor, appVersion)

	_, _ = w.Write([]byte(html))
}

func handleHealth(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	_, _ = fmt.Fprintln(w, "ok")
}

func handleVersion(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	_, _ = fmt.Fprintf(w, "%s\n", appVersion)
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		log.Printf("%s %s from %s took %s\n",
			r.Method, r.URL.Path, r.RemoteAddr, time.Since(start))
	})
}
