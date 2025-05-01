package hello_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/suite"

	"github.com/javiertelioz/clean_architecture/pkg/domain/entities/hello"
	"github.com/javiertelioz/clean_architecture/pkg/domain/entities/shared"
)

type HelloTestSuite struct {
	suite.Suite
}

func TestHelloTestSuite(t *testing.T) {
	suite.Run(t, new(HelloTestSuite))
}

func (suite *HelloTestSuite) givenValidHelloOptions() []hello.HelloOption {
	timestamp := time.Now().UnixMilli()
	return []hello.HelloOption{
		hello.WithName("John"),
		hello.WithTimestamp(timestamp),
	}
}

func (suite *HelloTestSuite) givenInvalidHelloOptions() []hello.HelloOption {
	return []hello.HelloOption{
		hello.WithName(""),
		hello.WithTimestamp(0),
	}
}

func (suite *HelloTestSuite) whenCreatingHello(opts []hello.HelloOption) *hello.Hello {
	return hello.NewHello(opts...)
}

func (suite *HelloTestSuite) whenValidatingHello(h *hello.Hello) *shared.ValidationErrors {
	return h.Validate()
}

func (suite *HelloTestSuite) whenGettingSayHello(h *hello.Hello) string {
	return h.SayHello()
}

func (suite *HelloTestSuite) thenExpectNoValidationErrors(errs *shared.ValidationErrors) {
	suite.Empty(errs.Errors)
}

func (suite *HelloTestSuite) thenExpectValidationErrors(errs *shared.ValidationErrors) {
	suite.False(errs.IsEmpty())
	suite.NotEmpty(errs.Errors)
}

func (suite *HelloTestSuite) thenExpectHelloFieldsToMatch(h *hello.Hello, name string) {
	suite.Equal(name, h.GetName())
	suite.NotZero(h.GetTimestamp())
}

func (suite *HelloTestSuite) thenExpectHelloMessage(message string, expected string) {
	suite.Equal(expected, message)
}

func (suite *HelloTestSuite) TestCreateValidHello() {
	// Given
	opts := suite.givenValidHelloOptions()

	// When
	h := suite.whenCreatingHello(opts)

	// Then
	suite.thenExpectHelloFieldsToMatch(h, "John")
}

func (suite *HelloTestSuite) TestValidateValidHello() {
	// Given
	opts := suite.givenValidHelloOptions()
	h := suite.whenCreatingHello(opts)

	// When
	errs := suite.whenValidatingHello(h)

	// Then
	suite.thenExpectNoValidationErrors(errs)
}

func (suite *HelloTestSuite) TestValidateInvalidHello() {
	// Given
	opts := suite.givenInvalidHelloOptions()
	h := suite.whenCreatingHello(opts)

	// When
	errs := suite.whenValidatingHello(h)

	// Then
	suite.thenExpectValidationErrors(errs)
}

func (suite *HelloTestSuite) TestSayHelloMessage() {
	// Given
	opts := suite.givenValidHelloOptions()
	h := suite.whenCreatingHello(opts)

	// When
	message := suite.whenGettingSayHello(h)

	// Then
	suite.thenExpectHelloMessage(message, "Hello, John!")
}
