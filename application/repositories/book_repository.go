package repositories

import (
	"log"

	"github.com/andrewmak/crud_golang/domain"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type BookRepository interface {
	GetAll() ([]*domain.Book, error)
	Insert(book *domain.Book) error
}

type BookRepositoryDb struct {
	Server   string
	Database string
}

var db *mgo.Database

const (
	COLLECTION = "books"
)

func (m *BookRepositoryDb) Connect() {
	session, err := mgo.Dial(m.Server)
	if err != nil {
		log.Fatal(err)
	}
	db = session.DB(m.Database)
}

func (m *BookRepositoryDb) GetAll() ([]*domain.Book, error) {
	var books []*domain.Book
	err := db.C(COLLECTION).Find(bson.M{}).All(&books)
	return books, err
}

func (repo *BookRepositoryDb) Insert(book *domain.Book) error {
	err := db.C(COLLECTION).Insert(&book)
	return err
}
