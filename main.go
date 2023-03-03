package main

import (
	"log"
	"net/http"
	"time"
)

// timeHandler is a function that accepts a time format as string and
// returns a http.HanlderFunc can be converted to an http.Handler.
func timeHandler(format string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tm := time.Now().Format(format)
		w.Write([]byte("The time is: " + tm))
	}
}

// Leverage mux ability to accept a function that can be converted to an http.Handler,
// thanks to mux.HandleFunc.
func main() {
	mux := http.NewServeMux()

	th := timeHandler(time.RFC1123)
	mux.HandleFunc("/time", th)

	log.Print("Listening...")
	http.ListenAndServe(":3000", mux)
}

// Variants:
//
// Converting the closure to an http.Handler by ourselves:
//
//func timeHandler(format string) http.Handler {
//	return http.Handler(func(w http.ResponseWriter, r *http.Request) {
//		tm := time.Now().Format(format)
//		w.Write([]byte("The time is: " + tm))
//	})
//}
//
//func main() {
//	mux := http.NewServeMux()
//
//	th := timeHandler(time.RFC1123)
//	mux.Handle("/time", th)
//
//	log.Print("Listening...")
//	http.ListenAndServe(":3000", mux)
//}
//
// Using the default global http Default ServeMux:
// Note: pay attention as it's global and any package is able to access it and register a route on.
//
//func main() {
//	th := timeHandler(time.RFC1123)
//
//	// Use the http package's Default ServeMux
//	http.Handle("/time", th)
//
//	log.Print("Listening...")
//	http.ListenAndServe(":3000", mux)
//}
//
// Define a custom type that adheres to the http.Handler interface:
// ```
// type Handler interface {
//     ServeHTTP(ResponseWriter, *Request)
// }
// ```
//
//type timeHandler struct {
//	format string
//}
//
//func (th timeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
//	tm := time.Now().Format(th.format)
//	w.Write([]byte("The time is: " + tm))
//}
//
//func main() {
//	mux := http.NewServeMux()
//
//	th := timeHandler{format: time.RFC1123}
//
//	// Use the http package's Default ServeMux
//	mux.Handle("/time", th)
//
//	log.Print("Listening...")
//	http.ListenAndServe(":3000", mux)
//}
