package handler

import (
	"errors"
	"fmt"
	"net/http"
	"project-root/common"
	authDTO "project-root/modules/auth/dto"
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

	createdUser, err := h.authUsecase.Register(ctx.Request.Context(), user)
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

// @Tags 					auth
// @Summary				Login
// @Description 	Login
// @Accept 				json
// @Produce 			json
// @Success				200 {object} common.BaseResponse[authDTO.LoginResponse]
// @Router				/auth/login [post]
// @Param					request body authDTO.LoginDTO true "request body for create an user [RAW]"
func (h *AuthHandler) Login(ctx *gin.Context) {
	var loginForm authDTO.LoginDTO
	if err := ctx.ShouldBindBodyWithJSON(&loginForm); err != nil {
		ctx.JSON(
			http.StatusBadRequest,
			gin.H{
				"errors": fmt.Sprintf("%s: %s", "failed to login an account", err.Error()),
			},
		)
	}

	response, code, err := h.authUsecase.Login(ctx.Request.Context(), loginForm)
	if err != nil {
		ctx.JSON(
			code,
			gin.H{
				"errors": fmt.Sprintf("%s: %s", "failed to login an account", err.Error()),
			},
		)
	}

	ctx.JSON(
		code,
		common.BaseResponse[authDTO.LoginResponse]{
			Message: "login success",
			Data:    *response,
		},
	)
}

// @Tags 					auth
// @Summary				Refresh Login
// @Description 	Get new access token
// @Accept 				json
// @Produce 			json
// @Success				200 {object} common.BaseResponse[authDTO.LoginResponse]
// @Router				/auth/refresh_login [post]
// @Param					request body authDTO.RefreshAccessToken true "request body for get new access token [RAW]"
func (h *AuthHandler) RefreshLogin(ctx *gin.Context) {
	var token authDTO.RefreshAccessToken
	if err := ctx.ShouldBindBodyWithJSON(&token); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"errors": fmt.Sprintf("%s: %s", "failed to generate new access token", err.Error())})
	}

	generatedAccessToken, err := h.authUsecase.RefreshLogin(ctx.Request.Context(), token.RefreshToken)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"errors": fmt.Sprintf("%s: %s", "failed to generate new access token", err.Error())})
		return
	}

	ctx.JSON(
		http.StatusCreated,
		common.BaseResponse[authDTO.LoginResponse]{
			Status:  http.StatusCreated,
			Message: "succesfully generated new access token",
			Data: authDTO.LoginResponse{
				AccessToken:  *generatedAccessToken,
				RefreshToken: token.RefreshToken,
			},
		},
	)
}
