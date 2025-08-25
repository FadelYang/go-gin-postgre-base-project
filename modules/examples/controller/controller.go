package controller

import (
	"log"
	"net/http"
	"project-root/common"
	"project-root/modules/examples/dto"
	"project-root/modules/examples/service"

	"github.com/gin-gonic/gin"
)

type ExampleController struct {
	exampleService service.ExampleService
}

func NewExampleController(example service.ExampleService) *ExampleController {
	return &ExampleController{
		exampleService: example,
	}
}

// @Tags 					examples
// @Summary				Get Example
// @Description 	get all example data
// @Accept 				json
// @Produce 			json
// @Success				200 {object} common.BaseResponse[dto.ExampleDTO]
// @Router				/examples [get]
func (c *ExampleController) GetExamples(ctx *gin.Context) {
	examples, err := c.exampleService.GetExamples()
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
func (c *ExampleController) Create(ctx *gin.Context) {
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

	createdExample, err := c.exampleService.CreateExample(example)
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
