package handlers

import (
	"auth_service/internal/apperrors"
	"auth_service/internal/models"
	"errors"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type IUserUseCase interface {
	Register(user models.RegisterUserRequest) (*string, error)
	Login(loginInfo models.LoginUserRequest) (*string, error)
	Verify(verifyInfo models.VerifyRequest) (*uint, error)
	Get(userID uint) (*models.User, error)
	Update(updInfo models.UpdateUserRequest) (*models.User, error)
	Delete(userID uint) error
}

type UserHandler struct {
	userUC IUserUseCase
}

func NewUserHandler(userUseCase IUserUseCase) *UserHandler {
	return &UserHandler{
		userUC: userUseCase,
	}
}

func (h *UserHandler) Register(ctx *gin.Context) {
	var req models.RegisterUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Couldn't bind request",
		})
		return
	}

	token, err := h.userUC.Register(req)

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

	rsp := models.RegisterUserResponse{
		Token: *token,
	}

	ctx.JSON(http.StatusOK, rsp)

}

func (h *UserHandler) Login(ctx *gin.Context) {
	var req models.LoginUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Couldn't bind request",
		})
		return
	}

	token, err := h.userUC.Login(req)

	if err != nil {
		var appErr *apperrors.AppError
		if errors.As(err, &appErr) {
			log.Print(appErr.Message)
			ctx.JSON(appErr.Status(), appErr.JSONResponse)
		} else {
			log.Print(err.Error())
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		}

		return

	}

	rsp := models.LoginUserResponse{
		Token: *token,
	}

	ctx.JSON(http.StatusOK, rsp)

}

func (h *UserHandler) Verify(ctx *gin.Context) {
	var req models.VerifyRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Couldn't bind request",
		})
		return
	}

	userID, err := h.userUC.Verify(req)

	if err != nil {
		var appErr *apperrors.AppError
		if errors.As(err, &appErr) {
			log.Print(appErr.Message)
			ctx.JSON(appErr.Status(), appErr.JSONResponse)
		} else {
			log.Print(err.Error())
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		}

		return

	}

	resp := models.VerifyResponse{
		UserID: *userID,
	}

	ctx.JSON(http.StatusOK, resp)
}

func (h *UserHandler) Get(ctx *gin.Context) {
	userIDstr := ctx.Param("id")
	userID, err := strconv.ParseUint(userIDstr, 10, 32)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	user, err := h.userUC.Get(uint(userID))

	if err != nil {
		var appErr *apperrors.AppError
		if errors.As(err, &appErr) {
			log.Print(appErr.Message)
			ctx.JSON(appErr.Status(), appErr.JSONResponse)
		} else {
			log.Print(err.Error())
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		}

		return

	}

	resp := models.GetUserResponse{
		ID:      user.ID,
		Name:    user.Name,
		Surname: user.Surname,
		Email:   user.Email,
		Country: user.Country,
	}

	ctx.JSON(http.StatusOK, &resp)

}

func (h *UserHandler) Update(ctx *gin.Context) {
	userIDstr := ctx.Param("id")
	userID, err := strconv.ParseUint(userIDstr, 10, 32)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	var req models.UpdateUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Couldn't bind request",
		})
		return
	}

	req.ID = uint(userID)

	user, err := h.userUC.Update(req)

	if err != nil {
		var appErr *apperrors.AppError
		if errors.As(err, &appErr) {
			log.Print(appErr.Message)
			ctx.JSON(appErr.Status(), appErr.JSONResponse)
		} else {
			log.Print(err.Error())
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		}

		return

	}

	resp := models.UpdateUserResponse{
		ID:      user.ID,
		Name:    user.Name,
		Surname: user.Surname,
		Email:   user.Email,
		Country: user.Country,
	}

	ctx.JSON(http.StatusOK, &resp)
}

func (h *UserHandler) Delete(ctx *gin.Context) {
	userIDstr := ctx.Param("id")
	userID, err := strconv.ParseUint(userIDstr, 10, 32)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	err = h.userUC.Delete(uint(userID))

	if err != nil {
		var appErr *apperrors.AppError
		if errors.As(err, &appErr) {
			log.Print(appErr.Message)
			ctx.JSON(appErr.Status(), appErr.JSONResponse)
		} else {
			log.Print(err.Error())
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		}

		return

	}

	ctx.Status(http.StatusOK)
}
