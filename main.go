package main

import (
	"bytes"
	"embed"
	"flag"
	"io/fs"
	"log"
	"net/http"
	"time"
)

//go:embed published
var staticFiles embed.FS

// Allow paths to render the default index.html
// for SPA-behavior that enables directly visiting URL paths
var allowedPathsForSPA []string

func setupServer() {
	// Set up API routes
	customizations()

	// Create a filesystem subtree'd within /published,
	// which contains the static assets embedded within the program at compile time
	publishedFS, err := fs.Sub(staticFiles, "published")
	if err != nil {
		log.Fatal("failed to create sub file system:", err)
	}

	// Grab the contents for index.html
	// This is a special case because we want to serve this file
	// as the catch-all for other routes, so we'll need its bytes for serving later
	indexHTMLContent, err := fs.ReadFile(publishedFS, "index.html")
	if err != nil {
		log.Fatal("failed to read index.html:", err)
	}
	lastModified := time.Now() // Use boot time as the index.html's last modified time since it's embedded within the binary and won't change

	http.Handle("/", http.FileServerFS(publishedFS))

	// Allow index.html to be served to get SPA-behavior
	// where certain allowed URL paths can be directly visited and not 404
	for _, path := range allowedPathsForSPA {
		p := path
		http.HandleFunc(p, func(w http.ResponseWriter, r *http.Request) {
			if r.URL.Path == p {
				reader := bytes.NewReader(indexHTMLContent)
				http.ServeContent(w, r, "index.html", lastModified, reader)
			} else {
				http.NotFound(w, r)
			}
		})
	}
}

func main() {
	setupServer()

	var serverAddress string
	flag.StringVar(&serverAddress, "addr", "127.0.0.1:8080", "HTTP server address")
	flag.Parse()

	log.Printf("Server started on http://%s", serverAddress)
	if err := http.ListenAndServe(serverAddress, nil); err != nil {
		log.Fatal(err)
	}
}
