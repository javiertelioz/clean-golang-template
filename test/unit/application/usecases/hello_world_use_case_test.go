package hello_test

import (
	"testing"

	"github.com/stretchr/testify/suite"

	dto "github.com/javiertelioz/clean_architecture/pkg/application/dto/hello"
	usecase "github.com/javiertelioz/clean_architecture/pkg/application/use_cases/hello"
	"github.com/javiertelioz/clean_architecture/pkg/domain/entities/hello"
)

type HelloUseCaseTestSuite struct {
	suite.Suite
	useCase *usecase.HelloUseCase
}

func TestHelloUseCaseTestSuite(t *testing.T) {
	suite.Run(t, new(HelloUseCaseTestSuite))
}

func (suite *HelloUseCaseTestSuite) givenAHelloUseCase() {
	suite.useCase = usecase.NewHelloUseCase()
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

func (suite *HelloUseCaseTestSuite) whenExecutingUseCase(input *dto.HelloInput) (*hello.Hello, error) {
	return suite.useCase.Execute(input)
}

func (suite *HelloUseCaseTestSuite) thenExpectValidResult(result *hello.Hello, err error, expectedName string) {
	suite.Nil(err)
	suite.NotNil(result)
	suite.Equal(expectedName, result.GetName())
}

func (suite *HelloUseCaseTestSuite) thenExpectError(result *hello.Hello, err error) {
	suite.NotNil(err)
	suite.Nil(result)
}

func (suite *HelloUseCaseTestSuite) SetupTest() {
	suite.givenAHelloUseCase()
}

func (suite *HelloUseCaseTestSuite) TestExecute_WithValidInput() {
	// Given
	input := suite.givenValidInput()

	// When
	result, err := suite.whenExecutingUseCase(input)

	// Then
	suite.thenExpectValidResult(result, err, input.Name)
}

func (suite *HelloUseCaseTestSuite) TestExecute_WithInvalidInput() {
	// Given
	input := suite.givenInvalidInput()

	// When
	result, err := suite.whenExecutingUseCase(input)

	// Then
	suite.thenExpectError(result, err)
}
