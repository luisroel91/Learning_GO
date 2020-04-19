package main

// We'll need the net/http package
import "net/http"

func main() {
	// This will serve the current directory over HTTP
	http.ListenAndServe(":8000", http.FileServer(http.Dir(".")))
}
