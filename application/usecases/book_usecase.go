package usecases

import (
	"github.com/andrewmak/crud_golang/application/repositories"
	"github.com/andrewmak/crud_golang/domain"
)

type BookUseCase struct {
	BookRepositoryDb repositories.BookRepositoryDb
}

func (u *BookUseCase) Create(book *domain.Book) error {

	err := u.BookRepositoryDb.Insert(book)

	if err != nil {
		return err
	}

	return nil
}