package service

import (
	"context"
	"fmt"
	"library/pkg/db"
	"library/pkg/model/book"
	"library/pkg/model/dto"
	"library/pkg/model/validate"

	"github.com/google/uuid"
)

type LibraryService struct {
	repo db.LibraryRepository
}

func New(repo db.LibraryRepository) *LibraryService {
	return &LibraryService{
		repo: repo,
	}
}

func (ls *LibraryService) Create(ctx context.Context, req *dto.CreateBookRequest) (*dto.CreateBookResponse, error) {
	if err := validate.Validate(req.Title, req.Genre); err != nil {
		return nil, fmt.Errorf("failed to create book: %w", err)
	}

	b := book.New(req.Title, req.Description, req.Genre, req.Year)
	if err := ls.repo.Create(ctx, b); err != nil {
		return nil, fmt.Errorf("failed to create book: %w", err)
	}

	return &dto.CreateBookResponse{
		Book: *b,
	}, nil
}

func (ls *LibraryService) AddPage(ctx context.Context, req *dto.AddPageRequest) error {
	p := book.Page{
		ID:      uuid.New(),
		BookID:  req.BookID,
		Content: req.Content,
		Number:  req.Number,
	}

	if err := ls.repo.AddPage(ctx, &p); err != nil {
		return fmt.Errorf("failed to add page: %w", err)
	}

	return nil
}

func (ls *LibraryService) GetByTitle(ctx context.Context, req *dto.GetBookRequest) (*dto.GetBookResponse, error) {
	if err := validate.Validate(req.Title, "tmp"); err != nil {
		return nil, fmt.Errorf("failed to validate get by title request: %w", err)
	}

	b, err := ls.repo.GetByTitle(ctx, req.Title)
	if err != nil {
		return nil, fmt.Errorf("failed to get book by title: %w", err)
	}

	return &dto.GetBookResponse{
		Book: *b,
	}, nil
}

func (ls *LibraryService) GetAll(ctx context.Context) (*dto.GetAllBookResponse, error) {
	books, err := ls.repo.GetAll(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get all books: %w", err)
	}

	return &dto.GetAllBookResponse{
		Books: books,
	}, nil
}

func (ls *LibraryService) GetPage(ctx context.Context, req *dto.GetPageRequest) (*dto.GetPageResponse, error) {
	p, err := ls.repo.GetPage(ctx, req.BookID, req.Number)
	if err != nil {
		return nil, fmt.Errorf("failed to get page for book: %w", err)
	}

	return &dto.GetPageResponse{
		Page: *p,
	}, nil
}

func (ls *LibraryService) UpdateTitle(ctx context.Context, req *dto.UpdateTitleRequest) error {
	if err := validate.Validate(req.Title, "tmp"); err != nil {
		return fmt.Errorf("failed to validate request: %w", err)
	}

	return ls.repo.UpdateTitle(ctx, req.BookID, req.Title)
}

func (ls *LibraryService) UpdateDescription(ctx context.Context, req *dto.UpdateDescriptionRequest) error {
	return ls.repo.UpdateDescription(ctx, req.BookID, req.Description)
}

func (ls *LibraryService) UpdateGenre(ctx context.Context, req *dto.UpdateGenreRequest) error {
	if err := validate.Validate("tmp", req.Genre); err != nil {
		return fmt.Errorf("failed to validate request: %w", err)
	}

	return ls.repo.UpdateGenre(ctx, req.BookID, req.Genre)
}

func (ls *LibraryService) Delete(ctx context.Context, bookID uuid.UUID) error {
	if err := ls.repo.Delete(ctx, bookID); err != nil {
		return fmt.Errorf("failed to delete book: %w", err)
	}

	return nil
}
