package handler

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"project-root/common"
	"project-root/modules/users/dto"
	"project-root/modules/users/usecase"
	"project-root/tools"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userUsecase usecase.UserUsecase
}

func NewUserHandler(userUsecase usecase.UserUsecase) *UserHandler {
	return &UserHandler{
		userUsecase: userUsecase,
	}
}

// @Tags 					users
// @Summary				Get Users
// @Description 	get all users
// @Accept 				json
// @Produce 			json
// @Success				200 {object} common.BaseResponse[dto.UserDTO]
// @Router				/users [get]
func (c *UserHandler) GetAll(ctx *gin.Context) {
	users, err := c.userUsecase.GetAll()
	if err != nil {
		log.Printf("Failed to get users: %v", err)

		ctx.JSON(http.StatusBadRequest, gin.H{"errors": "failed to get users data"})
		return
	}

	ctx.JSON(
		http.StatusOK,
		common.BaseResponse[[]dto.UserDTO]{
			Status:  http.StatusOK,
			Message: "Successfully get user data",
			Data:    users,
		},
	)
}

// @Tags 					users
// @Summary				Create User
// @Description 	create an user
// @Accept 				json
// @Produce 			json
// @Success				201 {object} common.BaseResponse[dto.UserDTO]
// @Router				/users [post]
// @Param					request body dto.CreateUser true "request body for create an user [RAW]"
func (c *UserHandler) Create(ctx *gin.Context) {
	var user dto.CreateUser
	if err := ctx.ShouldBindBodyWithJSON(&user); err != nil {
		log.Printf("Failed to create user: %v", err)

		ctx.JSON(http.StatusBadRequest, gin.H{"errors": fmt.Sprintf("%s: %s", usecase.ErrCreateUserValidate, err.Error())})
		return
	}

	createdExample, err := c.userUsecase.Create(ctx.Request.Context(), user)
	if err != nil {
		log.Printf("Failed to create user: %v", err)

		var vErr *tools.ValidationError
		if errors.As(err, &vErr) {
			ctx.JSON(http.StatusConflict, vErr)
			return
		}

		ctx.JSON(http.StatusBadRequest, gin.H{"errors": fmt.Sprintf("%s: %s", usecase.ErrCreateUserValidate, err.Error())})
		return
	}

	ctx.JSON(
		http.StatusOK,
		common.BaseResponse[dto.UserDTO]{
			Status:  http.StatusCreated,
			Message: "User created successfully",
			Data:    createdExample,
		},
	)
}

// @Tags 					users
// @Summary				Update User
// @Description 	update an user
// @Accept 				json
// @Produce 			json
// @Success				201 {object} common.BaseResponse[dto.UserDTO]
// @Router				/users/{uuid} [put]
// @Param					uuid path string true "UUID"
// @Param					request body dto.UpdateUser true "request body for update an example [RAW]"
func (c *UserHandler) Update(ctx *gin.Context) {
	stringUUID := ctx.Param("uuid")
	parsedUUID, err := tools.StringToUUID(stringUUID)
	if err != nil {
		log.Printf("Failed to update user: %v", err)

		ctx.JSON(http.StatusBadRequest, gin.H{"errors": fmt.Sprintf("failed to update a user: %s", err.Error())})
		return
	}

	var user dto.UpdateUser
	if err := ctx.ShouldBindBodyWithJSON(&user); err != nil {
		log.Printf("Failed to update user: %v", err)

		ctx.JSON(http.StatusBadRequest, gin.H{"errors": fmt.Sprintf("failed to update a user: %s", err.Error())})
		return
	}

	updatedUser, err := c.userUsecase.Update(ctx.Request.Context(), user, parsedUUID)
	if err != nil {
		log.Printf("Failed to update user: %v", err)

		var vErr *tools.ValidationError
		if errors.As(err, &vErr) {
			ctx.JSON(http.StatusConflict, vErr)
			return
		}

		ctx.JSON(http.StatusInternalServerError, gin.H{"errors": err.Error()})
		return
	}

	ctx.JSON(
		http.StatusOK,
		common.BaseResponse[dto.UserDTO]{
			Status:  http.StatusCreated,
			Message: "User updated successfully",
			Data:    updatedUser,
		},
	)
}

// @Tags 					users
// @Summary				Delete User
// @Description 	Delete an user
// @Accept 				json
// @Produce 			json
// @Success				200 {object} common.BaseResponse[dto.UserDTO]
// @Router				/users/{uuid} [delete]
// @Param					uuid path string true "UUID"
func (c *UserHandler) Delete(ctx *gin.Context) {
	stringUUID := ctx.Param("uuid")
	parsedUUID, err := tools.StringToUUID(stringUUID)
	if err != nil {
		log.Printf("Failed to delete user: %v", err)

		ctx.JSON(http.StatusBadRequest, gin.H{"errors": fmt.Sprintf("failed to delete a user: %s", err.Error())})
		return
	}

	deletedUser, err := c.userUsecase.Delete(ctx.Request.Context(), parsedUUID)
	if err != nil {
		log.Printf("Failed to delete user: %v", err)

		switch err {
		case usecase.ErrUserNotFound:
			ctx.JSON(http.StatusBadRequest, gin.H{"errors": fmt.Sprintf("%s: %s", usecase.ErrUserNotFound, err.Error())})
		default:
			ctx.JSON(http.StatusBadRequest, gin.H{"errors": fmt.Sprintf("failed to delete a user: %s", err.Error())})
		}

		return
	}

	ctx.JSON(
		http.StatusOK,
		common.BaseResponse[dto.UserDTO]{
			Status:  http.StatusCreated,
			Message: "User deleted successfully",
			Data:    deletedUser,
		},
	)
}

// @Tags 					users
// @Summary				Find User By Its UUID
// @Description 	Find user by id
// @Accept 				json
// @Produce 			json
// @Success				200 {object} common.BaseResponse[dto.UserDTO]
// @Router				/users/{uuid} [get]
// @Param					uuid path string true "UUID"
func (c *UserHandler) GetByID(ctx *gin.Context) {
	stringUUID := ctx.Param("uuid")
	parsedUUID, err := tools.StringToUUID(stringUUID)
	if err != nil {
		log.Printf("Failed to find user with id %s: %v", parsedUUID, err)

		ctx.JSON(http.StatusBadRequest, gin.H{"errors": err.Error()})
		return
	}

	user, err := c.userUsecase.FindByID(ctx.Request.Context(), parsedUUID)
	if err != nil {
		log.Printf("Failed to find user with id %s: %v", parsedUUID, err)

		ctx.JSON(http.StatusBadRequest, gin.H{"errors": err.Error()})
		return
	}

	ctx.JSON(
		http.StatusOK,
		common.BaseResponse[dto.UserDTO]{
			Status:  http.StatusOK,
			Message: fmt.Sprintf("Successfully get user with id %s", user.UUID),
			Data:    user,
		},
	)
}

// @Tags 					users
// @Summary				Find User By Its Email
// @Description 	Find user by email
// @Accept 				json
// @Produce 			json
// @Success				200 {object} common.BaseResponse[dto.UserDTO]
// @Router				/users/email/{email} [get]
// @Param					email path string true "User Email"
func (c *UserHandler) GetByEmail(ctx *gin.Context) {
	email := ctx.Param("email")

	user, err := c.userUsecase.FindByEmail(ctx, email)
	if err != nil {
		log.Printf("Failed to find user with email %s: %v", email, err)

		ctx.JSON(http.StatusBadRequest, gin.H{"errors": err.Error()})
		return
	}

	ctx.JSON(
		http.StatusOK,
		common.BaseResponse[dto.UserDTO]{
			Status:  http.StatusOK,
			Message: fmt.Sprintf("Successfully get user with email %s", email),
			Data:    user,
		},
	)
}

// @Tags 					users
// @Summary				Update User Role
// @Description 	update a role of user
// @Accept 				json
// @Produce 			json
// @Success				201 {object} common.BaseResponse[dto.UserDTO]
// @Router				/users/{uuid}/roles [put]
// @Param					uuid path string true "UUID"
// @Param					request body dto.UpdateUserRole true "request body for update a role of user [RAW]"
func (c *UserHandler) UpdateRole(ctx *gin.Context) {
	stringUserUUID := ctx.Param("uuid")
	parseUserUUID, err := tools.StringToUUID(stringUserUUID)
	if err != nil {
		errMsg := fmt.Sprintf("failed to update user role: %v", err)

		log.Println(errMsg)
		ctx.JSON(http.StatusBadRequest, gin.H{"errors": errMsg})
		return
	}

	var updatedRole dto.UpdateUserRole
	if err := ctx.ShouldBindBodyWithJSON(&updatedRole); err != nil {
		errMsg := fmt.Sprintf("failed to update user role: %v", err)

		log.Println(errMsg)
		ctx.JSON(http.StatusBadRequest, gin.H{"errors": errMsg})
		return
	}

	code, err := c.userUsecase.UpdateRole(ctx.Request.Context(), parseUserUUID, updatedRole)
	if err := ctx.ShouldBindBodyWithJSON(&updatedRole); err != nil {
		errMsg := fmt.Sprintf("failed to update user role: %v", err)

		log.Println(errMsg)
		ctx.JSON(code, gin.H{"errors": errMsg})
		return
	}
}
