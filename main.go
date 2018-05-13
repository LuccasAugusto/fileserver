package main

import "net/http"

func main() {

	// Building a server.
	http.ListenAndServe(":8080", http.FileServer(http.Dir(".")))
}
