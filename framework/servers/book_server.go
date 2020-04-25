package servers

import (
	"encoding/json"
	"net/http"

	"github.com/andrewmak/crud_golang/application/usecases"
	"github.com/andrewmak/crud_golang/domain"
	"gopkg.in/mgo.v2/bson"
)

type BookServer struct {
	BookUseCase usecases.BookUseCase
}

func NewBookServer() *BookServer {
	return &BookServer{}
}

func (BookServer *BookServer) CreateBook(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var book *domain.Book
	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	book.Id = bson.NewObjectId()

	err := BookServer.BookUseCase.Create(book)

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusCreated, book)

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
