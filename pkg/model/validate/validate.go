package validate

import (
	"fmt"

	"golang.org/x/sync/errgroup"
)

func validateTitle(title string) error {
	if len(title) < 3 {
		return fmt.Errorf("title most be at least 3 character")
	}

	if len(title) > 255 {
		return fmt.Errorf("title can't be over than 255 characters")
	}

	return nil
}

func validateGenre(genre string) error {
	if genre == "" {
		return fmt.Errorf("genre can't be empty")
	}

	if len(genre) > 100 {
		return fmt.Errorf("genre can't be over than 100 characters")
	}

	return nil
}

func Validate(title, genre string) error {
	var g errgroup.Group

	g.Go(func() error {
		return validateTitle(title)
	})

	g.Go(func() error {
		return validateGenre(genre)
	})

	return g.Wait()
}
