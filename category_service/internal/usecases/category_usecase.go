package usecases

import "category_service/internal/models"

type ICategoryRepo interface {
	Create(category models.Category) (*models.Category, error)
	ListByUserID(userID uint) ([]models.Category, error)
	Update(category models.Category) (*models.Category, error)
	GetByID(categoryID uint) (*models.Category, error)
	Delete(categoryID uint) error
}

type CategoryUC struct {
	categoryRepo ICategoryRepo
}

func NewCategoryUC(categoryR ICategoryRepo) *CategoryUC {
	return &CategoryUC{
		categoryRepo: categoryR,
	}
}

func (uc *CategoryUC) Create(category models.Category) (*models.Category, error) {
	createdCategory, err := uc.categoryRepo.Create(category)

	if err != nil {
		return nil, err
	}

	return createdCategory, nil

}

func (uc *CategoryUC) List(userID uint) ([]models.Category, error) {
	categoryList, err := uc.categoryRepo.ListByUserID(userID)

	if err != nil {
		return nil, err
	}

	return categoryList, nil

}

func (uc *CategoryUC) Update(updInfo models.Category) (*models.Category, error) {
	_, err := uc.categoryRepo.GetByID(updInfo.ID)

	if err != nil {
		return nil, err
	}

	updatedCategory, err := uc.categoryRepo.Update(updInfo)

	if err != nil {
		return nil, err
	}

	return updatedCategory, nil

}

func (uc *CategoryUC) Delete(categoryID uint) error {
	_, err := uc.categoryRepo.GetByID(categoryID)

	if err != nil {
		return err
	}

	err = uc.categoryRepo.Delete(categoryID)

	return err

}
