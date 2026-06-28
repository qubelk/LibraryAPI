package book

import (
	"time"

	"github.com/google/uuid"
)

type Page struct {
	ID        uuid.UUID `json:"id"`
	BookID    uuid.UUID `json:"book_id"`
	Content   string    `json:"content"`
	Number    uint      `json:"number"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Book struct {
	ID          uuid.UUID `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	CreateYear  uint      `json:"create_year"`
	Pages       []Page    `json:"pages"`
	TotalPages  uint      `json:"total_pages"`
	Genre       string    `json:"genre"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func New(title, description, genre string, year uint) *Book {
	return &Book{
		ID:          uuid.New(),
		Title:       title,
		Description: description,
		CreateYear:  year,
		Pages:       []Page{},
		TotalPages:  0,
		Genre:       genre,
	}
}
