package hello_test

import (
	"fmt"
	"github.com/javiertelioz/clean_architecture/test/mocks/services"
	"testing"

	"github.com/stretchr/testify/suite"

	dto "github.com/javiertelioz/clean_architecture/pkg/application/dto/hello"
	usecase "github.com/javiertelioz/clean_architecture/pkg/application/use_cases/hello"
	contracts "github.com/javiertelioz/clean_architecture/pkg/domain/contracts/services"
)

type HelloUseCaseTestSuite struct {
	suite.Suite
	useCase       *usecase.HelloUseCase
	loggerService contracts.LoggerService
}

func TestHelloUseCaseTestSuite(t *testing.T) {
	suite.Run(t, new(HelloUseCaseTestSuite))
}

func (suite *HelloUseCaseTestSuite) SetupTest() {
	suite.givenAHelloUseCase()
}

func (suite *HelloUseCaseTestSuite) givenAHelloUseCase() {
	suite.loggerService = services.NewMockLoggerService()
	suite.useCase = usecase.NewHelloUseCase(suite.loggerService)
}

func (suite *HelloUseCaseTestSuite) givenValidInput() *dto.HelloInput {
	return &dto.HelloInput{
		Name: "Joe",
	}
}

func (suite *HelloUseCaseTestSuite) givenInvalidInput() *dto.HelloInput {
	return &dto.HelloInput{
		Name: "",
	}
}

func (suite *HelloUseCaseTestSuite) whenExecutingUseCase(input *dto.HelloInput) (*dto.HelloOutput, error) {
	return suite.useCase.Execute(input)
}

func (suite *HelloUseCaseTestSuite) thenExpectValidResult(result *dto.HelloOutput, err error, expectedMessage string) {
	suite.Nil(err)
	suite.NotNil(result)
	suite.NotEmpty(result.Timestamp)
	suite.Equal(expectedMessage, result.Message)
}

func (suite *HelloUseCaseTestSuite) thenExpectError(result *dto.HelloOutput, err error) {
	suite.NotNil(err)
	suite.Nil(result)
}

func (suite *HelloUseCaseTestSuite) TestExecute_WithValidInput() {
	// Given
	input := suite.givenValidInput()

	// When
	result, err := suite.whenExecutingUseCase(input)

	// Then
	suite.thenExpectValidResult(result, err, fmt.Sprintf("Hello, %s!", input.Name))
}

func (suite *HelloUseCaseTestSuite) TestExecute_WithInvalidInput() {
	// Given
	input := suite.givenInvalidInput()

	// When
	result, err := suite.whenExecutingUseCase(input)

	// Then
	suite.thenExpectError(result, err)
}
