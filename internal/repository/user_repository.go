package repository

import "github.com.rifkifajarramadhani/golang-backend-boilerplate-1/internal/domain"

type UserRepository interface {
	GetAll() (*[]domain.User, error)
	GetByID(id int) (*domain.User, error)
	Create(user *domain.User) (*domain.User, error)
	Update(id int, user *domain.User) (*domain.User, error)
	Delete(id int) error
}