package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type respond struct {
	Name  string
	Value int
}

type message struct {
	Command string
}

func handler(w http.ResponseWriter, r *http.Request) {
	var m message
	if r.Body == nil {
		http.Error(w, "Please send a request body", 400)
		return
	}
	err := json.NewDecoder(r.Body).Decode(&m)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	fmt.Println(m.Command)

	ret := respond{Name: "sum", Value: 42}

	js, err := json.Marshal(ret)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
