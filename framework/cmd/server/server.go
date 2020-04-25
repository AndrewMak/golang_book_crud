package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/andrewmak/crud_golang/application/repositories"
	"github.com/andrewmak/crud_golang/application/usecases"
	"github.com/andrewmak/crud_golang/framework/servers"
	"github.com/andrewmak/crud_golang/framework/utils"
	"github.com/gorilla/mux"
)

var BookRepositoryDb = repositories.BookRepositoryDb{}
var config = utils.Config{}

func listBooks(w http.ResponseWriter, r *http.Request) {
	books, err := BookRepositoryDb.GetAll()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, books)
}

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Bem Vindo")
}

func configureRoutes() *mux.Router {
	bookServer := setUpBookServer()
	r := mux.NewRouter()
	r.HandleFunc("/", home).Methods("GET")
	r.HandleFunc("/books", listBooks).Methods("GET")
	r.HandleFunc("/books", bookServer.CreateBook).Methods("POST")
	return r
}

func init() {
	config.Read()

	BookRepositoryDb.Server = config.Server
	BookRepositoryDb.Database = config.Database
	BookRepositoryDb.Connect()
}

func MiddlewareServer() {
	r := configureRoutes()

	var port = ":3000"
	fmt.Println("Server running in port:", port)
	err := http.ListenAndServe(port, r)
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

func setUpBookServer() *servers.BookServer {
	bookServer := servers.NewBookServer()
	bookRepository := repositories.BookRepositoryDb{}
	bookServer.BookUseCase = usecases.BookUseCase{BookRepositoryDb: bookRepository}
	return bookServer
}
