package category

import (
	"webserver/app/entity"
)

//Reader interface
type Reader interface {
	Get(id entity.ID) (*entity.Category, error)
	Search(query string) ([]*entity.Category, error)
	List() ([]*entity.Category, error)
}

//Writer book writer
type Writer interface {
	Create(e *entity.Category) (entity.ID, error)
	Update(e *entity.Category) error
	Delete(id entity.ID) error
}

//Repository interface
type Repository interface {
	Reader
	Writer
}

//UseCase interface
type UseCase interface {
	GetCategory(id entity.ID) (*entity.Category, error)
	SearchCategories(query string) ([]*entity.Category, error)
	ListCategories() ([]*entity.Category, error)
	CreateCategory(title string, author string, pages int, quantity int) (entity.ID, error)
	UpdateCategory(e *entity.Category) error
	DeleteCategory(id entity.ID) error
}
