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

type ISubcategoryUC interface {
	Create(subcategory models.Subcategory) (*models.Subcategory, error)
	List(categoryID uint) ([]models.Subcategory, error)
	Update(updInfo models.Subcategory) (*models.Subcategory, error)
	Delete(subcategoryID uint) error
}

type SubcategoryHandler struct {
	subcategoryUC ISubcategoryUC
}

func NewSubcategoryHandler(subcatUC ISubcategoryUC) *SubcategoryHandler {
	return &SubcategoryHandler{
		subcategoryUC: subcatUC,
	}
}

func (h *SubcategoryHandler) Create(ctx *gin.Context) {
	var req models.CreateSubcategoryRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Couldn't bind request",
		})
		return
	}

	subcategory, err := h.subcategoryUC.Create(models.Subcategory{CategoryID: req.CategoryID, Name: req.Name})

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

	resp := models.CreateSubcategoryResponse{
		ID:         subcategory.ID,
		CategoryID: subcategory.CategoryID,
		Name:       subcategory.Name,
	}

	ctx.JSON(http.StatusOK, resp)

}

func (h *SubcategoryHandler) List(ctx *gin.Context) {
	var req models.ListSubcategoryRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Couldn't bind request",
		})
		return
	}

	subcategoryList, err := h.subcategoryUC.List(req.CategoryID)

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

	resp := models.ListSubcategoryResponse{
		SubcategoryList: []models.ElementOfSubcategoryList{},
	}

	for _, subcategory := range subcategoryList {
		categoryToAppend := models.ElementOfSubcategoryList{
			ID:         subcategory.ID,
			CategoryID: subcategory.CategoryID,
			Name:       subcategory.Name,
		}
		resp.SubcategoryList = append(resp.SubcategoryList, categoryToAppend)
	}

	ctx.JSON(http.StatusOK, &resp)

}

func (h *SubcategoryHandler) Update(ctx *gin.Context) {

	subcategoryIDstr := ctx.Param("id")
	subcategoryID, err := strconv.ParseUint(subcategoryIDstr, 10, 32)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	var req models.UpdateSubcategoryRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Couldn't bind request",
		})
		return
	}

	subcategory, err := h.subcategoryUC.Update(models.Subcategory{ID: uint(subcategoryID), Name: req.Name})

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

	resp := models.UpdateSubcategoryResponse{
		ID:         subcategory.ID,
		CategoryID: subcategory.CategoryID,
		Name:       subcategory.Name,
	}

	ctx.JSON(http.StatusOK, resp)
}

func (h *SubcategoryHandler) Delete(ctx *gin.Context) {

	subcategoryIDstr := ctx.Param("id")
	subcategoryID, err := strconv.ParseUint(subcategoryIDstr, 10, 32)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	err = h.subcategoryUC.Delete(uint(subcategoryID))

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
