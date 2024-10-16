package database

import (
	"auth_service/internal/apperrors"
	"auth_service/internal/models"
	"context"
	"database/sql"
)

type UserRepo struct {
	*sql.DB
}

func NewUserRepo(db *sql.DB) *UserRepo {
	return &UserRepo{
		db,
	}
}

func (repo *UserRepo) Create(user models.User) (*models.User, error) {
	query := "INSERT INTO users(name, surname, email, country, hashed_password) values ($1, $2, $3, $4, $5) RETURNING id, name, surname, email, country, hashed_password, created_at;"

	ctx, cancel := context.WithTimeout(context.Background(), DBTimeout)

	defer cancel()

	row := repo.QueryRowContext(ctx, query, user.Name, user.Surname, user.Email, user.Country, user.HashedPassword)

	var res models.User

	err := row.Scan(
		&res.ID,
		&res.Name,
		&res.Surname,
		&res.Email,
		&res.Country,
		&res.HashedPassword,
		&res.CreatedAt,
	)

	if err != nil {
		return nil, apperrors.ErrDatabase(err.Error())
	}

	return &res, nil
}

func (repo *UserRepo) GetByEmail(email string) (*models.User, error) {
	query := "SELECT id, name, surname, email, country, hashed_password, created_at FROM users WHERE email=$1;"

	ctx, cancel := context.WithTimeout(context.Background(), DBTimeout)

	defer cancel()

	row := repo.QueryRowContext(ctx, query, email)

	var res models.User

	err := row.Scan(
		&res.ID,
		&res.Name,
		&res.Surname,
		&res.Email,
		&res.Country,
		&res.HashedPassword,
		&res.CreatedAt,
	)

	if err != nil {
		return nil, apperrors.ErrDatabase(err.Error())
	}

	return &res, nil
}

func (repo *UserRepo) GetByID(userID uint) (*models.User, error) {
	query := "SELECT id, name, surname, email, country, hashed_password, created_at FROM users WHERE id=$1;"

	ctx, cancel := context.WithTimeout(context.Background(), DBTimeout)

	defer cancel()

	row := repo.QueryRowContext(ctx, query, userID)

	var res models.User

	err := row.Scan(
		&res.ID,
		&res.Name,
		&res.Surname,
		&res.Email,
		&res.Country,
		&res.HashedPassword,
		&res.CreatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, apperrors.ErrNotFound("user with such id not found")
		} else {
			return nil, apperrors.ErrDatabase(err.Error())
		}
	}

	return &res, nil

}

func (repo *UserRepo) Update(userObj models.User) (*models.User, error) {
	query := "UPDATE users SET name = $1, surname = $2, email=$3, country=$4 WHERE id=$5 RETURNING id, name, surname, email, country, hashed_password, created_at;"

	ctx, cancel := context.WithTimeout(context.Background(), DBTimeout)

	defer cancel()

	row := repo.QueryRowContext(ctx, query, userObj.Name, userObj.Surname, userObj.Email, userObj.Country, userObj.ID)

	var res models.User

	err := row.Scan(
		&res.ID,
		&res.Name,
		&res.Surname,
		&res.Email,
		&res.Country,
		&res.HashedPassword,
		&res.CreatedAt,
	)

	if err != nil {
		return nil, apperrors.ErrDatabase(err.Error())
	}

	return &res, nil

}

func (repo *UserRepo) Delete(userID uint) error {
	query := "DELETE FROM users WHERE id=$1;"

	ctx, cancel := context.WithTimeout(context.Background(), DBTimeout)

	defer cancel()

	_, err := repo.ExecContext(ctx, query, userID)

	if err != nil {
		return apperrors.ErrDatabase(err.Error())
	}

	return nil
}
