package usecases

import (
	"auth_service/internal/apperrors"
	"auth_service/internal/models"
	"fmt"
)

type IUserRepo interface {
	Create(user models.User) (*models.User, error)
	GetByEmail(email string) (*models.User, error)
	GetByID(userID uint) (*models.User, error)
	Update(userObj models.User) (*models.User, error)
	Delete(userID uint) error
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

type ICategoryUC interface {
	CreateDefaultCategories(userID uint)
}

type UserUC struct {
	userRepo       IUserRepo
	tokenMaker     ITokenMaker
	passwordHasher IPasswordHasher
	categoryUC     ICategoryUC
}

func NewUserUC(repo IUserRepo, tokenM ITokenMaker, passwordH IPasswordHasher, categoryUC ICategoryUC) *UserUC {
	return &UserUC{
		userRepo:       repo,
		tokenMaker:     tokenM,
		passwordHasher: passwordH,
		categoryUC:     categoryUC,
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

	uc.categoryUC.CreateDefaultCategories(createdUser.ID)

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

func (uc *UserUC) Get(userID uint) (*models.User, error) {

	return uc.userRepo.GetByID(userID)
}

func (uc *UserUC) Delete(userID uint) error {
	_, err := uc.userRepo.GetByID(userID)

	if err != nil {
		return err
	}

	return uc.userRepo.Delete(userID)
}

func (uc *UserUC) Update(updInfo models.UpdateUserRequest) (*models.User, error) {
	userBeforeUpd, err := uc.userRepo.GetByID(updInfo.ID)

	if err != nil {
		return nil, err
	}

	userUpdObj := models.User{
		ID:      userBeforeUpd.ID,
		Name:    userBeforeUpd.Name,
		Surname: userBeforeUpd.Surname,
		Email:   userBeforeUpd.Email,
		Country: userBeforeUpd.Country,
	}
	userUpdObj.ID = updInfo.ID

	if updInfo.Name != "" {
		userUpdObj.Name = updInfo.Name
	}

	if updInfo.Surname != "" {
		userUpdObj.Surname = updInfo.Surname
	}

	if updInfo.Email != "" {
		userUpdObj.Email = updInfo.Email
	}

	if updInfo.Country != "" {
		userUpdObj.Country = updInfo.Country
	}

	userAfterUpd, err := uc.userRepo.Update(userUpdObj)

	if err != nil {
		return nil, err
	}

	return userAfterUpd, nil
}
