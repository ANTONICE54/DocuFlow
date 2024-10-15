package usecases

import "category_service/internal/models"

type ISubcategoryRepo interface {
	Create(subcategory models.Subcategory) (*models.Subcategory, error)
	ListByCategoryID(categoryID uint) ([]models.Subcategory, error)
	Update(subcategory models.Subcategory) (*models.Subcategory, error)
	GetByID(subcategoryID uint) (*models.Subcategory, error)
	Delete(subcategoryID uint) error
}

type SubcategoryUC struct {
	subcategoryRepo ISubcategoryRepo
}

func NewSubcategoryUC(subcatRepo ISubcategoryRepo) *SubcategoryUC {
	return &SubcategoryUC{
		subcategoryRepo: subcatRepo,
	}
}

func (uc *SubcategoryUC) Create(subcategory models.Subcategory) (*models.Subcategory, error) {
	createdSubcategory, err := uc.subcategoryRepo.Create(subcategory)

	if err != nil {
		return nil, err
	}

	return createdSubcategory, nil

}

func (uc *SubcategoryUC) List(categoryID uint) ([]models.Subcategory, error) {
	subcategoryList, err := uc.subcategoryRepo.ListByCategoryID(categoryID)

	if err != nil {
		return nil, err
	}

	return subcategoryList, nil

}

func (uc *SubcategoryUC) Update(updInfo models.Subcategory) (*models.Subcategory, error) {
	_, err := uc.subcategoryRepo.GetByID(updInfo.ID)

	if err != nil {
		return nil, err
	}

	updatedSubcategory, err := uc.subcategoryRepo.Update(updInfo)

	if err != nil {
		return nil, err
	}

	return updatedSubcategory, nil

}

func (uc *SubcategoryUC) Delete(subcategoryID uint) error {
	_, err := uc.subcategoryRepo.GetByID(subcategoryID)

	if err != nil {
		return err
	}

	err = uc.subcategoryRepo.Delete(subcategoryID)

	return err

}
