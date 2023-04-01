package repository

import (
	"book/model"
	"book/repository/query"
	"errors"
)

// interface employee
type BookRepo interface {
	CreateBook(in model.Book) (res model.Book, err error)
	GetBookID(id string) (res model.Book, err error)
	GetBookAll() (res []model.Book, err error)
	DeleteBook(id string) (err error)
	UpdateBook(id string, book model.Book) (model.Book, error)
}

func (r Repo) CreateBook(in model.Book) (res model.Book, err error) {
	r.db.QueryRow(
		query.AddBook,
		in.Title,
		in.Author,
		in.Decs,
	).Scan(
		&res.BookID,
		&res.Title,
		&res.Author,
		&res.Decs,
	)
	if err != nil {
		return res, err
	}

	return res, nil
}

func (r *Repo) GetBookID(id string) (res model.Book, err error) {
	r.db.QueryRow(query.GetBookID+id).Scan(
		&res.BookID,
		&res.Title,
		&res.Author,
		&res.Decs,
	)
	if err != nil {
		return res, err
	}
	return res, nil
}

func (r *Repo) GetBookAll() (res []model.Book, err error) {

	rows, err := r.db.Query(query.GetBookAll)
	if err != nil {
		return res, err
	}
	defer rows.Close()
	for rows.Next() {
		var book model.Book
		err := rows.Scan(
			&book.BookID,
			&book.Title,
			&book.Author,
			&book.Decs,
		)
		if err != nil {
			return res, err
		}
		res = append(res, book)
	}
	return res, nil
}

func (r *Repo) DeleteBook(id string) (err error) {
	res, err := r.db.Exec(query.DeleteBook + id)
	if err != nil {
		return err
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return errors.New("book not found")
	}
	return nil
}
func (r *Repo) UpdateBook(id string, book model.Book) (model.Book, error) {
	res, err := r.db.Exec(query.UpdateBook+id, book.Title, book.Author, book.Decs)

	if err != nil {
		return model.Book{}, err
	}
	rowsAffected, err := res.RowsAffected()

	if err != nil {
		return model.Book{}, err
	}
	if rowsAffected == 0 {
		return model.Book{}, errors.New("book not found")
	}
	return book, nil
}
