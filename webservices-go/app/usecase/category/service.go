package category

import (
	"fmt"
	"strings"
	"time"
	"webserver/app/entity"
)

//Service book usecase
type Service struct {
	repo Repository
}

//NewService create new service
func NewService(r Repository) *Service {
	return &Service{
		repo: r,
	}
}

//CreateBook create a book
func (s *Service) CreateCategory(name string, description string) (entity.ID, error) {
	b, err := entity.NewCategory(name, description)
	if err != nil {
		return entity.NewID(), err
	}
	return s.repo.Create(b)
}

//GetBook get a book
func (s *Service) GetCategory(id entity.ID) (*entity.Category, error) {
	b, err := s.repo.Get(id)
	if b == nil {
		return nil, entity.ErrNotFound
	}
	if err != nil {
		return nil, err
	}

	return b, nil
}

//SearchBooks search books
func (s *Service) SearchCategories(query string) ([]*entity.Category, error) {
	books, err := s.repo.Search(strings.ToLower(query))
	fmt.Print(len(books))
	if err != nil {
		return nil, err
	}
	if len(books) == 0 {
		return nil, entity.ErrNotFound
	}
	return books, nil
}

//ListBooks list books
func (s *Service) ListCategories() ([]*entity.Category, error) {
	books, err := s.repo.List()
	if err != nil {
		return nil, err
	}
	if len(books) == 0 {
		return nil, entity.ErrNotFound
	}
	return books, nil
}

//DeleteBook Delete a book
func (s *Service) DeleteCategory(id entity.ID) error {
	_, err := s.GetCategory(id)
	if err != nil {
		return err
	}
	return s.repo.Delete(id)
}

//UpdateBook Update a book
func (s *Service) UpdateCategory(e *entity.Category) error {
	err := e.Validate()
	if err != nil {
		return err
	}
	e.UpdatedAt = time.Now()
	return s.repo.Update(e)
}
