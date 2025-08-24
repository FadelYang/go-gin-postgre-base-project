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

func (c *ExampleController) GetExamples(ctx *gin.Context) {
	examples, err := c.exampleService.GetExamples()
	if err != nil {
		log.Printf("Failed to get examples: %v", err)

		ctx.JSON(
			http.StatusInternalServerError,
			common.BaseResponse{
				Status:  http.StatusInternalServerError,
				Message: "Failed to get examples",
				Data:    nil,
			},
		)
		return
	}

	ctx.JSON(
		http.StatusOK,
		common.BaseResponse{
			Status:  http.StatusOK,
			Message: "Success get examples",
			Data:    examples,
		},
	)
}

func (c *ExampleController) Create(ctx *gin.Context) {
	var example dto.ExampleDTO
	if err := ctx.ShouldBindBodyWithJSON(&example); err != nil {
		log.Printf("Failed to create an example: %v", err)

		ctx.JSON(
			http.StatusInternalServerError,
			common.BaseResponse{
				Status:  http.StatusInternalServerError,
				Message: "Failed to create an example",
				Data:    nil,
			},
		)
		return
	}

	createdExample, err := c.exampleService.CreateExample(example)
	if err != nil {
		log.Printf("Failed to create an example: %v", err)

		ctx.JSON(
			http.StatusInternalServerError,
			common.BaseResponse{
				Status:  http.StatusInternalServerError,
				Message: "Failed to create an example",
				Data:    nil,
			},
		)
		return
	}

	ctx.JSON(
		http.StatusOK,
		common.BaseResponse{
			Status:  http.StatusInternalServerError,
			Message: "Success create an example",
			Data:    createdExample,
		},
	)
}
