package controllers

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"net/http"

	dto "github.com/javiertelioz/clean_architecture/pkg/application/dto/hello"
	usecase "github.com/javiertelioz/clean_architecture/pkg/application/use_cases/hello"
	"github.com/javiertelioz/clean_architecture/pkg/domain/contracts/services"
	"github.com/javiertelioz/clean_architecture/pkg/interfaces/presenters"
	"github.com/javiertelioz/clean_architecture/pkg/interfaces/serializers"
)

type HelloController struct {
	useCase       *usecase.HelloUseCase
	loggerService services.LoggerService
}

func NewHelloController(
	useCase *usecase.HelloUseCase,
	loggerService services.LoggerService,
) *HelloController {
	return &HelloController{
		useCase:       useCase,
		loggerService: loggerService,
	}
}

// HelloHandler godoc
//
//	@Summary		Say Hello
//	@Description	Say Hello
//	@Tags			Hello
//	@Accept			json
//	@Produce		json
//	@Param			name			path		string	true	"Name"		default(Joe)
//	@Param			Accept-Language	header		string	false	"Language"	default(en-US)
//	@Success		200				{object}	serializers.HelloSerializer
//	@Router			/api/v1/hello/{name} [get]
func (c *HelloController) HelloHandler(w http.ResponseWriter, r *http.Request) {
	input := dto.HelloInput{
		Name: chi.URLParam(r, "name"),
	}

	output, err := c.useCase.Execute(&input)
	if err != nil {
		c.loggerService.Error(fmt.Sprintf("Failed to execute HelloUseCase: %s", err))
		presenters.JSONError(w, http.StatusBadRequest, "Invalid input or execution error", 1001, err.Error())
		return
	}

	response := serializers.NewHelloSerializer(output)

	c.loggerService.Trace("HelloHandler executed successfully")
	c.loggerService.Debug(fmt.Sprintf("HelloHandler executed successfully with name: %s", input))

	presenters.JSON(w, http.StatusOK, response)
}
