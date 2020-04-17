package main

import (
	"fmt"
	"time"

	// We'll need the net/http package
	"net/http"
)

// We need to define a handler for incoming requests
// Every function can be a handler if it takes in an IOWriter
// (http.ResponseWriter) and a request object
func genResponse(writer http.ResponseWriter, request *http.Request) {
	// We can grab the path that the request used by accessing request.URL.Path
	requestPath := request.URL.Path[1:]
	// If path is root...
	if requestPath == "" {
		requestPath = "World!"
	}
	// Grab current time in RFC1123 format
	currentTime := time.Now().Format(time.RFC1123)
	// An IOWriter is technically a file, so we'll use Fprintf
	// to write formatted text (the path the request came from
	// and current time) to it
	fmt.Fprintf(writer, "Hello %s!\n%s", requestPath, currentTime)
}

func main() {
	// http.HandleFunc allows us to define what
	// function will handle requests on what path
	// http.HandleFunc(path_name, handler_function_name)
	http.HandleFunc("/", genResponse)
	// http.ListenAndServe starts the server to start listening
	// This function blocks until program exits
	// http.ListenAndServe(address_or_port, handler *usually nil*)
	fmt.Println("Listening...")
	http.ListenAndServe(":8000", nil)
}
