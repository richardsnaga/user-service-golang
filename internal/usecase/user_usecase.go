package usecase

import (
	"fmt"
	"user-service/internal/domain/user"
)

type UserUsecase struct {
	repo user.Repository
}

func NewUserUsecase(r user.Repository) *UserUsecase {
	return &UserUsecase{repo: r}
}

func (u *UserUsecase) GetUsers() ([]user.User, error) {
	return u.repo.FindAll()
}

func (u *UserUsecase) CreateUser(newUser *user.User) error {
	// contoh validasi
	if newUser.Email == "" {
		return fmt.Errorf("email is required")
	}
	return u.repo.Create(newUser)
}
