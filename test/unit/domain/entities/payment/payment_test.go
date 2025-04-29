package payment_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/suite"

	"github.com/javiertelioz/clean_architecture/pkg/domain/entities/payment"
	"github.com/javiertelioz/clean_architecture/pkg/domain/entities/payment/value_object"
	"github.com/javiertelioz/clean_architecture/pkg/domain/entities/shared"
)

type PaymentTestSuite struct {
	suite.Suite
}

func TestPaymentTestSuite(t *testing.T) {
	suite.Run(t, new(PaymentTestSuite))
}

// ----------- GIVEN ------------

func (suite *PaymentTestSuite) givenValidPaymentOptions() []payment.PaymentOption {
	now := time.Now()

	return []payment.PaymentOption{
		payment.WithID("pay-123"),
		payment.WithUserID("user-456"),
		payment.WithAmount(150.75),
		payment.WithMethod("card"),
		payment.WithTraceabilityBranchID("branch-01"),
		payment.WithTraceabilityDate(now),
		payment.WithTraceabilityExternalID("ext-123"),
		payment.WithTraceabilityTerminalID("term-001"),
	}
}

func (suite *PaymentTestSuite) givenInvalidPaymentOptions() []payment.PaymentOption {
	return []payment.PaymentOption{
		payment.WithID("pay-000"),
		payment.WithUserID(""), // invalid
		payment.WithAmount(0),  // invalid
		payment.WithMethod(""), // invalid
		// traceability missing completely → invalid
	}
}

func (suite *PaymentTestSuite) whenCreatingPayment(opts []payment.PaymentOption) *payment.Payment {
	return payment.NewPayment(opts...)
}

func (suite *PaymentTestSuite) whenValidatingPayment(p *payment.Payment) *shared.ValidationErrors {
	return p.Validate()
}

func (suite *PaymentTestSuite) thenExpectNoValidationErrors(errs *shared.ValidationErrors) {
	suite.Empty(errs.Errors)
}

func (suite *PaymentTestSuite) thenExpectValidationErrors(errs *shared.ValidationErrors) {
	suite.False(errs.IsEmpty())
	suite.NotEmpty(errs.Errors)
}

func (suite *PaymentTestSuite) thenExpectPaymentFieldsToMatch(p *payment.Payment, id, userID string, amount float64, method string) {
	suite.Equal(id, p.GetID())
	suite.Equal(userID, p.GetUserID())
	suite.Equal(value_object.Amount(amount), p.GetAmount())
	suite.Equal(value_object.Method(method), p.GetMethod())
	suite.NotNil(p.GetTraceability().Date)
}

// ----------- TESTS ------------

func (suite *PaymentTestSuite) TestCreateValidPayment() {
	// Given
	opts := suite.givenValidPaymentOptions()

	// When
	p := suite.whenCreatingPayment(opts)

	// Then
	suite.thenExpectPaymentFieldsToMatch(p, "pay-123", "user-456", 150.75, "card")
}

func (suite *PaymentTestSuite) TestValidateValidPayment() {
	// Given
	opts := suite.givenValidPaymentOptions()
	p := suite.whenCreatingPayment(opts)

	// When
	errs := suite.whenValidatingPayment(p)

	// Then
	suite.thenExpectNoValidationErrors(errs)
}

func (suite *PaymentTestSuite) TestValidateInvalidPayment() {
	// Given
	opts := suite.givenInvalidPaymentOptions()
	p := suite.whenCreatingPayment(opts)

	// When
	errs := suite.whenValidatingPayment(p)

	// Then
	suite.thenExpectValidationErrors(errs)
}
