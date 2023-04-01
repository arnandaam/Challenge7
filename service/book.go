package service

import "book/model"

type BookService interface {
	CreateBook(in model.Book) (res model.Book, err error)
	GetBookID(id string) (res model.Book, err error)
	GetBookAll() (res []model.Book, err error)
	DeleteBook(id string) error
	UpdateBook(id string, book model.Book) (res model.Book, err error)
}

func (s *Service) CreateBook(in model.Book) (res model.Book, err error) {
	//cuma manggil dari repo
	res, err = s.repo.CreateBook(in)
	if err != nil {
		return res, err
	}

	return res, nil

	// return s.repo.CreateEmployee(in)
}

func (s *Service) GetBookID(id string) (res model.Book, err error) {
	return s.repo.GetBookID(id)
}

func (s *Service) GetBookAll() (res []model.Book, err error) {
	return s.repo.GetBookAll()
}

func (s *Service) DeleteBook(id string) error {
	return s.repo.DeleteBook(id)
}

func (s *Service) UpdateBook(id string, book model.Book) (res model.Book, err error) {
	return s.repo.UpdateBook(id, book)
}
