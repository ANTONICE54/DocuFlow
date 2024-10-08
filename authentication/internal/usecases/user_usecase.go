package usecases

import "auth_service/internal/models"

type IUserRepo interface {
	Create(user models.User) error
	GetByEmail(email string) (*models.User, error)
}

type UserUC struct {
	userRepo IUserRepo
}

func NewUserUC(repo IUserRepo) *UserUC {
	return &UserUC{
		userRepo: repo,
	}
}
