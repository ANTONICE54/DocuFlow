package database

import (
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

func (repo *UserRepo) Create(user models.User) error {
	query := "INSERT INTO users(name, surname, email, country, hashed_password) values ($1, $2, $3, $4, $5);"

	ctx, cancel := context.WithTimeout(context.Background(), DBTimeout)

	defer cancel()

	_, err := repo.ExecContext(ctx, query, user.Name, user.Surname, user.Email, user.Country, user.HashedPassword)

	if err != nil {
		return err
	}

	return nil
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
		return nil, err
	}

	return &res, nil
}
