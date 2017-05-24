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
	router.HandleFunc("/getAllTodos", GetAllTodos).Methods("GET")
	router.HandleFunc("/newTodo", CreateTodo).Methods("POST")

	log.Fatal(http.ListenAndServe(":80", router))
}

func Index(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Hey Brendan!")
}


func GetAllTodos(w http.ResponseWriter, req *http.Request) {

	var todoItems = querySql("SELECT id, value, checked FROM todo")
	todoJson, err := json.Marshal(todoItems)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
 	w.Write(todoJson)
}


func CreateTodo(w http.ResponseWriter, req *http.Request) {
	todoItems := DecodeTodoRequest(req)

	for _, todoItem := range todoItems {
		log.Print("Todo params are:" + todoItem.Value)
		executeSql("INSERT INTO todo (value) values ('" + todoItem.Value + "')")
	}
}

func DecodeTodoRequest(req *http.Request) (todoItems []Todo) {
	decoder := json.NewDecoder(req.Body)

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

func querySql(query string) (todoItems []Todo) {
	db, err := sql.Open("mysql", "root:bob@tcp(db:3306)/todo")
	if err != nil {
 		log.Panic(err);
		log.Fatal("Error: Connection to the DB messed up")
	}
	defer db.Close()

	rows, err := db.Query(query)
	if err != nil {
		log.Panic(err);
		log.Fatal("Error: SQL query messed up")
	}
	defer rows.Close()

	for rows.Next() {
		var todoItem Todo
		rows.Scan(&todoItem.ID, &todoItem.Value, &todoItem.Checked)
		todoItems = append(todoItems, todoItem)
	}

	if err != nil {
		log.Panic(err);
		log.Fatal("Error: SQL execution messed up")
	}

	return
}
