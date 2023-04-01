package query

const (
	AddBook = `
		INSERT INTO
			books
		(
			title,
			author,
			des
			
		)
		VALUES ($1, $2, $3)
		RETURNING *;
	`
)

const (
	GetBookID = `
		SELECT * FROM books WHERE id =
	`
)

const (
	GetBookAll = `
		SELECT * FROM books ORDER BY id ASC;
	`
)
const (
	DeleteBook = `
		DELETE FROM books WHERE id =
	`
)

const (
	UpdateBook = `
		UPDATE books SET title = $1, author = $2, des = $3 WHERE id =
	`
)
