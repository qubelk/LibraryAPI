package dto

import (
	"library/pkg/model/book"
	"time"
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

type GetBookRequest struct {
	Title string `json:"title" binding:"required"`
}

type GetBookResponse struct {
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Year        uint      `json:"year"`
	PageCount   uint      `json:"page_count"`
	Genre       string    `json:"genre"`
	CreatedAt   time.Time `json:"created_at"`
}

type ReadBookRequest struct {
	Title      string `json:"title" binding:"required"`
	PageNumber uint   `json:"page_number" binding:"required"`
}

type ReadBookResponse struct {
	Title       string `json:"title"`
	PageContent string `json:"page_content"`
}
