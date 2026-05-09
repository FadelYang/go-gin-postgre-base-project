package handler

import (
	"errors"
	"fmt"
	"net/http"
	"project-root/common"
	"project-root/modules/auth/usecase"
	userDTO "project-root/modules/users/dto"
	"project-root/tools"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	authUsecase usecase.AuthUsecase
}

func NewAuthHandler(authUsecase usecase.AuthUsecase) *AuthHandler {
	return &AuthHandler{
		authUsecase: authUsecase,
	}
}

// @Tags 					auth
// @Summary				Register new account
// @Description 	Register a new account
// @Accept 				json
// @Produce 			json
// @Success				201 {object} common.BaseResponse[userDTO.UserDTO]
// @Router				/auth/register [post]
// @Param					request body userDTO.CreateUser true "request body for create an user [RAW]"
func (h *AuthHandler) Register(ctx *gin.Context) {
	var user userDTO.CreateUser
	if err := ctx.ShouldBindBodyWithJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"errors": fmt.Sprintf("%s: %s", "failed to create a new account", err.Error())})
		return
	}

	createdUser, err := h.authUsecase.Register(user)
	if err != nil {
		var vErr *tools.ValidationError
		if errors.As(err, &vErr) {
			ctx.JSON(http.StatusConflict, vErr)
			return
		}

		fmt.Printf("errors: %+v", err)

		ctx.JSON(http.StatusBadRequest, gin.H{"errors": fmt.Sprintf("%s: %s", "failed to create a new account", err.Error())})
		return
	}

	ctx.JSON(
		http.StatusCreated,
		common.BaseResponse[userDTO.UserDTO]{
			Status:  http.StatusCreated,
			Message: "register successfully",
			Data: userDTO.UserDTO{
				UUID:        createdUser.ID,
				Username:    createdUser.Username,
				Email:       createdUser.Email,
				FirstName:   createdUser.FirstName,
				LastName:    createdUser.LastName,
				PhoneNumber: createdUser.Phonenumber,
			},
		},
	)
}
