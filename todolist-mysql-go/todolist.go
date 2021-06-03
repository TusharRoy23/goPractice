package main

import (
	"encoding/json"
	"io"
	"net/http"

	"strconv"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/rs/cors"
	log "github.com/sirupsen/logrus"
)

var db, _ = gorm.Open("mysql", "root:root@/todolist?charset=utf8&parseTime=True&loc=Local")

type todoItemModel struct {
	ID          int `gorm:"primary_key"`
	Description string
	Completed   bool
}

func init() {
	log.SetFormatter(&log.TextFormatter{})
	log.SetReportCaller(true)
}

func healthz(res http.ResponseWriter, req *http.Request) {
	log.Info("API health is OK")
	res.Header().Set("Content-Type", "application/json")
	io.WriteString(res, `{alive: true}\n`)
}

func createItem(res http.ResponseWriter, req *http.Request) {
	description := req.FormValue("description")
	log.WithFields(log.Fields{
		"description": description,
	}).Info("Add new TodoItem. Saving to database")
	todo := &todoItemModel{Description: description, Completed: false}
	db.Create(&todo)
	result := db.Last(&todo)
	res.Header().Set("Content-Type", "application/json")
	json.NewEncoder(res).Encode(result.Value)
}

func updateItem(res http.ResponseWriter, req *http.Request) {
	// GET URL parameter from mux
	vars := mux.Vars(req)
	id, _ := strconv.Atoi(vars["id"])
	// Check if the Todoitem exist in DB or not

	err := getItemByID(id)
	if !err {
		res.Header().Set("Content_Type", "application/json")
		io.WriteString(res, `{"updated": false, "error": "Record Not Found"}`)
	} else {
		completed, _ := strconv.ParseBool(req.FormValue("completed"))
		log.WithFields(log.Fields{"ID": id, "Completed": completed}).Info("Updating TodoItem")
		todo := &todoItemModel{}
		db.First(&todo, id)
		todo.Completed = completed
		db.Save(&todo)
		res.Header().Set("Content-Type", "application/json")
		io.WriteString(res, `{"updated": true}`)
	}
}

func deleteItem(res http.ResponseWriter, req *http.Request) {
	// GET URL parameter from mux
	vars := mux.Vars(req)
	id, _ := strconv.Atoi(vars["id"])

	// Check if the TodoItem exist or not
	err := getItemByID(id)
	if !err {
		res.Header().Set("Content-Type", "application/json")
		io.WriteString(res, `{deleted: false, "error": "Record Not Found"}`)
	} else {
		log.WithFields(log.Fields{"ID": id}).Info("Deleting TodoItem")
		todo := &todoItemModel{}
		db.First(&todo, id)
		db.Delete(&todo)
		res.Header().Set("Content-Type", "application/json")
		io.WriteString(res, `{"deleted": true}`)
	}
}

func getItemByID(ID int) bool {
	todo := &todoItemModel{}
	result := db.First(&todo, ID)
	if result.Error != nil {
		log.Warn("TodoItem not found in Database")
		return false
	}
	return true
}

func getCompletedItems(res http.ResponseWriter, req *http.Request) {
	log.Info("Get Completed TodoItems")
	completedTodoItems := getTodoItems(true)
	res.Header().Set("Content-Type", "application/json")
	json.NewEncoder(res).Encode(completedTodoItems)
}

func getIncompletedItems(res http.ResponseWriter, req *http.Request) {
	log.Info("Get Incomplete TodoItems")
	inCompletedTodoItems := getTodoItems(false)
	res.Header().Set("Content-Type", "application/json")
	json.NewEncoder(res).Encode(inCompletedTodoItems)
}

func getTodoItems(completed bool) interface{} {
	var todos []todoItemModel
	TodoItems := db.Where("completed = ?", completed).Find(&todos).Value
	return TodoItems
}

func main() {
	defer db.Close()
	db.Debug().DropTableIfExists(&todoItemModel{})
	db.Debug().AutoMigrate(&todoItemModel{})

	log.Info("Starting Todolist API server")
	router := mux.NewRouter()
	router.HandleFunc("/healthz", healthz).Methods("GET")
	router.HandleFunc("/todo", createItem).Methods("POST")
	router.HandleFunc("/todo-completed", getCompletedItems).Methods("GET")
	router.HandleFunc("/todo-incomplete", getIncompletedItems).Methods("GET")
	router.HandleFunc("/todo/{id}", updateItem).Methods("PUT")
	router.HandleFunc("/todo/{id}", deleteItem).Methods("DELETE")
	// http.ListenAndServe(":8080", router)
	handle := cors.New(cors.Options{
		AllowedMethods: []string{"GET", "POST", "DELETE", "PUT", "OPTIONS"},
	}).Handler(router)

	http.ListenAndServe(":8080", handle)
}
