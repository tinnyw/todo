package main

import (
	"fmt"
	// "html"
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	/*http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "LOGOS, i <3 YOU!!!!")
	})*/
	router := mux.NewRouter()
	router.HandleFunc("/", Index).Methods("GET")
	router.HandleFunc("/newTodo", CreateTodo).Methods("POST")

	log.Fatal(http.ListenAndServe(":80", router))
}

func Index(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Hey Brendan!")
}


func CreateTodo(w http.ResponseWriter, req *http.Request) {
	// db, err := sql.Open("mysql", "root:bob@/todo"); 
	db, err := sql.Open("mysql", "root:bob@tcp(db:3306)/todo")
	if err != nil {
		log.Panic(err);
		log.Fatal("Error: Connection to the DB messed up 1")
	} else {
		log.Print("Cool DB 1")
	}

	err = db.Ping()
	if err != nil {
		log.Panic(err);
 		log.Fatal("Error: Connection to the DB messed up 2")
 	} else {
		log.Print("Cool DB 2")
	}
}

