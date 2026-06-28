package db

import (
	"context"
	"errors"
	"fmt"
	"library/pkg/model/book"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type pgLibraryRepository struct {
	pool *pgxpool.Pool
}

func (pg *pgLibraryRepository) Create(ctx context.Context, b *book.Book) error {
	createQuery := `INSERT INTO books (title, description, created_year, total_pages, genre) 
	VALUES ($1, $2, $3, $4, $5)
	RETURNING created_at, updated_at;
	`

	err := pg.pool.QueryRow(
		ctx,
		createQuery,
		b.Title,
		b.Description,
		b.CreateYear,
		b.TotalPages,
		b.Genre,
	).Scan(&b.CreatedAt, &b.UpdatedAt)

	if err != nil {
		return fmt.Errorf("failed to execute create book query: %w", err)
	}

	return nil
}

func (pg *pgLibraryRepository) AddPage(ctx context.Context, p *book.Page) error {
	addQuery := `
	INSERT INTO book_pages (book_id, number, content) 
	VALUES ($1, $2, $3)
	RETURNING created_at, updated_at;
	`

	err := pg.pool.QueryRow(
		ctx,
		addQuery,
		p.BookID,
		p.Number,
		p.Content,
	).Scan(&p.CreatedAt, &p.UpdatedAt)

	if err != nil {
		return fmt.Errorf("failed to execute add page to book query: %w", err)
	}

	updateQuery := `
	UPDATE books
	SET total_pages = (SELECT COUNT(*) FROM book_pages WHERE book_id = $1)
	SET updated_at = NOW()
	WHERE id = $1
	`

	_, err = pg.pool.Exec(
		ctx,
		updateQuery,
		p.BookID,
	)

	if err != nil {
		return fmt.Errorf("failed to update total pages number for book: %w", err)
	}

	return nil
}

func (pg *pgLibraryRepository) GetByTitle(ctx context.Context, title string) (*book.Book, error) {
	getQuery := `SELECT * FROM books WHERE title = $1`

	var b book.Book
	err := pg.pool.QueryRow(
		ctx,
		getQuery,
	).Scan(
		&b.ID,
		&b.Title,
		&b.Description,
		&b.CreateYear,
		&b.TotalPages,
		&b.Genre,
		&b.CreatedAt,
		&b.UpdatedAt,
	)

	if err != nil {
		return nil, fmt.Errorf("failed to execute get book by title query: %w", err)
	}

	return &b, nil
}

func (pg *pgLibraryRepository) GetAll(ctx context.Context) ([]book.Book, error) {
	getQuery := `SELECT * FROM books`

	rows, err := pg.pool.Query(ctx, getQuery)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, fmt.Errorf("no one books not found")
		}

		return nil, fmt.Errorf("failed to execute get all books query: %w", err)
	}
	defer rows.Close()

	var books []book.Book
	for rows.Next() {
		var b book.Book
		rows.Scan(
			&b.ID,
			&b.Title,
			&b.Description,
			&b.CreateYear,
			&b.TotalPages,
			&b.Genre,
			&b.CreatedAt,
			&b.UpdatedAt,
		)

		books = append(books, b)
	}

	return books, nil
}

func (pg *pgLibraryRepository) GetPage(ctx context.Context, bookID uuid.UUID, pageNumber uint) (*book.Page, error) {
	getQuery := `SELECT * FROM book_pages WHERE book_id = $1, number = $1`

	var p book.Page
	err := pg.pool.QueryRow(
		ctx,
		getQuery,
		bookID,
		pageNumber,
	).Scan(
		&p.ID,
		&p.BookID,
		&p.Number,
		&p.Content,
		&p.CreatedAt,
		&p.UpdatedAt,
	)

	if err != nil {
		return nil, fmt.Errorf("failed to execute get page query: %w", err)
	}

	return &p, nil
}

func (pg *pgLibraryRepository) Delete(ctx context.Context, bookID uuid.UUID) error {
	deleteQuery := `DELETE FROM books WHERE id = $1`

	_, err := pg.pool.Exec(ctx, deleteQuery)
	if err != nil {
		return fmt.Errorf("failed to execute delete book query: %w", err)
	}

	return nil
}
