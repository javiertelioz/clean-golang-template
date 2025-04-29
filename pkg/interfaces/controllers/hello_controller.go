package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/javiertelioz/clean_architecture/pkg/application/dto/hello"
	"github.com/javiertelioz/clean_architecture/pkg/interfaces/serializers"
	"net/http"

	"github.com/go-chi/chi/v5"

	usecase "github.com/javiertelioz/clean_architecture/pkg/application/use_cases/hello"
	"github.com/javiertelioz/clean_architecture/pkg/domain/contracts/services"
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
	input := hello.HelloInput{
		Name: chi.URLParam(r, "name"),
	}

	result, err := c.useCase.Execute(&input)
	if err != nil {
		c.loggerService.Error(fmt.Sprintf("Failed to execute HelloUseCase: %s", err))
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	output := serializers.NewHelloSerializer(result.SayHello())

	c.loggerService.Trace("HelloHandler executed successfully")
	c.loggerService.Debug(fmt.Sprintf("HelloHandler executed successfully with name: %s", input))

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	err = json.NewEncoder(w).Encode(output)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
