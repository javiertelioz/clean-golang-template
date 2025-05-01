package payment_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/suite"

	dto "github.com/javiertelioz/clean_architecture/pkg/application/dto/payment"
	domain "github.com/javiertelioz/clean_architecture/pkg/domain/entities/payment"
	"github.com/javiertelioz/clean_architecture/pkg/domain/entities/payment/value_object"
)

type CreateTransactionDtoTestSuite struct {
	suite.Suite
}

func TestCreateTransactionDtoTestSuite(t *testing.T) {
	suite.Run(t, new(CreateTransactionDtoTestSuite))
}

func (suite *CreateTransactionDtoTestSuite) givenValidDTO() *dto.CreateTransactionDto {
	now := time.Now()

	return &dto.CreateTransactionDto{
		UserID:      "user-001",
		Amount:      99.99,
		Method:      "card",
		ReferenceID: "ref-abc-123",
		Traceability: dto.TraceabilityDto{
			BranchId:   "branch-01",
			Date:       &now,
			ExternalId: "external-xyz",
			TerminalId: "terminal-09",
		},
	}
}

func (suite *CreateTransactionDtoTestSuite) givenDTOWithoutDate() *dto.CreateTransactionDto {
	return &dto.CreateTransactionDto{
		UserID:      "user-001",
		Amount:      99.99,
		Method:      "card",
		ReferenceID: "ref-abc-123",
		Traceability: dto.TraceabilityDto{
			BranchId:   "branch-01",
			Date:       nil,
			ExternalId: "external-xyz",
			TerminalId: "terminal-09",
		},
	}
}

func (suite *CreateTransactionDtoTestSuite) whenConvertingToDomainOptions(input *dto.CreateTransactionDto) *domain.Payment {
	return domain.NewPayment(input.ToDomainOptions()...)
}

func (suite *CreateTransactionDtoTestSuite) thenExpectFieldsToMatch(p *domain.Payment) {
	suite.Equal("user-001", p.GetUserID())
	suite.Equal(value_object.Amount(99.99), p.GetAmount())
	suite.Equal(value_object.Method("card"), p.GetMethod())
	suite.Equal("branch-01", p.GetTraceability().BranchID)
	suite.Equal("external-xyz", p.GetTraceability().ExternalID)
	suite.Equal("terminal-09", p.GetTraceability().TerminalID)
	suite.NotNil(p.GetTraceability().Date)
}

func (suite *CreateTransactionDtoTestSuite) thenExpectDateToBeNil(p *domain.Payment) {
	suite.Nil(p.GetTraceability().Date)
}

func (suite *CreateTransactionDtoTestSuite) TestToDomainOptions_WithValidDate() {
	// Given
	dto := suite.givenValidDTO()

	// When
	payment := suite.whenConvertingToDomainOptions(dto)

	// Then
	suite.thenExpectFieldsToMatch(payment)
}

func (suite *CreateTransactionDtoTestSuite) TestToDomainOptions_WithoutDate() {
	// Given
	dto := suite.givenDTOWithoutDate()

	// When
	payment := suite.whenConvertingToDomainOptions(dto)

	// Then
	suite.thenExpectDateToBeNil(payment)
}
