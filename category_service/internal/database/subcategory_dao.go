package database

import (
	"category_service/internal/apperrors"
	"category_service/internal/models"
	"context"
	"database/sql"
)

type SubcategoryRepo struct {
	*sql.DB
}

func NewSubcategoryRepo(db *sql.DB) *SubcategoryRepo {
	return &SubcategoryRepo{db}
}

func (repo *SubcategoryRepo) Create(subcategory models.Subcategory) (*models.Subcategory, error) {
	query := "INSERT INTO subcategories (category_id, name) VALUES ($1, $2) RETURNING id, category_id, name, created_at;"

	ctx, cancel := context.WithTimeout(context.Background(), DBTimeout)
	defer cancel()

	row := repo.QueryRowContext(ctx, query, subcategory.CategoryID, subcategory.Name)

	var res models.Subcategory

	err := row.Scan(
		&res.ID,
		&res.CategoryID,
		&res.Name,
		&res.CreatedAt,
	)

	if err != nil {
		return nil, apperrors.ErrDatabase(err.Error())
	}

	return &res, nil

}

func (repo *SubcategoryRepo) ListByCategoryID(categoryID uint) ([]models.Subcategory, error) {
	list := []models.Subcategory{}
	query := "SELECT id, category_id, name, created_at FROM subcategories WHERE category_id = $1"

	ctx, cancel := context.WithTimeout(context.Background(), DBTimeout)
	defer cancel()

	rows, err := repo.QueryContext(ctx, query, categoryID)

	if err != nil {
		return nil, apperrors.ErrDatabase(err.Error())
	}

	for rows.Next() {
		var subcategory models.Subcategory

		err := rows.Scan(
			&subcategory.ID,
			&subcategory.CategoryID,
			&subcategory.Name,
			&subcategory.CreatedAt,
		)
		if err != nil {
			return nil, apperrors.ErrDatabase(err.Error())
		}
		list = append(list, subcategory)
	}

	if err := rows.Close(); err != nil {
		return nil, apperrors.ErrDatabase(err.Error())
	}

	return list, nil

}

func (repo *SubcategoryRepo) Update(subcategory models.Subcategory) (*models.Subcategory, error) {
	query := "UPDATE subcategories SET name = $1 WHERE id = $2 RETURNING id, category_id, name, created_at;"

	ctx, cancel := context.WithTimeout(context.Background(), DBTimeout)
	defer cancel()

	row := repo.QueryRowContext(ctx, query, subcategory.Name, subcategory.ID)

	var res models.Subcategory

	err := row.Scan(
		&res.ID,
		&res.CategoryID,
		&res.Name,
		&res.CreatedAt,
	)

	if err != nil {
		return nil, apperrors.ErrDatabase(err.Error())
	}

	return &res, nil
}

func (repo *SubcategoryRepo) GetByID(subcategoryID uint) (*models.Subcategory, error) {
	query := "SELECT id, category_id, name, created_at FROM subcategories WHERE id = $1"

	ctx, cancel := context.WithTimeout(context.Background(), DBTimeout)

	defer cancel()

	row := repo.QueryRowContext(ctx, query, subcategoryID)

	var res models.Subcategory

	err := row.Scan(
		&res.ID,
		&res.CategoryID,
		&res.Name,
		&res.CreatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, apperrors.ErrNotFound("subcategory with such id not found")
		} else {
			return nil, apperrors.ErrDatabase(err.Error())
		}
	}

	return &res, nil
}

func (repo *SubcategoryRepo) Delete(subcategoryID uint) error {
	query := "DELETE FROM subcategories WHERE id=$1;"

	ctx, cancel := context.WithTimeout(context.Background(), DBTimeout)

	defer cancel()

	_, err := repo.ExecContext(ctx, query, subcategoryID)

	if err != nil {
		return apperrors.ErrDatabase(err.Error())
	}

	return nil
}
