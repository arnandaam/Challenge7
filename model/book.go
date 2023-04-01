package model

type Book struct {
	BookID int    `json:"id" db:"id"`
	Title  string `json:"title" db:"title"`
	Author string `json:"author" db:"autor"`
	Decs   string `json:"des" db:"des"`
}
