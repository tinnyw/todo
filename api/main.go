package main

import (
	"fmt"
	// "html"
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"encoding/json"
)

type Todo  struct {
	ID int `json:"id"`
	Value string `json:"value"`
	Checked bool `json:"checked"`
}

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
	todoItems := DecodeTodoRequest(req)
	// var todoItem Todo
	// _ = json.NewDecoder(req.Body).Decode(&todoItem)
	for _, todoItem := range todoItems {
		log.Print("Todo params are:" + todoItem.Value)
		executeSql("INSERT INTO todo (value) values ('" + todoItem.Value + "')")
	}
}

func DecodeTodoRequest(req *http.Request) (todoItems []Todo) {
	decoder := json.NewDecoder(req.Body)
	/*var todoItems = new (Todo)
  err := json.Unmarshal(req.Body, &todoItems)*/

	err := decoder.Decode(&todoItems)
	if err != nil {
 		log.Panic(err)
	}

	return
}

func executeSql(stmnt string) {
	db, err := sql.Open("mysql", "root:bob@tcp(db:3306)/todo")
	if err != nil {
 		log.Panic(err);
		log.Fatal("Error: Connection to the DB messed up")
	}
	defer db.Close()

	_, err = db.Exec(stmnt)

	if err != nil {
		log.Panic(err);
		log.Fatal("Error: SQL execution messed up")
	}
}
