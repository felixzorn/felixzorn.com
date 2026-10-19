// felixzorn.com: liefert die Inselreise (web/index.html) aus einem einzigen Go-Binary aus.
//
//	go run .            # http://localhost:8080
package main

import (
	"embed"
	"fmt"
	"io/fs"
	"log"
	"net/http"
	"os"
	"time"
)

//go:embed web
var webFS embed.FS

var started = time.Now()

func main() {
	addr := env("ADDR", ":8080")

	static, err := fs.Sub(webFS, "web")
	if err != nil {
		log.Fatal(err)
	}

	mux := http.NewServeMux()
	mux.HandleFunc("GET /healthz", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "ok · uptime %s\n", time.Since(started).Round(time.Second))
	})
	mux.Handle("GET /", http.FileServerFS(static))

	srv := &http.Server{
		Addr:              addr,
		Handler:           mux,
		ReadHeaderTimeout: 5 * time.Second,
	}
	log.Printf("felixzorn.com läuft auf %s", addr)
	log.Fatal(srv.ListenAndServe())
}

func env(k, def string) string {
	if v := os.Getenv(k); v != "" {
		return v
	}
	return def
}
