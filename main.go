package main

import (
	"net/http"
)

func main() {

	mux := http.NewServeMux()
	mux.Handle("/", http.HandlerFunc(defaultHandler))
	// 使用logging中间件
	mux.Handle("/add", http.Handler(loggingHandler(http.HandlerFunc(addHandler))))
	// 使用auth中间件
	mux.Handle("/update", http.Handler(authHandler(http.HandlerFunc(updateHandler))))
	// 使用logging，auth两种中间件
	mux.Handle("/get", http.Handler(loggingHandler(http.Handler(authHandler(http.HandlerFunc(findHandler))))))
	mux.HandleFunc("/delete", http.HandlerFunc(deleteHandler))
	mux.Handle("/all",http.HandlerFunc(findAllHandler))

	http.ListenAndServe(":8000", mux)

}
