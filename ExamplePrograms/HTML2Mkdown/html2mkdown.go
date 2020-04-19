package main

import (
	"net/http"
	// Markdown processor for Go + HTML sanitizer
	"github.com/microcosm-cc/bluemonday"
	"gopkg.in/russross/blackfriday.v2"
)

// GenerateMkdown Take in an HTTP request, grab the value of the
// body into slice of bytes. Send to BF for processing, to BM
// for sanitization and then write the slice of bytes to the
// ResponseWriter
func GenerateMkdown(resw http.ResponseWriter, r *http.Request) {
	mkdown := blackfriday.Run([]byte(r.FormValue("body")))
	sanitized := bluemonday.UGCPolicy().SanitizeBytes(mkdown)
	resw.Write(sanitized)
}

func main() {
	// Bind handler for /mkdown route
	http.HandleFunc("/mkdown", GenerateMkdown)
	// Server our HTML
	http.Handle("/", http.FileServer(http.Dir("static")))
	// Start default server
	http.ListenAndServe(":8000", nil)
}
