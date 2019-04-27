package main

import (
	"encoding/json"
	"fmt"
	"gopkg.in/mgo.v2"
	"log"
	"goji.io/pat"
	"net/http"
)

func ResponseWithJSON(w http.ResponseWriter, json []byte, code int) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(code)
	w.Write(json)
}

func ErrorWithJSON(w http.ResponseWriter, message string, code int) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(code)
	fmt.Fprintf(w, "{message: %q}", message)
}

func CompleteWithJSON(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "{message: %q}", "ok")
}

func defaultHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("in default controller")
}

func addHandler(w http.ResponseWriter, r *http.Request) {
	//fmt.Println("mongodb add something")
	var good Goods
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&good)
	if err != nil {
		ErrorWithJSON(w, "Incorrect body", http.StatusBadRequest)
		return
	}
	err = good.save();
	if err != nil {
		ErrorWithJSON(w, "Database error", http.StatusInternalServerError)
		log.Println("Failed insert good: ", err)
		return
	}

	CompleteWithJSON(w)

}

func findHandler(w http.ResponseWriter, r *http.Request) {
	//fmt.Println("mongodb search something")

	id := pat.Param(r, "id")
	var good Goods
	err := good.find(id)

	if err != nil {
		ErrorWithJSON(w, "Database error", http.StatusInternalServerError)
		log.Println("Failed find goods: ", err)
		return
	}

	if good.Name == ""{
		ErrorWithJSON(w, "good not found", http.StatusNotFound)
		return
	}

	respBody, err := json.MarshalIndent(good, "", "  ")
	if err != nil {
		log.Fatal(err)
	}

	ResponseWithJSON(w, respBody, http.StatusOK)
}

func updateHandler(w http.ResponseWriter, r *http.Request) {
	//fmt.Println("mongodb update something")
	//id := pat.Param(r, "id")

	var good Goods
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&good)
	if err != nil {
		ErrorWithJSON(w, "Incorrect body", http.StatusBadRequest)
		return
	}
	// 更新
	err = good.update()

	if err != nil {
		switch err {
		default:
			ErrorWithJSON(w, "Database error", http.StatusInternalServerError)
			log.Println("Failed update good: ", err)
			return
		case mgo.ErrNotFound:
			ErrorWithJSON(w, "good not found", http.StatusNotFound)
			return
		}
	}

	w.WriteHeader(http.StatusNoContent)

}

func deleteHandler(w http.ResponseWriter, r *http.Request) {
	//fmt.Println("mongodb delete something")
	var good Goods
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&good)
	if err != nil {
		ErrorWithJSON(w, "Incorrect body", http.StatusBadRequest)
		return
	}
	// 删除
	err = good.delete()

	if err != nil {
		switch err {
		default:
			ErrorWithJSON(w, "Database error", http.StatusInternalServerError)
			log.Println("Failed update good: ", err)
			return
		case mgo.ErrNotFound:
			ErrorWithJSON(w, "good not found", http.StatusNotFound)
			return
		}
	}

	w.WriteHeader(http.StatusNoContent)


}

func findAllHandler(w http.ResponseWriter, r *http.Request) {

	var goods []Goods
	goods, err := Goods.findAll(Goods{})
	if err != nil {
		ErrorWithJSON(w, "Database error", http.StatusInternalServerError)
		log.Fatal(err)
	}

	respBody, err := json.MarshalIndent(goods, "", "  ")
	if err != nil {
		log.Fatal(err)
	}

	ResponseWithJSON(w, respBody, http.StatusOK)

}
