package handler

import (
	"book/helper"
	"book/model"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h HttpServer) CreateBook(c *gin.Context) {
	in := model.Book{}

	err := c.BindJSON(&in)
	if err != nil {
		helper.BadRequest(c, err.Error())
		return
	}

	// err = in.Validation()
	// if err != nil {
	// 	helper.BadRequest(c, err.Error())
	// 	return
	// }

	// panggil servise
	res, err := h.app.CreateBook(in)
	if err != nil {
		helper.InternalServerError(c, err.Error())
		return
	}

	helper.Ok(c, res)
}
func (h HttpServer) GetBookID(c *gin.Context) {
	in := c.Param("id")
	res, err := h.app.GetBookID(in)

	if err != nil {
		helper.BadRequest(c, c.Err().Error())
	}
	// res, err := h.app.GetBookID(in)
	// if err != nil {
	// 	helper.InternalServerError(c, err.Error())
	// 	return
	// }

	helper.Ok(c, res)
}

func (h HttpServer) GetBookAll(c *gin.Context) {
	book, err := h.app.GetBookAll()

	if err != nil {
		helper.BadRequest(c, c.Err().Error())
	}
	// res, err := h.app.GetBookID(in)
	// if err != nil {
	// 	helper.InternalServerError(c, err.Error())
	// 	return
	// }

	helper.Ok(c, book)
}

func (h *HttpServer) DeleteBook(c *gin.Context) {
	id := c.Param("id")
	err := h.app.DeleteBook(id)
	if err != nil {
		helper.BadRequest(c, c.Err().Error())
	}
	c.JSON(http.StatusOK, "Deleted")
	helper.Ok(c, err)
}

func (h *HttpServer) UpdateBook(c *gin.Context) {
	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		helper.BadRequest(c, err.Error())
		return
	}

	// Parse the book data from the request body
	var book model.Book
	err = c.BindJSON(&book)
	if err != nil {
		helper.BadRequest(c, err.Error())
		return
	}

	// Set the book ID to the value from the request URL
	book.BookID = idInt

	// Call the service to update the book
	msg, err := h.app.UpdateBook(id, book)
	if err != nil {
		helper.InternalServerError(c, err.Error())
		return
	}

	// Respond with a success message
	c.JSON(http.StatusOK, msg)
}
