package handlers

import (
	"category_service/internal/apperrors"
	"category_service/internal/models"
	"errors"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ICategoryUC interface {
	Create(category models.Category) (*models.Category, error)
	List(userID uint) ([]models.Category, error)
	Update(updInfo models.Category) (*models.Category, error)
	Delete(categoryID uint) error
}

type CategoryHandler struct {
	categoryUC ICategoryUC
}

func NewCategoryHandler(categoryUC ICategoryUC) *CategoryHandler {
	return &CategoryHandler{
		categoryUC: categoryUC,
	}
}

func (h *CategoryHandler) Create(ctx *gin.Context) {
	var req models.CreateCategoryRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Couldn't bind request",
		})
		return
	}

	category, err := h.categoryUC.Create(models.Category{UserID: req.UserID, Name: req.Name})

	if err != nil {
		var appErr *apperrors.AppError
		if errors.As(err, &appErr) {
			log.Print(appErr.Message)
			ctx.JSON(appErr.Status(), appErr.JSONResponse)
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		}

		return

	}

	resp := models.CreateCategoryResponse{
		ID:     category.ID,
		UserID: category.UserID,
		Name:   category.Name,
	}

	ctx.JSON(http.StatusOK, resp)

}

func (h *CategoryHandler) List(ctx *gin.Context) {
	var req models.ListCategoryRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Couldn't bind request",
		})
		return
	}

	categoryList, err := h.categoryUC.List(req.UserID)

	if err != nil {
		var appErr *apperrors.AppError
		if errors.As(err, &appErr) {
			log.Print(appErr.Message)
			ctx.JSON(appErr.Status(), appErr.JSONResponse)
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		}

		return

	}

	resp := models.ListCategoryResponse{
		CategoryList: []models.ElementOfCategoryList{},
	}

	for _, category := range categoryList {
		categoryToAppend := models.ElementOfCategoryList{
			ID:              category.ID,
			UserID:          category.UserID,
			Name:            category.Name,
			SubcategoryList: []models.ElementOfSubcategoryList{},
		}

		for _, subcategory := range category.SubcategoryList {
			subcategoryToAppend := models.ElementOfSubcategoryList{
				ID:         subcategory.ID,
				CategoryID: subcategory.CategoryID,
				Name:       subcategory.Name,
			}
			categoryToAppend.SubcategoryList = append(categoryToAppend.SubcategoryList, subcategoryToAppend)
		}
		resp.CategoryList = append(resp.CategoryList, categoryToAppend)
	}

	ctx.JSON(http.StatusOK, &resp)

}

func (h *CategoryHandler) Update(ctx *gin.Context) {

	categoryIDstr := ctx.Param("id")
	categoryID, err := strconv.ParseUint(categoryIDstr, 10, 32)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	var req models.UpdateCategoryRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Couldn't bind request",
		})
		return
	}

	category, err := h.categoryUC.Update(models.Category{ID: uint(categoryID), Name: req.Name})

	if err != nil {
		var appErr *apperrors.AppError
		if errors.As(err, &appErr) {
			log.Print(appErr.Message)
			ctx.JSON(appErr.Status(), appErr.JSONResponse)
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		}

		return

	}

	resp := models.UpdateCategoryResponse{
		ID:     category.ID,
		UserID: category.UserID,
		Name:   category.Name,
	}

	ctx.JSON(http.StatusOK, resp)
}

func (h *CategoryHandler) Delete(ctx *gin.Context) {

	categoryIDstr := ctx.Param("id")
	categoryID, err := strconv.ParseUint(categoryIDstr, 10, 32)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	err = h.categoryUC.Delete(uint(categoryID))

	if err != nil {
		var appErr *apperrors.AppError
		if errors.As(err, &appErr) {
			log.Print(appErr.Message)
			ctx.JSON(appErr.Status(), appErr.JSONResponse)
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		}

		return

	}

	ctx.Status(http.StatusOK)
}
