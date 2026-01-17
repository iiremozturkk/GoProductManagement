package service

import (
	"GoProductManagement/models"
	"GoProductManagement/repository"
)

type UserService interface {
	Register(user models.User) (int, error)
	FindByEmail(email string) (models.User, error)
}

type userService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{repo: repo}
}

func (s *userService) Register(user models.User) (int, error) {
	return s.repo.Create(user)
}

func (s *userService) FindByEmail(email string) (models.User, error) {
	return s.repo.FindByEmail(email)
} 