package main

import (
	"log"
	"net/http"
)

// func hello(w http.ResponseWriter, req *http.Request) {
// 	fmt.Fprintf(w, "hello\n")
// }

// func headers(w http.ResponseWriter, req *http.Request) {
// 	for name, headers := range req.Header {
// 		for _, h := range headers {
// 			fmt.Fprintf(w, "%v: %v\n", name, h)
// 		}
// 	}
// }

func main() {

	mux := http.NewServeMux()
	mux.HandleFunc("/", okHandler)

	// Wrap the servemux with the limit middleware.
	log.Print("Listening on :4000...")
	http.ListenAndServe(":4000", limit(limitByUsers(mux)))
}

func okHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Пока все в порядке"))
}
