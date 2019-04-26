package main

import (
	"net/http"
)

func main() {

	mux := &ControllerRegistor{}

	mux.Add("/", &DefaultController{})
	mux.Add("/add", &AddController{})
	mux.Add("/delete", &DeleteController{})
	mux.Add("/search", &SearchController{})
	mux.Add("/update", &UpdateController{})

	s := &http.Server{
		Addr:    ":8001",
		Handler: mux,
	}

	s.ListenAndServe()
}

