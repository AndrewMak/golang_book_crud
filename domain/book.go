package domain

import (
	"gopkg.in/mgo.v2/bson"
)

type Book struct {
	Id     bson.ObjectId `bson:"_id" json:"id"`
	Titulo string        `bson:"titutlo" json:"titulo"`
	Autor  string        `bson:"autor" json:"autor"`
	Active bool          `bson:"active" json:"active"`
}

func NewBook() *Book {
	return &Book{}
}

func (book *Book) Prepare() error {

	err := book.validate()

	if err != nil {
		return err
	}

	return nil
}

func (book *Book) validate() error {
	// validar o com govalidate
	return nil
}
