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


	//setup the DB
	executeSql("CREATE TABLE todo (ID INT(7) USIGNED AUTO_INCREMENT PRIMARY KEY, VALUE VARCHAR(40) NOT NULL, CHECKED BOOLEAN NOT NULL DEFAULT false)")	
	log.Fatal(http.ListenAndServe(":80", router))
}

func Index(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Hey Brendan!")
}


func CreateTodo(w http.ResponseWriter, req *http.Request) {
	executeSql("CREATE TABLE todo (ID INT(7) USIGNED AUTO_INCREMENT PRIMARY KEY, VALUE VARCHAR(40) NOT NULL, CHECKED BOOLEAN NOT NULL DEFAULT false)")
}

func executeSql(stmnt string) {
	db, err := sql.Open("mysql", "root:bob@tcp(db:3306)/todo")
	if err != nil {
 		log.Panic(err);
		log.Fatal("Error: Connection to the DB messed up")
	}
	defer db.Close()

	_, err = db.Exec(stmnt)
}

