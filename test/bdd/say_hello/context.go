package say_hello

import (
	"github.com/javiertelioz/clean_architecture/test/unit/mocks/service"
	"net/http/httptest"

	"github.com/cucumber/godog"
	"github.com/go-chi/chi/v5"

	"github.com/javiertelioz/clean_architecture/pkg/application/use_cases/hello"
	"github.com/javiertelioz/clean_architecture/pkg/interfaces/controllers"
	"github.com/javiertelioz/clean_architecture/pkg/interfaces/routes"
)

type HelloFeatureContext struct {
	router           *chi.Mux
	responseRecorder *httptest.ResponseRecorder
	response         string
}

func NewHelloFeatureContext() *HelloFeatureContext {
	loggerService := service.NewMockLoggerService()
	helloUseCase := hello.NewHelloUseCase(loggerService)
	helloController := controllers.NewHelloController(helloUseCase, loggerService)

	router := chi.NewRouter()
	router.Mount("/api/v1/hello", routes.NewHelloRoutes(helloController).Mount())

	return &HelloFeatureContext{
		router:           router,
		responseRecorder: httptest.NewRecorder(),
	}
}

func (ctx *HelloFeatureContext) InitializeScenario(s *godog.ScenarioContext) {
	s.Step(`^I send a GET request to "([^"]*)"$`, ctx.iSendAGETRequestTo)
	s.Step(`^I should get status code (\d+)$`, ctx.iShouldGetStatusCode)
	s.Step(`^the response should contain "([^"]*)"$`, ctx.theResponseShouldContain)
}
