package handlers

import (
	"auth_service/internal/apperrors"
	"auth_service/internal/models"
	"errors"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type IUserUseCase interface {
	Register(user models.RegisterUserRequest) (*string, error)
	Login(loginInfo models.LoginUserRequest) (*string, error)
	Verify(verifyInfo models.VerifyRequest) (*uint, error)
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
