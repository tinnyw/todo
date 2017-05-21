package main

import (
	"fmt"
	// "html"
	"log"
	"net/http"
	"github.com/gorilla/mux"
)

func main() {
	/*http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "LOGOS, i <3 YOU!!!!")
	})*/
	router := mux.NewRouter()
	router.HandleFunc("/", Index).Methods("GET")

	log.Fatal(http.ListenAndServe(":80", router))
}

func Index(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Hey Brendan!")
}
