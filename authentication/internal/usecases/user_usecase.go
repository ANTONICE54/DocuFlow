package usecases

import (
	"auth_service/internal/apperrors"
	"auth_service/internal/models"
	"fmt"
)

type IUserRepo interface {
	Create(user models.User) (*models.User, error)
	GetByEmail(email string) (*models.User, error)
}

type IPasswordHasher interface {
	HashPassword(password string) (string, error)
	VerifyPassword(password string, hashedPassword string) error
}

type ITokenMaker interface {
	GenerateToken(userID int) (string, error)
	VerifyToken(tokenString string) error
	ExtractClaims(tokenString string) (*uint, error)
}

type UserUC struct {
	userRepo       IUserRepo
	tokenMaker     ITokenMaker
	passwordHasher IPasswordHasher
}

func NewUserUC(repo IUserRepo, tokenM ITokenMaker, passwordH IPasswordHasher) *UserUC {
	return &UserUC{
		userRepo:       repo,
		tokenMaker:     tokenM,
		passwordHasher: passwordH,
	}
}

// TODO -- this function have to return token
func (uc *UserUC) Register(user models.RegisterUserRequest) (*string, error) {

	alreadyExists, err := uc.userRepo.GetByEmail(user.Email)

	if alreadyExists != nil {
		return nil, apperrors.NewError("User with such email already exists", apperrors.BadRequest, "User with such email already exists")
	}

	hashedPassword, err := uc.passwordHasher.HashPassword(user.Password)

	if err != nil {
		return nil, apperrors.ErrInternalServer(fmt.Sprintf("Failed to hash password + %s", err.Error()))
	}

	userForDB := models.User{
		Name:           user.Name,
		Surname:        user.Surname,
		Email:          user.Email,
		Country:        user.Country,
		HashedPassword: hashedPassword,
	}

	createdUser, err := uc.userRepo.Create(userForDB)

	if err != nil {
		return nil, err
	}

	token, err := uc.tokenMaker.GenerateToken(int(createdUser.ID))

	return &token, nil
}

func (uc *UserUC) Login(loginInfo models.LoginUserRequest) (*string, error) {

	user, _ := uc.userRepo.GetByEmail(loginInfo.Email)
	/*
		if err != nil {
			return err
		}
	*/

	if user == nil {
		return nil, apperrors.NewError("Wrong email", apperrors.Unauthorized, "Wrong email or password")
	}

	err := uc.passwordHasher.VerifyPassword(loginInfo.Password, user.HashedPassword)

	if err != nil {
		return nil, apperrors.NewError("Wrong email", apperrors.Unauthorized, "Wrong email or password")
	}

	token, err := uc.tokenMaker.GenerateToken(int(user.ID))

	if err != nil {
		return nil, apperrors.ErrInternalServer("Failed to generate token")
	}

	return &token, nil
}

func (uc *UserUC) Verify(verifyInfo models.VerifyRequest) (*uint, error) {
	err := uc.tokenMaker.VerifyToken(verifyInfo.Token)

	if err != nil {
		return nil, err
	}

	userID, err := uc.tokenMaker.ExtractClaims(verifyInfo.Token)

	if err != nil {
		return nil, err
	}
	return userID, nil
}
