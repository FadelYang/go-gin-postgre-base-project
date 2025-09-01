package controller

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"project-root/common"
	"project-root/modules/users/dto"
	"project-root/modules/users/service"
	"project-root/tools"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService service.UserService
}

func NewUserController(userService service.UserService) *UserController {
	return &UserController{
		userService: userService,
	}
}

// @Tags 					users
// @Summary				Get Users
// @Description 	get all users
// @Accept 				json
// @Produce 			json
// @Success				200 {object} common.BaseResponse[dto.UserDTO]
// @Router				/users [get]
func (c *UserController) GetAll(ctx *gin.Context) {
	users, err := c.userService.GetAll()
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
func (c *UserController) Create(ctx *gin.Context) {
	var user dto.CreateUser
	if err := ctx.ShouldBindBodyWithJSON(&user); err != nil {
		log.Printf("Failed to create user: %v", err)

		ctx.JSON(http.StatusBadRequest, gin.H{"errors": fmt.Sprintf("%s: %s", service.ErrCreateUserValidate, err.Error())})
		return
	}

	createdExample, err := c.userService.Create(user)
	if err != nil {
		log.Printf("Failed to create user: %v", err)

		var vErr *tools.ValidationError
		if errors.As(err, &vErr) {
			ctx.JSON(http.StatusConflict, vErr)
			return
		}

		ctx.JSON(http.StatusBadRequest, gin.H{"errors": fmt.Sprintf("%s: %s", service.ErrCreateUserValidate, err.Error())})
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
func (c *UserController) Update(ctx *gin.Context) {
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

	updatedUser, err := c.userService.Update(user, parsedUUID)
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
func (c *UserController) Delete(ctx *gin.Context) {
	stringUUID := ctx.Param("uuid")
	parsedUUID, err := tools.StringToUUID(stringUUID)
	if err != nil {
		log.Printf("Failed to delete user: %v", err)

		ctx.JSON(http.StatusBadRequest, gin.H{"errors": fmt.Sprintf("failed to delete a user: %s", err.Error())})
		return
	}

	deletedUser, err := c.userService.Delete(parsedUUID)
	if err != nil {
		log.Printf("Failed to delete user: %v", err)

		switch err {
		case service.ErrUserNotFound:
			ctx.JSON(http.StatusBadRequest, gin.H{"errors": fmt.Sprintf("%s: %s", service.ErrUserNotFound, err.Error())})
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
func (c *UserController) GetByID(ctx *gin.Context) {
	stringUUID := ctx.Param("uuid")
	parsedUUID, err := tools.StringToUUID(stringUUID)
	if err != nil {
		log.Printf("Failed to find user with id %s: %v", parsedUUID, err)

		ctx.JSON(http.StatusBadRequest, gin.H{"errors": err.Error()})
		return
	}

	user, err := c.userService.FindByID(parsedUUID)
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
func (c *UserController) GetByEmail(ctx *gin.Context) {
	email := ctx.Param("email")

	user, err := c.userService.FindByEmail(email)
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
