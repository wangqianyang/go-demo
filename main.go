package main

import (
	"net/http"
)

func main() {

	mux := http.NewServeMux()
	mux.Handle("/", http.HandlerFunc(defaultHandler))
	mux.Handle("/add", http.Handler(loggingHandler(http.HandlerFunc(addHandler))))
	mux.Handle("/update", http.Handler(authHandler(http.HandlerFunc(updateHandler))))
	mux.Handle("/get", http.Handler(loggingHandler(http.Handler(authHandler(http.HandlerFunc(findHandler))))))
	mux.HandleFunc("/delete", http.HandlerFunc(deleteHandler))
	mux.Handle("/all",http.HandlerFunc(findAllHandler))

	http.ListenAndServe(":8000", mux)

}
