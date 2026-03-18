package handler

import (
	"log"
	"net/http"
	"project-root/common"
	"project-root/modules/examples/dto"
	"project-root/modules/examples/usecase"

	"github.com/gin-gonic/gin"
)

type ExampleHandler struct {
	exampleUsecase usecase.ExampleUsecase
}

func NewExampleHandler(example usecase.ExampleUsecase) *ExampleHandler {
	return &ExampleHandler{
		exampleUsecase: example,
	}
}

// @Tags 					examples
// @Summary				Get Example
// @Description 	get all example data
// @Accept 				json
// @Produce 			json
// @Success				200 {object} common.BaseResponse[dto.ExampleDTO]
// @Router				/examples [get]
func (c *ExampleHandler) GetExamples(ctx *gin.Context) {
	examples, err := c.exampleUsecase.GetExamples()
	if err != nil {
		log.Printf("Failed to get examples: %v", err)

		ctx.JSON(
			http.StatusInternalServerError,
			common.BaseResponse[dto.ExampleDTO]{
				Status:  http.StatusInternalServerError,
				Message: "Failed to get examples",
				Data:    dto.ExampleDTO{},
			},
		)
		return
	}

	ctx.JSON(
		http.StatusOK,
		common.BaseResponse[[]dto.ExampleDTO]{
			Status:  http.StatusOK,
			Message: "Success get examples",
			Data:    examples,
		},
	)
}

// @Tags 					examples
// @Summary				Post Example
// @Description 	create an example
// @Accept 				json
// @Produce 			json
// @Success				201 {object} common.BaseResponse[dto.ExampleDTO]
// @Router				/examples [post]
// @Param					request body dto.CreateExample true "request body for create an example [RAW]"
func (c *ExampleHandler) Create(ctx *gin.Context) {
	var example dto.ExampleDTO
	if err := ctx.ShouldBindBodyWithJSON(&example); err != nil {
		log.Printf("Failed to create an example: %v", err)

		ctx.JSON(
			http.StatusInternalServerError,
			common.BaseResponse[dto.ExampleDTO]{
				Status:  http.StatusInternalServerError,
				Message: "Failed to create an example",
				Data:    dto.ExampleDTO{},
			},
		)
		return
	}

	createdExample, err := c.exampleUsecase.CreateExample(example)
	if err != nil {
		log.Printf("Failed to create an example: %v", err)

		ctx.JSON(
			http.StatusInternalServerError,
			common.BaseResponse[dto.ExampleDTO]{
				Status:  http.StatusInternalServerError,
				Message: "Failed to create an example",
				Data:    dto.ExampleDTO{},
			},
		)
		return
	}

	ctx.JSON(
		http.StatusOK,
		common.BaseResponse[dto.ExampleDTO]{
			Status:  http.StatusInternalServerError,
			Message: "Success create an example",
			Data:    createdExample,
		},
	)
}
