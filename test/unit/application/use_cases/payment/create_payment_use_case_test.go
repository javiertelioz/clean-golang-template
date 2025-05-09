package payment_test

import (
	"errors"
	"github.com/javiertelioz/clean_architecture/test/mocks/repository"
	"testing"
	"time"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"

	dto "github.com/javiertelioz/clean_architecture/pkg/application/dto/payment"
	"github.com/javiertelioz/clean_architecture/pkg/application/use_cases/payment"
	domain "github.com/javiertelioz/clean_architecture/pkg/domain/entities/payment"
)

type CreatePaymentUseCaseTestSuite struct {
	suite.Suite
	useCase    *payment.CreatePaymentUseCase
	repository *repository.MockPaymentRepository
}

func TestCreatePaymentUseCaseTestSuite(t *testing.T) {
	suite.Run(t, new(CreatePaymentUseCaseTestSuite))
}

func (suite *CreatePaymentUseCaseTestSuite) SetupTest() {
	suite.repository = repository.NewMockPaymentRepository()
	suite.useCase = payment.NewCreatePaymentUseCase(suite.repository)
}

func (suite *CreatePaymentUseCaseTestSuite) givenValidDTO() *dto.CreateTransactionInput {
	now := time.Now()

	return &dto.CreateTransactionInput{
		UserID:      "user-001",
		Amount:      250.0,
		Method:      "debit_card",
		ReferenceID: "ref-123",
		Traceability: dto.TraceabilityDto{
			BranchId:   "BR001",
			Date:       &now,
			ExternalId: "EXT123",
			TerminalId: "TERM456",
		},
	}
}

func (suite *CreatePaymentUseCaseTestSuite) givenInvalidDTO() *dto.CreateTransactionInput {
	return &dto.CreateTransactionInput{
		UserID: "",
		Amount: 0,
		Method: "",
		Traceability: dto.TraceabilityDto{
			BranchId:   "",
			Date:       nil,
			ExternalId: "",
			TerminalId: "",
		},
	}
}

func (suite *CreatePaymentUseCaseTestSuite) whenExecuting(input *dto.CreateTransactionInput) (*domain.Payment, error) {
	return suite.useCase.Execute(input)
}

func (suite *CreatePaymentUseCaseTestSuite) thenExpectSuccess(p *domain.Payment, err error) {
	suite.NoError(err)
	suite.NotNil(p)
}

func (suite *CreatePaymentUseCaseTestSuite) thenExpectValidationError(p *domain.Payment, err error) {
	suite.Nil(p)
	suite.Error(err)
}

func (suite *CreatePaymentUseCaseTestSuite) thenExpectRepositoryCalledOnce() {
	suite.repository.AssertNumberOfCalls(suite.T(), "Save", 1)
}

func (suite *CreatePaymentUseCaseTestSuite) TestExecute_WithValidDTO_ShouldSucceed() {
	// Given
	dto := suite.givenValidDTO()
	expected := domain.NewPayment(dto.ToDomainOptions()...)
	suite.repository.On("Save", mock.Anything).Return(expected, nil)

	// When
	result, err := suite.whenExecuting(dto)

	// Then
	suite.thenExpectSuccess(result, err)
	suite.thenExpectRepositoryCalledOnce()
}

func (suite *CreatePaymentUseCaseTestSuite) TestExecute_WithInvalidDTO_ShouldFailValidation() {
	// Given
	dto := suite.givenInvalidDTO()

	// When
	result, err := suite.whenExecuting(dto)

	// Then
	suite.thenExpectValidationError(result, err)
	suite.repository.AssertNotCalled(suite.T(), "Save")
}

func (suite *CreatePaymentUseCaseTestSuite) TestExecute_WhenRepositoryFails_ShouldReturnError() {
	// Given
	dto := suite.givenValidDTO()
	expected := domain.NewPayment(dto.ToDomainOptions()...)
	suite.repository.On("Save", mock.Anything).Return(expected, errors.New("repository failure"))

	// When
	result, err := suite.whenExecuting(dto)

	// Then
	suite.Error(err)
	suite.Nil(result)
	suite.thenExpectRepositoryCalledOnce()
}
