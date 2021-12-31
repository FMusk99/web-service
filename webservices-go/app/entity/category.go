package entity

import (
	"time"
)

//Book data
type Category struct {
	ID          ID        `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

//NewBook create a new book
func NewCategory(name string, description string) (*Category, error) {
	b := &Category{
		ID:          NewID(),
		Name:        name,
		Description: description,
		CreatedAt:   time.Now(),
	}
	err := b.Validate()
	if err != nil {
		return nil, ErrInvalidEntity
	}
	return b, nil
}

//Validate validate book
func (b *Category) Validate() error {
	if b.Name == "" {
		return ErrInvalidEntity
	}
	return nil
}
