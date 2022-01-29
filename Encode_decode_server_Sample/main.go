package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type person struct {
	First string
}

func main() {
	http.HandleFunc("/encode", encodefunc)
	http.HandleFunc("/decode", decodefunc)
	http.ListenAndServe(":8080", nil)
}

func encodefunc(w http.ResponseWriter, r *http.Request) {
	println("you are in encode function")
	p1 := person{
		First: "Jagadeesh",
	}
	err := json.NewEncoder(w).Encode(p1)
	if err != nil {
		log.Println("Encoded bad data :", err)
	}
}

func decodefunc(w http.ResponseWriter, r *http.Request) {
	println("you are in decode function")
	var p1 person
	err := json.NewDecoder(r.Body).Decode(&p1)
	if err != nil {
		log.Println("Decoded bad data :", err)
	}

	log.Println("Person :", p1)
}
