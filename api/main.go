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
	"strconv"
)

type Todo  struct {
	ID int `json:"id"`
	Value string `json:"value"`
	Checked bool `json:"checked"`
}

// handle all the different API routes
func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", Index).Methods("GET")
	router.HandleFunc("/getAllTodos", GetAllTodos).Methods("GET")
	router.HandleFunc("/createTodos", CreateTodos).Methods("POST")
	router.HandleFunc("/checkTodos", CheckTodos).Methods("POST")
	router.HandleFunc("/uncheckTodos", UncheckTodos).Methods("POST")

	log.Fatal(http.ListenAndServe(":80", router))
}

// main page for the API service
func Index(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Hey Brendan!")
}

// get all the todos
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

// create a list of new todos
func CreateTodos(w http.ResponseWriter, req *http.Request) {
	todoItems := DecodeTodoRequest(req)

	for _, todoItem := range todoItems {
		log.Print("Todo params are:" + todoItem.Value)
		executeSql("INSERT INTO todo (value) values ('" + todoItem.Value + "')")
	}
}

func CheckTodos(w http.ResponseWriter, req *http.Request) {
	todoItemsToCheck := DecodeTodoRequest(req)

	for _, todoItemToCheck := range todoItemsToCheck {
		log.Print("Todo params are:" + todoItemToCheck.Value)
		executeSql("UPDATE todo SET checked = 1 WHERE id = " + strconv.Itoa(todoItemToCheck.ID))
	}
}

func UncheckTodos(w http.ResponseWriter, req *http.Request) {
	todoItemsToUncheck := DecodeTodoRequest(req)

	for _, todoItemToUncheck := range todoItemsToUncheck {
		log.Print("Todo params are:" + todoItemToUncheck.Value)
		executeSql("UPDATE todo SET checked = 0 WHERE id = " + strconv.Itoa(todoItemToUncheck.ID))
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
