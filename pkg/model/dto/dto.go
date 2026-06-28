package dto

import (
	"library/pkg/model/book"

	"github.com/google/uuid"
)

type CreateBookRequest struct {
	Title       string      `json:"title" binding:"required"`
	Description string      `json:"description"`
	Pages       []book.Page `json:"pages" binding:"required"`
	Year        uint        `json:"year" binding:"required"`
	Genre       string      `json:"genre" binding:"required"`
}

type CreateBookResponse struct {
	Book book.Book `json:"book"`
}

type AddPageRequest struct {
	BookID  uuid.UUID `json:"book_id"`
	Content string    `json:"content"`
	Number  uint      `json:"number"`
}

type GetBookRequest struct {
	Title string `json:"title" binding:"required"`
}

type GetBookResponse struct {
	Book book.Book `json:"book"`
}

type GetAllBookResponse struct {
	Books []book.Book `json:"books"`
}

type GetPageRequest struct {
	BookID uuid.UUID `json:"uuid" binding:"required"`
	Number uint      `json:"page_number" binding:"required"`
}

type GetPageResponse struct {
	Page book.Page `json:"page"`
}

type ReadBookRequest struct {
	Title      string `json:"title" binding:"required"`
	PageNumber uint   `json:"page_number" binding:"required"`
}

type ReadBookResponse struct {
	Title       string `json:"title"`
	PageContent string `json:"page_content"`
}

type UpdateTitleRequest struct {
	BookID uuid.UUID `json:"-"`
	Title  string    `json:"title" binding:"required"`
}

type UpdateDescriptionRequest struct {
	BookID      uuid.UUID `json:"-"`
	Description string    `json:"description" binding:"required"`
}

type UpdateGenreRequest struct {
	BookID uuid.UUID `json:"-"`
	Genre  string    `json:"genre" binding:"required"`
}
