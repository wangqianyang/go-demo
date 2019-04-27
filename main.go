package main

import (
	"net/http"
)

func main() {

	mux := http.NewServeMux()
	mux.Handle("/", http.HandlerFunc(defaultHandler))
	mux.Handle("/add", http.Handler(loggingHandler(http.HandlerFunc(addHandler))))
	mux.Handle("/update", http.Handler(authHandler(http.HandlerFunc(updateHandler))))
	mux.Handle("/search", http.Handler(loggingHandler(http.Handler(authHandler(http.HandlerFunc(searchHandler))))))
	mux.HandleFunc("/delete", http.HandlerFunc(deleteHandler))

	http.ListenAndServe(":8000", mux)

}
