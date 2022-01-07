package http

import (
	"net/http"
	"webserver/app/entity"
)

type Reader interface {
	GetCategory(id entity.ID) (*entity.Category, error)
	SearchCategories(query string) ([]*entity.Category, error)
	ListCategories() ([]*entity.Category, error)
}

//Writer Category writer
type Writer interface {
	CreateCategory(name string, description string) (entity.ID, error)
	UpdateCategory(e *entity.Category) error
	DeleteCategory(id entity.ID) error
}

//Repository interface
type Repository interface {
	Reader
	Writer
}

//UseCase interface
type ICategoryHandler interface {
	GetCategory(id entity.ID) (*entity.Category, error)
	SearchCategories(w http.ResponseWriter, r *http.Request)
	ListCategories() ([]*entity.Category, error)
	CreateCategory(name string, description string) (entity.ID, error)
	CreateCategories() (string, error)
	UpdateCategory(e *entity.Category) error
	DeleteCategory(id entity.ID) error
}
