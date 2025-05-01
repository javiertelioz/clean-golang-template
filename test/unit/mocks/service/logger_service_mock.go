package service

import "github.com/stretchr/testify/mock"

type MockLoggerService struct {
	mock.Mock
	// Disable output flag for testing
	DisableOutput bool
}

func NewMockLoggerService() *MockLoggerService {
	return &MockLoggerService{
		DisableOutput: true, // Default to no output for benchmarks
	}
}

func (ls *MockLoggerService) Trace(msg string) {
	if !ls.DisableOutput {
		ls.Called(msg)
	}
}

func (ls *MockLoggerService) Info(msg string) {
	if !ls.DisableOutput {
		ls.Called(msg)
	}
}

func (ls *MockLoggerService) Debug(msg string) {
	if !ls.DisableOutput {
		ls.Called(msg)
	}
}

func (ls *MockLoggerService) Warn(msg string) {
	if !ls.DisableOutput {
		ls.Called(msg)
	}
}

func (ls *MockLoggerService) Error(msg string) {
	if !ls.DisableOutput {
		ls.Called(msg)
	}
}
