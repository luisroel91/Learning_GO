package main

import (
	"fmt"
	"net/http"
)

// He're we'll make a simple program that uses goroutines &
// net/http to make a non-blocking website checker.
// We'll learn how goroutines work/communicate via channels

// Let's write the function that will check our website list
// by making a GET request to the servers
// * Note that it takes in a link (string) and c (channel string)
// * To send data to channel ch <- v (send v to channel ch)
// * To receive from channel and assign to v v := <-ch
// * To receiveData flows in the direction of the arrow
// * You must create them before use using make(chan *type*)
func checkServer(uri string, ch chan string) {
	// We don't really care about the http response, we just
	// care whether the server errors out or not
	_, err := http.Get(uri)
	// If there's an error then we know server is down
	if err != nil {
		// Feed info into channel
		ch <- uri + " is down"
		return
	}
	// If we don't get an error then we can assume server is up
	// Feed data into channel
	ch <- uri + " is up"
	return
}

func printResult(ch chan string) {
	for result := range ch {
		fmt.Println(result)
	}
}

func main() {
	servers := []string{
		"http://google.com",
		"http://facebook.com",
	}
	// Before using a channel with goroutines, we need to
	// create it using make(), channels have two "types"
	// chan and then string, int, etc. There are channel types
	// for every data type
	c := make(chan string)

	// Loop through server list
	for _, server := range servers {
		// Execute our function as a goroutine
		// Notice how we pass our channel
		go checkServer(server, c)
	}
	printResult(c)
	close(c)
}
