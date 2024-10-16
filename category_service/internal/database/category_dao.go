package database

import (
	"category_service/internal/apperrors"
	"category_service/internal/models"
	"context"
	"database/sql"
	"time"
)

type CategoryRepo struct {
	*sql.DB
}

func NewCategoryRepo(db *sql.DB) *CategoryRepo {
	return &CategoryRepo{db}
}

func (repo *CategoryRepo) Create(category models.Category) (*models.Category, error) {
	query := "INSERT INTO categories (user_id, name) VALUES ($1, $2) RETURNING id, user_id, name, created_at;"

	ctx, cancel := context.WithTimeout(context.Background(), DBTimeout)
	defer cancel()

	row := repo.QueryRowContext(ctx, query, category.UserID, category.Name)

	var res models.Category

	err := row.Scan(
		&res.ID,
		&res.UserID,
		&res.Name,
		&res.CreatedAt,
	)

	if err != nil {
		return nil, apperrors.ErrDatabase(err.Error())
	}

	return &res, nil

}

func (repo *CategoryRepo) ListByUserID(userID uint) ([]models.Category, error) {
	list := []models.Category{}
	categoryMap := make(map[uint]models.Category)

	query := "SELECT categories.id, categories.user_id, categories.name, categories.created_at, subcategories.id, subcategories.name, subcategories.created_at FROM categories LEFT JOIN subcategories ON categories.id = subcategories.category_id  WHERE user_id = $1"

	ctx, cancel := context.WithTimeout(context.Background(), DBTimeout)
	defer cancel()

	rows, err := repo.QueryContext(ctx, query, userID)

	if err != nil {
		return nil, apperrors.ErrDatabase(err.Error())
	}

	for rows.Next() {

		var subcategory models.Subcategory
		var categoryID uint
		var userID uint
		var createdAt time.Time
		var categoryName string
		var subcategoryID sql.NullInt64
		var subcategoryName sql.NullString
		var subcategoryCreatedAt sql.NullTime

		err := rows.Scan(
			&categoryID,
			&userID,
			&categoryName,
			&createdAt,
			&subcategoryID,
			&subcategoryName,
			&subcategoryCreatedAt,
		)

		if err != nil {
			return nil, apperrors.ErrDatabase(err.Error())
		}

		category, ok := categoryMap[categoryID]

		if !ok {

			category = models.Category{
				ID:              categoryID,
				UserID:          userID,
				Name:            categoryName,
				CreatedAt:       createdAt,
				SubcategoryList: make([]models.Subcategory, 0),
			}

		}

		if subcategoryID.Valid {
			subcategory.ID = uint(subcategoryID.Int64)
			subcategory.Name = subcategoryName.String
			subcategory.CreatedAt = subcategoryCreatedAt.Time
			subcategory.CategoryID = categoryID

			category.SubcategoryList = append(category.SubcategoryList, subcategory)
		}

		categoryMap[categoryID] = category

	}

	for _, category := range categoryMap {
		list = append(list, category)
	}

	if err := rows.Close(); err != nil {
		return nil, apperrors.ErrDatabase(err.Error())
	}

	return list, nil

}

func (repo *CategoryRepo) Update(category models.Category) (*models.Category, error) {
	query := "UPDATE categories SET name = $1 WHERE id = $2 RETURNING id, user_id, name, created_at;"

	ctx, cancel := context.WithTimeout(context.Background(), DBTimeout)
	defer cancel()

	row := repo.QueryRowContext(ctx, query, category.Name, category.ID)

	var res models.Category

	err := row.Scan(
		&res.ID,
		&res.UserID,
		&res.Name,
		&res.CreatedAt,
	)

	if err != nil {
		return nil, apperrors.ErrDatabase(err.Error())
	}

	return &res, nil
}

func (repo *CategoryRepo) GetByID(categoryID uint) (*models.Category, error) {
	query := "SELECT id, user_id, name, created_at FROM categories WHERE id = $1"

	ctx, cancel := context.WithTimeout(context.Background(), DBTimeout)

	defer cancel()

	row := repo.QueryRowContext(ctx, query, categoryID)

	var res models.Category

	err := row.Scan(
		&res.ID,
		&res.UserID,
		&res.Name,
		&res.CreatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, apperrors.ErrNotFound("category with such id not found")
		} else {
			return nil, apperrors.ErrDatabase(err.Error())
		}
	}

	return &res, nil
}

func (repo *CategoryRepo) Delete(categoryID uint) error {
	query := "DELETE FROM categories WHERE id=$1;"

	ctx, cancel := context.WithTimeout(context.Background(), DBTimeout)

	defer cancel()

	_, err := repo.ExecContext(ctx, query, categoryID)

	if err != nil {
		return apperrors.ErrDatabase(err.Error())
	}

	return nil
}
