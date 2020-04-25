package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/andrewmak/crud_golang/application/repositories"
	"github.com/andrewmak/crud_golang/domain"
	"github.com/andrewmak/crud_golang/framework/utils"
	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2/bson"
)

var BookRepositoryDb = repositories.BookRepositoryDb{}
var config = utils.Config{}

func init() {
	config.Read()

	BookRepositoryDb.Server = config.Server
	BookRepositoryDb.Database = config.Database
	BookRepositoryDb.Connect()
}
func listBooks(w http.ResponseWriter, r *http.Request) {
	books, err := BookRepositoryDb.GetAll()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, books)
}

func Create(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var book *domain.Book
	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	book.Id = bson.NewObjectId()
	if err := BookRepositoryDb.Insert(book); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusCreated, book)
}
func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Bem Vindo")
}

func configureRoutes() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", home).Methods("GET")
	r.HandleFunc("/books", listBooks).Methods("GET")
	r.HandleFunc("/books", Create).Methods("POST")
	return r
}

func MiddlewareServer() {
	r := configureRoutes()

	fmt.Println("Server running in port 1337")

	err := http.ListenAndServe(":1337", r)

	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	MiddlewareServer()
}

func respondWithError(w http.ResponseWriter, code int, msg string) {
	respondWithJson(w, code, map[string]string{"error": msg})
}

func respondWithJson(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
