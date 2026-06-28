package db

import (
	"context"
	"library/pkg/model/book"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type LibraryRepository interface {
	Create(ctx context.Context, b *book.Book) error
	AddPage(ctx context.Context, p *book.Page) error
	GetByTitle(ctx context.Context, title string) (*book.Book, error)
	GetAll(ctx context.Context) ([]book.Book, error)
	GetPage(ctx context.Context, bookID uuid.UUID, pageNumber uint) (*book.Page, error)
	UpdateTitle(ctx context.Context, bookID uuid.UUID, title string) error
	UpdateDescription(ctx context.Context, bookID uuid.UUID, description string) error
	UpdateGenre(ctx context.Context, bookID uuid.UUID, genre string) error
	Delete(ctx context.Context, bookID uuid.UUID) error
}

func New(pool *pgxpool.Pool) LibraryRepository {
	return &pgLibraryRepository{
		pool: pool,
	}
}
