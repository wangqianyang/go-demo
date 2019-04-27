package main

import (
	"fmt"
	"net/http"
)

func defaultHandler(w http.ResponseWriter, r *http.Request){
	fmt.Println("in default controller")
}

func addHandler(w http.ResponseWriter, r *http.Request){
	fmt.Println("mongodb add something")
}

func searchHandler(w http.ResponseWriter, r *http.Request){
	fmt.Println("mongodb search something")
}

func updateHandler(w http.ResponseWriter, r *http.Request){
	fmt.Println("mongodb update something")
}

func deleteHandler(w http.ResponseWriter, r *http.Request){
	fmt.Println("mongodb delete something")

}